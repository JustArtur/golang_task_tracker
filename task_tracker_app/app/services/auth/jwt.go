package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strconv"
	"task_tracker_app/app/helpers"
	"task_tracker_app/app/models"
	"task_tracker_app/config"
	"time"
)

func NewJWT(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(config.Envs.JWTExpiration))),
	})

	signedToken, err := token.SignedString([]byte(config.Envs.JWTSecret))

	if err != nil {
		log.Fatal(err)

		return "", err
	}

	return signedToken, nil
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := helpers.GetTokenFromRequest(r)

		token, err := validateJWT(tokenStr)
		if err != nil {
			log.Printf("validate token error: %v", err)
			unauthorized(w)
			return
		}

		userID, err := getUserIDFromJWT(token)
		if err != nil {
			log.Printf("failed to get user id from token: %v", err)
			unauthorized(w)
			return
		}

		u, err := models.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			unauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", u.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateJWT(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func getUserIDFromJWT(token *jwt.Token) (int, error) {
	claims := token.Claims.(jwt.MapClaims)

	userIDStr, err := claims.GetSubject()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func unauthorized(w http.ResponseWriter) {
	helpers.SendErrorResponse(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
}
