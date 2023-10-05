package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
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
