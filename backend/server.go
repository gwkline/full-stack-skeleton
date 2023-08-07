package main

import (
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gwkline/full-stack-infra/backend/database"
	"github.com/gwkline/full-stack-infra/backend/graph"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8888"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	databaseUser := os.Getenv("POSTGRES_USER")
	databasePassword := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB")

	database.InitDB(databaseUser, databasePassword, databaseName)

	router := gin.Default()

	env := os.Getenv("ENV")
	if env == "production" {
		router.Use(cors.New(cors.Config{
			AllowOrigins:  []string{"http://localhost:8080"},
			AllowMethods:  []string{"OPTIONS", "POST", "GET", "PUT"},
			AllowHeaders:  []string{"*"},
			ExposeHeaders: []string{"Content-Length"},
		}))
	} else {
		router.Use(cors.Default())
	}

	router.POST("/graphql", graphqlHandler())
	router.GET("/", playgroundHandler())
	router.Run(":" + port)
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
