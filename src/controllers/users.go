package controllers

import "net/http"

// CreateUser cria um usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create user"))
}

// GetUsers recupera todos os ussuários
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get users"))
}

// GetUser recupera um usuário
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

// UpdateUser atualiza um usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("udate user"))
}

// DeleteUser apaga um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
