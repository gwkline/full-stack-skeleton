package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/hibiken/asynqmon"
	"github.com/redis/go-redis/v9"

	_ "time/tzdata"

	"strings"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gwkline/full-stack-skeleton/backend/graph/resolver"
	"github.com/gwkline/full-stack-skeleton/backend/handlers"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/auth"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/cron"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/queue"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/search"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/some_service"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/util"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/workers"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

const defaultPort = "8888"

func main() {
	router, cron := apiInit()
	defer cron.Stop()

	port := util.GetEnvWithFallback("PORT", defaultPort)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router.Handler(),
	}

	go func() {
		fmt.Printf("Environment: %s\n", os.Getenv("ENV"))
		fmt.Printf("Listening on port %v\n", port)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("Surpassed timeout of 5 seconds.")
	}

	fmt.Println("Goodbye!")
}

func apiInit() (*gin.Engine, *cron.Service) {
	db, err := repo.Init()
	if err != nil {
		fmt.Printf("failed to initialize database: %v\n", err)
	}

	fmt.Println("Initializing Redis")
	ops, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		fmt.Printf("failed parsing redis: %v", err)
	}
	redis := redis.NewClient(ops)

	nrApp, err := newRelicInit()
	if err != nil {
		fmt.Printf("failed getting NR app: %s\n", err)
	}

	services, err := servicesInit(db, redis, nrApp, http.DefaultClient, "https://login.salesforce.com/services/oauth2/token", "")
	if err != nil {
		fmt.Printf("failed to initialize services: %v", err)
	}

	go services.Queue.StartServer(context.Background(), services.Workers)

	middleware := middlewareInit(db, nrApp, services.Auth)
	router := routerInit(db, redis, services, middleware)
	handler := handlers.New(db, redis, services)

	sentryInit()
	routesInit(router, handler)

	cron := cron.Init(services.Queue)
	cron.SetupJobs(services)

	return router, cron
}

func servicesInit(db *repo.Repository, r *redis.Client, nrApp *newrelic.Application, httpClient *http.Client, sfURL, mgURL string) (*types.Services, error) {
	fmt.Println("Initializing Services")
	ops := util.NewAsynqOps()

	auth := auth.Init(db)
	queue := queue.Init(&ops)
	example := some_service.Init(db)
	search := search.Init(db)
	worker := workers.Init(db, nrApp, example)

	return &types.Services{
		Auth:           auth,
		Queue:          queue,
		Search:         search,
		Workers:        worker,
		ExampleService: example,
	}, nil
}

func routerInit(db *repo.Repository, r *redis.Client, s *types.Services, middleware []gin.HandlerFunc) *gin.Engine {
	fmt.Println("Initializing Router")
	g := gin.New()

	g.Use(
		loggerMiddleware(),
		gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression),
	)

	g.Use(middleware...)

	pprof.Register(g)

	return g
}

func routesInit(router *gin.Engine, h *handlers.Handler) {
	router.StaticFile("/schema.gql", "./graph/schema.gql")

	router.OPTIONS("/graphql", func(c *gin.Context) { c.Status(http.StatusOK) })
	router.GET("/graphql", h.Graphql())
	router.POST("/graphql", h.Graphql())

	router.GET("/", h.Playground())

	router.POST("/someWebhookEndpoint", h.SomeWebhookHandler(nil))

	handler := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitoring", // RootPath specifies the root for asynqmon app
		RedisConnOpt: util.NewAsynqOps(),
	})

	router.Any("/monitoring/*a", gin.WrapH(handler))
}

func sentryInit() bool {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_BE_DSN"),
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		fmt.Printf("failed sentry initialization: %v\n", err)
		return false
	}

	return true
}

func newRelicInit() (*newrelic.Application, error) {
	enabled := os.Getenv("ENV") == "production"
	return newrelic.NewApplication(
		newrelic.ConfigAppName("NR_APP_NAME"),
		newrelic.ConfigLicense("NR_APP_LICENSE"),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigEnabled(enabled),
	)
}

func middlewareInit(db *repo.Repository, nrApp *newrelic.Application, auth types.IAuth) []gin.HandlerFunc {
	fmt.Println("Initializing Middleware")
	sentryMiddleware := sentrygin.New(sentrygin.Options{})

	var corsFunc gin.HandlerFunc
	switch os.Getenv("ENV") {
	case "production":
		corsFunc = cors.New(cors.Config{
			AllowOrigins:     []string{os.Getenv("FE_DOMAIN")},
			AllowMethods:     []string{"OPTIONS", "POST", "GET", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		})
	default:
		corsFunc = cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"OPTIONS", "POST", "GET", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			AllowCredentials: true,
		})
	}

	middlewares := []gin.HandlerFunc{
		sentryMiddleware,
		corsFunc,
		authenticationMiddleware(db, auth),
		dataloaderMiddleware(db),
	}

	if os.Getenv("ENV") == "production" {
		newRelicMiddleware := nrgin.Middleware(nrApp)
		middlewares = append(middlewares, newRelicMiddleware)
	}

	return middlewares
}

func authenticationMiddleware(db *repo.Repository, auth types.IAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = strings.ReplaceAll(token, "Bearer", "")
		token = strings.Trim(token, " ")

		if token == "" {
			c.Next()
			return
		}

		claims, err := auth.ValidateToken(c, token)
		if err != nil || claims == nil {
			fmt.Printf("failed to validate token: %v\n", err)
			c.Next()
			return
		}

		email := claims.Email
		user, err := db.User.FindBy([]types.Filter{{Key: "email", Value: email}})
		if err != nil {
			fmt.Printf("failed to find user by email: %v\n", err)
			c.Next()
			return
		}

		ctx := auth.SetCurrentUser(c, user)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func dataloaderMiddleware(repo *repo.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		loaders := resolver.NewLoaders(repo)
		ctx := context.WithValue(c.Request.Context(), types.LoadersKey, loaders)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if !strings.HasPrefix(path, "/monitoring/") {
			// Use the default logger
			gin.Logger()(c)
		} else {
			// Skip logging and call the next middleware directly
			c.Next()
		}
	}
}
