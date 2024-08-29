package auth

import (
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
