package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	jwtSecret            = []byte(os.Getenv("JWT_SECRET")) // Change this to a secure key
	AccessTokenDuration  = time.Hour * 2                   // Token valid for 2 hours
	RefreshTokenDuration = time.Hour * 24 * 7              // Refresh token valid for 7 days
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	OTP      string `json:"otp"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string, duration time.Duration) (string, error) {
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
