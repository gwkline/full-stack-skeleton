package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gwkline/full-stack-infra/backend/internal/database"
	"github.com/gwkline/full-stack-infra/backend/internal/graph/model"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtSecret            = []byte(os.Getenv("JWT_SECRET")) // Change this to a secure key
	AccessTokenDuration  = time.Hour * 2                   // Token valid for 2 hours
	RefreshTokenDuration = time.Hour * 24 * 7              // Refresh token valid for 7 days
	// googleOauthConfig    = oauth2.Config{
	// 	RedirectURL:  "http://localhost:8888/auth/google/callback", // Adjust this
	// 	ClientID:     "YOUR_GOOGLE_CLIENT_ID",
	// 	ClientSecret: "YOUR_GOOGLE_CLIENT_SECRET",
	// 	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	// 	Endpoint:     google.Endpoint,
	// }
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating user"})
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

	key, err := generateTOTPKey(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate TOTP key"})
	}

	secret := key.Secret()
	user.OtpSecret = &secret

	_, err = database.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set key on user"})
	}

	c.JSON(http.StatusOK, key.URL())
}

func RefreshTokenHandler(c *gin.Context, database *database.Database) {
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

	claims, err := validateToken(refreshToken)
	if err != nil {
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

func generateToken(email string, duration time.Duration) (string, error) {
	claims := Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func validateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func generateTOTPKey(email string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      "YourAppName",
		AccountName: email,
	})
}

func validOtpCode(userKey string, token string) bool {
	return totp.Validate(token, userKey)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func HandleGoogleAuth(c *gin.Context) {
// 	url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
// 	c.Redirect(http.StatusTemporaryRedirect, url)
// }

// func HandleGoogleCallback(c *gin.Context) {
// 	token, err := googleOauthConfig.Exchange(c, c.DefaultQuery("code", ""))
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "Could not get token")
// 		return
// 	}
// 	fmt.Println(token)
// 	// Use token to get user info, then authenticate the user in your system.
// 	// ...
// }
