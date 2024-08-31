package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang_task_tracker/config"
	"log"
	"strconv"
	"time"
)

func NewJWT(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject":   strconv.Itoa(userID),
		"expiresAt": time.Now().Add(time.Second * time.Duration(config.Envs.JWTExpiration)).Unix(),
	})

	signedToken, err := token.SignedString([]byte(config.Envs.JWTSecret))

	if err != nil {
		log.Fatal(err)

		return "", err
	}

	return signedToken, nil
}

//func AuthWithJWT(token string) error {
//
//}

func validateJWT(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}
