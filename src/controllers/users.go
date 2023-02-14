package controllers

import "net/http"

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FindAllUsers"))
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FindUserByID"))
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUserByID"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateUser"))
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}
