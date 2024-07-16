package middleware

import (
	"fmt"
	"log"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func VerifyToken(tokenString string) (*jwt.Token, error) {
	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mySigningKey := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, nil
	} else {
		return token, err
	}
}
