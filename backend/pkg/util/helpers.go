package util

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/badoux/checkmail"
	"github.com/gwkline/full-stack-skeleton/backend/mocks"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/hibiken/asynq"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ValidateEmail(email string) error {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return fmt.Errorf("failed validating email format: %w", err)
	}

	err = checkmail.ValidateHost(email)
	if err != nil {
		return fmt.Errorf("failed validating email host: %w", err)
	}

	return nil
}

func AddBusinessDays(daysToAdd int) string {
	// Get the current time
	currentTime := time.Now()

	// Add the given number of days
	futureTime := currentTime.AddDate(0, 0, daysToAdd)

	// Check if the resulting day is a weekend
	switch futureTime.Weekday() {
	case time.Saturday:
		futureTime = futureTime.AddDate(0, 0, 2) // Move to Monday
	case time.Sunday:
		futureTime = futureTime.AddDate(0, 0, 1) // Move to Monday
	}

	// Return as a string in "yyyy-MM-dd" format
	return futureTime.Format("2006-01-02")
}

func GetEnvWithFallback(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func NewAsynqOps() asynq.RedisClientOpt {
	ops, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		fmt.Printf("failed parsing redis: %v", err)
	}

	return asynq.RedisClientOpt{
		Addr:     ops.Addr,
		DB:       ops.DB,
		Password: ops.Password,
		Username: ops.Username,
	}
}

func MockDB() (*repo.Repository, sqlmock.Sqlmock) {
	sql, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Failed to create sqlmock: %s", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sql,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to create sqlmock: %s", err)
	}

	mockDatabase := repo.NewRepository(gormDB)
	return mockDatabase, mock
}

func ValToPtr[T any](val T) *T {
	return &val
}

func NewMockRepository(t *testing.T) *repo.Repository {
	return &repo.Repository{
		User:  mocks.NewIGenericRepo[types.User](t),
		Apple: mocks.NewIGenericRepo[types.Apple](t),
	}
}

func NewGoroutineContext(ctx context.Context) context.Context {
	txn := newrelic.FromContext(ctx)
	return newrelic.NewContext(ctx, txn.NewGoroutine())
}
