package utils

import(
	"github.com/dgrijalva/jwt-go"
	"time"
	"os"
)

type Claims struct {
	Username string `json:username`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
			Issuer: "go-fiber-mvc-app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}