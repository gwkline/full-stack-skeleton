package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	introspectionfilter "github.com/ec2-software/gqlgen-introspect-filter"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/gwkline/full-stack-skeleton/backend/graph/generated"
	"github.com/gwkline/full-stack-skeleton/backend/graph/resolver"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (h *Handler) Graphql() gin.HandlerFunc {
	resolver := &resolver.Resolver{
		Repository: h.Repository,
		Redis:      h.Redis,
		Service:    h.Service,
	}
	c := generated.Config{Resolvers: resolver}
	c.Directives.HasRole = resolver.HasRole

	serv := handler.New(generated.NewExecutableSchema(c))

	serv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		if os.Getenv("ENV") == "production" {
			if !strings.Contains(e.Error(), "user not found in context") {
				sentry.CaptureException(e)
			}
		}
		return err
	})

	serv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		InitFunc:              h.webSocketInit,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	serv.AddTransport(transport.Options{})
	serv.AddTransport(transport.GET{})
	serv.AddTransport(transport.POST{})
	serv.AddTransport(transport.MultipartForm{})

	serv.SetQueryCache(lru.New(1000))

	serv.Use(extension.Introspection{})
	serv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	serv.Use(extension.FixedComplexityLimit(1500))
	serv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		return h.introspectionMiddleware(ctx, next)
	})
	serv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		return h.maintainanceMiddleware(ctx, next)
	})
	serv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		return h.dbContextMiddleware(ctx, next)
	})

	plugin := introspectionfilter.Plugin{
		FieldFilter: func(ctx context.Context, d *ast.FieldDefinition) bool {
			for _, directive := range d.Directives {
				for _, arg := range directive.Arguments {
					if arg.Name == "role" && arg.Value.Raw == "ADMIN" && os.Getenv("ENV") != "development" {
						user, err := h.Service.Auth.CurrentUser(ctx)
						if err != nil || !user.IsAdmin() {
							return false
						}
					}
				}
			}
			return true
		},
	}

	serv.Use(&plugin)

	return func(c *gin.Context) {
		serv.ServeHTTP(c.Writer, c.Request)
	}
}

func (h *Handler) introspectionMiddleware(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	opContext := graphql.GetOperationContext(ctx)
	tx := newrelic.FromContext(ctx)
	tx.SetName(fmt.Sprintf("%s::%s", opContext.Operation.Operation, opContext.Operation.Name))
	tx.AddAttribute("variables", opContext.Variables)
	if opContext.OperationName == "IntrospectionQuery" {
		_, err := h.Service.Auth.CurrentUser(ctx)
		if err != nil {
			return graphql.OneShot(graphql.ErrorResponse(ctx, "Unauthorized"))
		}
	}

	if opContext.Operation.Name == "queueSubscription" {
		tx.Ignore()
	}
	return next(ctx)
}
func (h *Handler) dbContextMiddleware(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	h.Repository.DB = h.Repository.DB.WithContext(ctx)
	return next(ctx)
}

func (h *Handler) maintainanceMiddleware(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	if os.Getenv("MAINTAINANCE_MODE") == "true" {
		return graphql.OneShot(graphql.ErrorResponse(ctx, "The app is currently in maintainance mode. Please try again later."))
	}
	return next(ctx)
}

func (h *Handler) webSocketInit(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
	any := initPayload["Authorization"]
	token, ok := any.(string)
	if !ok {
		return ctx, &initPayload, nil
	}
	token = strings.ReplaceAll(token, "Bearer", "")
	token = strings.Trim(token, " ")

	ctx, err := h.authenticateUser(ctx, token)
	if err != nil {
		return ctx, &initPayload, nil
	}

	return ctx, &initPayload, nil
}

func (h *Handler) authenticateUser(ctx context.Context, token string) (context.Context, error) {
	if token == "" {
		return ctx, fmt.Errorf("user not found")
	}

	claims, err := h.Service.Auth.ValidateToken(ctx, token)
	if err != nil {
		return ctx, err
	}

	email := claims.Email
	user, err := h.Repository.User.FindBy([]types.Filter{{Key: "email", Value: email}})
	if err != nil {
		return ctx, err
	}

	ctx = h.Service.Auth.SetCurrentUser(ctx, user)

	return ctx, nil
}
