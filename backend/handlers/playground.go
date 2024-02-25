package handlers

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Playground() gin.HandlerFunc {
	serv := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		serv.ServeHTTP(c.Writer, c.Request)
	}
}
