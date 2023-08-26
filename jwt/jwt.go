package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("secret")

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	// Set token claims
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with our secret
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	} else {
		return "", err
	}
}
