package repo

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gwkline/full-stack-skeleton/backend/types"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx5"
	"github.com/rafaelhl/gorm-newrelic-telemetry-plugin/telemetry"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	maxRetries    = 10
	retryInterval = 5 * time.Second
)

func Init() (*Repository, error) {
	psqlInfo := os.Getenv("DATABASE_URL")
	if psqlInfo == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	if os.Getenv("ENV") == "development" {
		fmt.Println(psqlInfo)
	} else {
		psqlInfo = psqlInfo + "?sslmode=require"
	}

	fmt.Println("Waiting for database connection")

	gormLogger := createLogger()

	for i := 0; i < maxRetries; i++ {
		dbConn, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{Logger: gormLogger})
		if err != nil {
			logRetryAttempt(i)
			time.Sleep(retryInterval)
			continue
		}

		if os.Getenv("ENV") == "production" {
			err = dbConn.Use(telemetry.NewNrTracer("Main", "Heroku", "Postgres"))
			if err != nil {
				return nil, fmt.Errorf("failed to use newrelic tracer: %w", err)
			}
		}

		fmt.Println("Connected to the database")
		return setupDatabase(dbConn)
	}

	return nil, fmt.Errorf("failed to connect to database after %d tries", maxRetries)
}

func createLogger() logger.Interface {
	var loggerLev logger.LogLevel
	switch os.Getenv("ENV") {
	case "development":
		// loggerLev = logger.Info
		loggerLev = logger.Error
	default:
		loggerLev = logger.Error
	}
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  loggerLev,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
}

func setupDatabase(dbConn *gorm.DB) (*Repository, error) {
	x, err := dbConn.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	x.SetMaxOpenConns(110)
	x.SetMaxIdleConns(10)
	db := NewRepository(dbConn)
	models := getModels()

	dbConn.Exec("DO $$ BEGIN CREATE TYPE action AS ENUM ('drip', 'stop', 'fail', 'send'); EXCEPTION WHEN duplicate_object THEN END $$;")
	dbConn.Exec("DO $$ BEGIN CREATE TYPE status AS ENUM ('pending', 'queued', 'working', 'dripping', 'completed', 'failed', 'stopped'); EXCEPTION WHEN duplicate_object THEN END $$;")
	dbConn.Exec("DO $$ BEGIN CREATE TYPE new_or_reply AS ENUM ('NewEmail', 'Reply'); EXCEPTION WHEN duplicate_object THEN END $$;")

	for _, model := range models {
		err := dbConn.AutoMigrate(model)
		if err != nil {
			return nil, fmt.Errorf("failed migrating DB model '%T': %w", model, err)
		}
	}

	return db, nil
}

func getModels() []interface{} {
	return []interface{}{
		&types.User{},
		&types.Apple{},
	}
}

func logRetryAttempt(i int) {
	fmt.Printf("Failed to connect to r.db. Retry %d/%d. Waiting for %v before retrying...", i+1, maxRetries, retryInterval)
}
