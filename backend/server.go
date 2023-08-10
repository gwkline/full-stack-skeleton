package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

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
	router.StaticFile("/schema.graphqls", "./graph/schema.graphqls")
	router.POST("/graphql", graphqlHandler())
	router.GET("/", playgroundHandler())

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

func signupLoginHandler(c *gin.Context) {
	var login auth.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle Signup
	if login.Password == "" {
		// Add your signup logic here
		// Save the username to your database

		c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
		return
	}

	// Handle Login
	// Validate username and password with your database

	// If 2FA code is provided, validate it

	// Generate JWT tokens
	accessToken, _ := auth.GenerateToken(login.Username, auth.AccessTokenDuration)
	refreshToken, _ := auth.GenerateToken(login.Username, auth.RefreshTokenDuration)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func refreshTokenHandler(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshToken, ok := data["refresh_token"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token required"})
		return
	}

	claims, err := auth.ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired"})
		return
	}

	newAccessToken, _ := auth.GenerateToken(claims.Username, auth.AccessTokenDuration)
	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
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
