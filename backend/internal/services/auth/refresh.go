package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gwkline/full-stack-infra/backend/internal/database"
)

func RefreshTokenHandler(c *gin.Context, database *database.Database) {
	var data JWT
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	claims, err := validateToken(data.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims2, _ := validateToken(data.AccessToken)
	if claims.Email != claims2.Email {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired"})
		return
	}

	newAccessToken, _ := generateToken(claims.Email, AccessTokenDuration)
	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}
