package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Samaraferreira/go-api-rest/models"
	"github.com/Samaraferreira/go-api-rest/repositories"
	"github.com/gorilla/mux"
)

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
	var users = repositories.FindAll()
	json.NewEncoder(response).Encode(users)
}

func GetUserById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)
	user, err := repositories.FindUserById(idInt)

	if err != nil {
		errorForm := ErrorForm{
			Date:    time.Now().Format("02/01/2006 15:04:05"),
			Type:    "NotFound",
			Message: err.Error(),
		}
		json.NewEncoder(response).Encode(errorForm)
		return
	}

	json.NewEncoder(response).Encode(user)
}

func AddNewUser(response http.ResponseWriter, request *http.Request) {
	var newUser models.User
	json.NewDecoder(request.Body).Decode(&newUser)

	userResponse := repositories.AddNewUser(newUser)

	json.NewEncoder(response).Encode(userResponse)
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)
	user := repositories.DeleteUser(idInt)

	json.NewEncoder(response).Encode(user)
}

type ErrorForm struct {
	Date    string `json:"date"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
