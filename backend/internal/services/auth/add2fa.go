package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwkline/full-stack-infra/backend/internal/database"
)

func Add2FA(c *gin.Context, database *database.Database) {
	var login Login

	err := c.BindJSON(&login)
	if err != nil || login.Email == "" || login.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	user, err := database.FindUser(login.Email, "email")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if user.OtpSecret != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "2FA already enabled"})
		return
	}

	key, _ := generateTOTPKey(user.Email)
	secret := key.Secret()
	user.OtpSecret = &secret

	_, err = database.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set key on user"})
	}

	c.JSON(http.StatusOK, key.URL())
}
