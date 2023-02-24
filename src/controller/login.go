package controller

import (
	"devbook-api/src/auth"
	"devbook-api/src/database"
	"devbook-api/src/repository"
	"devbook-api/src/response"
	"devbook-api/src/security"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	userFound, error := repository.NewUserRepository(db).FindByEmail(user.Email)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.Compare(userFound.Password, user.Password); error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	token, error := auth.GetToken(userFound.Id)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusOK, token)
}
