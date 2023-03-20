package controller

import "net/http"

func FindAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FindAllPosts"))
}

func FindPostByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FindPostByID"))
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreatePost"))
}

func DeletePostByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeletePostByID"))
}

func UpdatePostByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePostByID"))
}
