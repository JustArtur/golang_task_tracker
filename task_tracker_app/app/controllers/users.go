package controllers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"task_tracker_app/app/helpers"
	"task_tracker_app/app/models"
	"task_tracker_app/app/services/auth"
	"task_tracker_app/app/types"
)

var user types.UserPayload

func Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	err := helpers.ParseRequest(r, &user)
	if err != nil {
		log.Println(err)
		helpers.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	userRecord, err := models.GetUserByEmail(user.Email)
	if err != nil {
		log.Println(err)
		helpers.SendErrorResponse(w, http.StatusBadRequest, fmt.Errorf("invalid user email or password"))
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userRecord.Password), []byte(user.Password)); err != nil {
		log.Println(err)
		helpers.SendErrorResponse(w, http.StatusBadRequest, fmt.Errorf("invalid user email or password"))
		return
	}

	token, err := auth.NewJWT(userRecord.ID)
	if err != nil {
		log.Println(err)
		helpers.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	helpers.SendResponse(w, http.StatusOK, map[string]string{"access_token": token})
}

func Register(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	err := helpers.ParseRequest(r, &user)
	if err != nil {
		helpers.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	result, _ := models.GetUserByEmail(user.Email)
	if result != nil {
		helpers.SendErrorResponse(w, http.StatusBadRequest, fmt.Errorf("user with this email already exists"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	err = models.CreateUser(types.User{
		Email:    user.Email,
		Password: string(hashedPassword),
		Name:     user.Name,
	})
	if err != nil {
		helpers.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	helpers.SendResponse(w, http.StatusCreated, nil)
}
