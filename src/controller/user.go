package controller

import (
	"devbook-api/src/auth"
	"devbook-api/src/database"
	"devbook-api/src/model"
	"devbook-api/src/repository"
	"devbook-api/src/response"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := r.URL.Query().Get("user")

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	users, error := repository.NewUserRepository(db).FindAll(nameOrNick)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	id, error := NewPathVariable(r).uint64("id")
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

	user, error := repository.NewUserRepository(db).FindById(id)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	if user.Id == 0 {
		response.Error(w, http.StatusNotFound, fmt.Errorf("user not found with ID [%d]", id))
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	id, error := NewPathVariable(r).uint64("id")
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if error := isSameUserFromAuthorization(id, r); error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	if error = repository.NewUserRepository(db).DeleteById(id); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, error := requestBody(r)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if error := user.Prepare(); error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	newId, error := repository.NewUserRepository(db).Create(user)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	user.Id = newId

	response.JSON(w, http.StatusCreated, user)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	user, error := requestBody(r)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if error := user.Prepare(); error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	id, error := NewPathVariable(r).uint64("id")
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if error := isSameUserFromAuthorization(id, r); error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	user.Id = id

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	if error := repository.NewUserRepository(db).Update(user); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	id, error := NewPathVariable(r).uint64("id")
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserId, error := auth.GetUserId(GetAuthorizationHeader(r))
	if error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	if id == tokenUserId {
		response.Error(w, http.StatusForbidden, errors.New("an user cannot be followed by himself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	if error := repository.NewUserRepository(db).Follow(id, tokenUserId); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	id, error := NewPathVariable(r).uint64("id")
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserId, error := auth.GetUserId(GetAuthorizationHeader(r))
	if error != nil {
		response.Error(w, http.StatusUnauthorized, error)
		return
	}

	if id == tokenUserId {
		response.Error(w, http.StatusForbidden, errors.New("an user cannot be unfollowed by himself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	if error := repository.NewUserRepository(db).Unfollow(id, tokenUserId); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func FindAllFollowersById(w http.ResponseWriter, r *http.Request) {
	id, error := NewPathVariable(r).uint64("id")
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

	followers, error := repository.NewUserRepository(db).FindAllFollowersById(id)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func isSameUserFromAuthorization(userId uint64, r *http.Request) error {
	tokenUserId, error := auth.GetUserId(GetAuthorizationHeader(r))
	if error != nil {
		return error
	}

	if tokenUserId != userId {
		return errors.New("authorized user is not the same that is being changed")
	}

	return nil
}

func requestBody(r *http.Request) (model.User, error) {
	requestBody, error := getRequestBody(r)
	if error != nil {
		return model.User{}, error
	}

	var user model.User

	if error = json.Unmarshal(requestBody, &user); error != nil {
		return model.User{}, error
	}

	return user, nil
}
