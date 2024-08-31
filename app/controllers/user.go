package controllers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang_task_tracker/app/models"
	"golang_task_tracker/app/services/auth"
	"golang_task_tracker/app/types"
	"log"
	"net/http"
)

var user types.UserPayload

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	err := ParseRequest(r, &user)

	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	userRecord, err := models.GetUserByEmail(user.Email)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusBadRequest, fmt.Errorf("invalid user email or password"))
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userRecord.Password), []byte(user.Password)); err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusBadRequest, fmt.Errorf("invalid user email or password"))
		return
	}

	token, err := auth.NewJWT(userRecord.ID)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	SendResponse(w, http.StatusOK, map[string]string{"access_token": token})
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	err := ParseRequest(r, &user)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	result, _ := models.GetUserByEmail(user.Email)
	if result != nil {
		SendErrorResponse(w, http.StatusBadRequest, fmt.Errorf("user with this email already exists"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	err = models.CreateUser(types.User{
		Email:    user.Email,
		Password: string(hashedPassword),
		Name:     user.Name,
	})
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	SendResponse(w, http.StatusCreated, nil)
}
