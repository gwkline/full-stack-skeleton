package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwkline/full-stack-infra/backend/internal/database"
)

func LoginHandler(c *gin.Context, db *database.Database) {
	var login Login

	// Bind JSON body to struct
	err := c.BindJSON(&login)
	if err != nil || login.Email == "" || login.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	user, err := db.FindUser(login.Email, "email")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if user.OtpSecret != nil {
		if !validOtpCode(*user.OtpSecret, login.OTP) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
			return
		}
	}

	if !validPassword(login.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT tokens
	accessToken, _ := generateToken(login.Email, AccessTokenDuration)
	refreshToken, _ := generateToken(login.Email, RefreshTokenDuration)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
