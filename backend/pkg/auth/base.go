package auth

import (
	"os"
	"time"

	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
)

var (
	jwtSecret            = []byte(os.Getenv("JWT_SECRET")) // TODO: Change this to a secure key
	AccessTokenDuration  = time.Hour * 24 * 30             // Token valid for 1 month
	RefreshTokenDuration = time.Hour * 24 * 365            // Refresh token valid for 1 year
)

type Service struct {
	Repository *repo.Repository
}

func Init(repo *repo.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
