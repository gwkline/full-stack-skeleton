package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gwkline/full-stack-infra/backend/database"
	"github.com/gwkline/full-stack-infra/backend/graph"
	"github.com/gwkline/full-stack-infra/backend/services/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8888"

type Config struct {
	Port, Env, DatabaseUser, DatabasePassword, DatabaseName, SentryDSN string
}

func main() {
	cfg := LoadConfig()

	InitSentry(cfg.SentryDSN)

	if cfg.Env == "production" || cfg.Env == "development" {
		database.InitDB(cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName)
	}

	router := SetupRouter(cfg.Env)
	router.Run(":" + cfg.Port)
}

func SetupRouter(env string) *gin.Engine {

	// Router config fun
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Middleware
	sentryMiddleware := sentrygin.New(sentrygin.Options{})
	defaultCORS := cors.Default()
	prodCORS := cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:8080"},
		AllowMethods:  []string{"OPTIONS", "POST", "GET", "PUT", "DELETE"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
	})

	// Apply Middleware based on environment
	switch env {
	case "production":
		router.Use(sentryMiddleware, prodCORS)
	default:
		router.Use(sentryMiddleware, defaultCORS)
	}

	// Routes
	router.OPTIONS("/graphql", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Expose schema for introspection
	// TODO: Add authorization here
	router.StaticFile("/schema.graphqls", "./graph/schema.graphqls")
	router.GET("/", playgroundHandler())

	// GraphQL Endpoint
	router.POST("/graphql", graphqlHandler())
	router.POST("/login", func(c *gin.Context) {
		auth.LoginHandler(c)
	})
	router.POST("/signup", func(c *gin.Context) {
		auth.SignupHandler(c)
	})
	router.GET("/auth/google", auth.HandleGoogleAuth)
	router.GET("/auth/google/callback", auth.HandleGoogleCallback)
	router.POST("/refresh", func(c *gin.Context) {
		auth.RefreshTokenHandler(c)
	})

	return router
}

func InitSentry(dsn string) bool {
	// Sentry Initialization
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
		return false
	}

	return true
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	if os.Getenv("ENV") == "development" {
		h.Use(extension.Introspection{})
	}

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func LoadConfig() Config {
	return Config{
		Port:             getEnv("PORT", defaultPort),
		Env:              os.Getenv("ENV"),
		DatabaseUser:     os.Getenv("POSTGRES_USER"),
		DatabasePassword: os.Getenv("POSTGRES_PASSWORD"),
		DatabaseName:     os.Getenv("POSTGRES_DB"),
		SentryDSN:        os.Getenv("SENTRY_BE_DSN"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
