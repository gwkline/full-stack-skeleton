package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwkline/full-stack-infra/backend/internal/database"
	"github.com/gwkline/full-stack-infra/backend/internal/graph/model"
)

func SignupHandler(c *gin.Context, database *database.Database) {
	var newUser model.NewUser

	// Bind JSON body to struct
	err := c.BindJSON(&newUser)
	if err != nil || newUser.Email == "" || newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	_, err = database.FindUser(newUser.Email, "email")
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	newUser.Password, err = hashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
		return
	}

	_, err = database.InsertUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// Generate JWT tokens
	accessToken, _ := generateToken(newUser.Email, AccessTokenDuration)
	refreshToken, _ := generateToken(newUser.Email, RefreshTokenDuration)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
