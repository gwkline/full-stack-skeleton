package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gwkline/full-stack-infra/backend/database"
	"github.com/gwkline/full-stack-infra/backend/graph"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8888"

type Config struct {
	Port, Env, DatabaseUser, DatabasePassword, DatabaseName string
}

func LoadConfig() Config {
	return Config{
		Port:             getEnv("PORT", defaultPort),
		Env:              os.Getenv("ENV"),
		DatabaseUser:     os.Getenv("POSTGRES_USER"),
		DatabasePassword: os.Getenv("POSTGRES_PASSWORD"),
		DatabaseName:     os.Getenv("POSTGRES_DB"),
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

	if env == "production" {
		router.Use(cors.New(cors.Config{
			AllowOrigins:  []string{"http://localhost:8080"},
			AllowMethods:  []string{"OPTIONS", "POST", "GET", "PUT", "DELETE"},
			AllowHeaders:  []string{"*"},
			ExposeHeaders: []string{"Content-Length"},
		}))
	} else if env == "development" {
		router.Use(cors.Default())
	} else {
		router.Use(cors.Default())
	}

	router.OPTIONS("/graphql", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	router.POST("/graphql", graphqlHandler())
	router.GET("/", playgroundHandler())
	return router
}

func main() {
	cfg := LoadConfig()

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
