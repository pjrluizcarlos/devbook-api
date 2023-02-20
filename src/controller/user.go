package controller

import (
	"devbook-api/src/database"
	"devbook-api/src/model"
	"devbook-api/src/repository"
	"devbook-api/src/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	users, error := repository.NewUser(db).FindAll()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusCreated, users)
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id, error := strconv.ParseUint(parameters["id"], 10, 64)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	user, error := repository.NewUser(db).FindById(id)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	if user.Id == 0 {
		response.Error(w, http.StatusNotFound, fmt.Errorf("User not found with ID [%d]", id))
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUserByID"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, error := requestBody(r)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	newId, error := repository.NewUser(db).Create(user)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	user.Id = newId

	response.JSON(w, http.StatusCreated, user)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

func requestBody(r *http.Request) (model.User, error) {
	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		return model.User{}, error
	}

	var user model.User

	if error = json.Unmarshal(requestBody, &user); error != nil {
		return model.User{}, error
	}

	return user, nil
}
