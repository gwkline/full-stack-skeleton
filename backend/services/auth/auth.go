package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gwkline/full-stack-infra/backend/database"
	"github.com/gwkline/full-stack-infra/backend/graph/model"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	jwtSecret            = []byte(os.Getenv("JWT_SECRET")) // Change this to a secure key
	AccessTokenDuration  = time.Hour * 2                   // Token valid for 2 hours
	RefreshTokenDuration = time.Hour * 24 * 7              // Refresh token valid for 7 days
	googleOauthConfig    = oauth2.Config{
		RedirectURL:  "http://localhost:8888/auth/google/callback", // Adjust this
		ClientID:     "YOUR_GOOGLE_CLIENT_ID",
		ClientSecret: "YOUR_GOOGLE_CLIENT_SECRET",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	OTP      string `json:"otp"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func LoginHandler(c *gin.Context) {
	var login Login

	// Bind JSON body to struct
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	fmt.Println(login) // { }

	// User has just entered the email field - we need to check if the account exists
	// to determine the user flow
	if login.Email == "" || login.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must pass the email key without a password"})
		return
	}

	user, err := database.FindUser(login.Email, "email")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if login.Password != "" && login.OTP == "" {
		if !CheckPasswordHash(login.Password, user.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
			return
		}
	} else if login.OTP != "" {
		if !ValidateTOTPToken(*user.Otp, login.OTP) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
			return
		}
	}

	// Generate JWT tokens
	accessToken, _ := GenerateToken(login.Email, AccessTokenDuration)
	refreshToken, _ := GenerateToken(login.Email, RefreshTokenDuration)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          user,
	})
}

func SignupHandler(c *gin.Context) {
	var login model.NewUser

	// Bind JSON body to struct
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	fmt.Println(login)

	if login.Email == "" || login.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must pass an email and password in your request body as non-empty strings."})
		return
	}

	login.Password, err = HashPassword(login.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
		return
	}

	_, err = database.InsertUser(login)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating user"})
		return
	}

	// Generate JWT tokens
	accessToken, _ := GenerateToken(login.Email, AccessTokenDuration)
	refreshToken, _ := GenerateToken(login.Email, RefreshTokenDuration)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshTokenHandler(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshToken, ok := data["refresh_token"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token required"})
		return
	}

	claims, err := ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired"})
		return
	}

	newAccessToken, _ := GenerateToken(claims.Email, AccessTokenDuration)
	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}

func GenerateToken(email string, duration time.Duration) (string, error) {
	claims := Claims{
		Email: email,
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

func HandleGoogleAuth(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGoogleCallback(c *gin.Context) {
	token, err := googleOauthConfig.Exchange(c, c.DefaultQuery("code", ""))
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Could not get token")
		return
	}
	fmt.Println(token)
	// Use token to get user info, then authenticate the user in your system.
	// ...
}

// Generate a new TOTP Key for a user
func GenerateTOTPKey(email string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      "YourAppName",
		AccountName: email,
	})
}

// Validate a provided TOTP token for a user
func ValidateTOTPToken(userKey string, token string) bool {
	return totp.Validate(token, userKey)
}

// HashPassword hashes a password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword validates a password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
