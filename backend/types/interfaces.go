package types

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type Services struct {
	Auth           IAuth
	Workers        IWorker
	Queue          IQueue
	Search         ISearch
	ExampleService ISomeService
}

type ISearch interface {
	Apples(ctx context.Context, input *ConnectionInput) (*AppleConnection, error)
}

type ISomeService interface {
	SomeMethod(ctx context.Context) error
}

type IQueue interface {
	Clear(ctx context.Context, name string) ([]*asynq.TaskInfo, error)
	TogglePause(ctx context.Context, name string) error
	StartServer(ctx context.Context, workPkg IWorker)
	Get(ctx context.Context, name string) (*asynq.QueueInfo, error)
	Enqueue(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
	RunTask(ctx context.Context, queueName string, taskId string) error
	DeleteTask(ctx context.Context, queueName string, taskId string) error
	ArchiveTask(ctx context.Context, queueName string, taskId string) error
}

type IAuth interface {
	Login(ctx context.Context, li LoginInput) (*JWT, error)
	Signup(ctx context.Context, newuser User, password string) (*JWT, error)
	RefreshToken(ctx context.Context, jwt JWT) (*JWT, error)
	GenerateToken(ctx context.Context, email string, duration time.Duration) (string, error)
	ValidateToken(ctx context.Context, tokenStr string) (*Claims, error)
	CurrentUser(ctx context.Context) (*User, error)
	SetCurrentUser(ctx context.Context, user *User) context.Context
}

type IWorker interface {
	HandleProcessSomeDataTask(ctx context.Context, t *asynq.Task) error
}

type HandlerDef interface {
	Graphql(repo any) gin.HandlerFunc
	Playground() gin.HandlerFunc
	SomeWebhookHandler(c *gin.Context, e IWorker)
}

type IGenericRepo[T any] interface {
	Create(entity *T) (*T, error)
	Update(entity *T) (*T, error)
	Archive(entity *T) error // soft delete
	Delete(entity *T) error  // hard delete
	ListBy(filters []Filter, preloads ...string) ([]*T, error)
	FindBy(filters []Filter, preloads ...string) (*T, error)
	FuzzyFindBy(key string, value any) ([]*T, error)
	Connection(tx *gorm.DB, tableName string, filter []FilterInput, sortBy string, direction AscOrDesc, limit int, after int) ([]*T, *PageInfo, error)
}

type Modeler interface {
	GetID() uint
}
