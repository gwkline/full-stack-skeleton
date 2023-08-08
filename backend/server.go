package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gwkline/full-stack-infra/backend/database"
	"github.com/gwkline/full-stack-infra/backend/graph"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8888"

type Config struct {
	Port, Env, DatabaseUser, DatabasePassword, DatabaseName, SentryDSN string
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

func SetupRouter(env string) *gin.Engine {
	router := gin.Default()

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
	router.StaticFile("/schema.graphqls", "./graph/schema.graphqls")
	router.POST("/graphql", graphqlHandler())
	router.GET("/", playgroundHandler())

	return router
}

func main() {
	cfg := LoadConfig()

	// Sentry Initialization
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.SentryDSN,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}

	if cfg.Env == "production" || cfg.Env == "development" {
		database.InitDB(cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName)
	}

	router := SetupRouter(cfg.Env)
	router.Run(":" + cfg.Port)
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

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
