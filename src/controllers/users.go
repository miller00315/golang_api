package controllers

import (
	"api/src/authentication"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser cria um usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)

	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User

	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("register"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	// Instancia um novo repositorio através de um generator
	repository := repositories.NewUserRepository(db)

	user.ID, error = repository.Create(user)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// GetUsers recupera todos os ussuários
func GetUsers(w http.ResponseWriter, r *http.Request) {
	userQuery := strings.ToLower(r.URL.Query().Get("user"))

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, error := repository.Search(userQuery)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// GetUser recupera um usuário
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Pega todos os parâmetros da rota

	userID, error := strconv.ParseUint(params["userId"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	user, error := repository.Get(userID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["userId"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserID, error := authentication.GetUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("Its not possible update others user data"))
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	var user models.User

	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("update"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	if error = repository.Update(userID, user); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser apaga um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["userId"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserID, error := authentication.GetUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("Its not possible delete others users"))
		return
	}

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	if error = repository.Delete(userID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser permits follow a user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, error := authentication.GetUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["userID"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("Its not valid follow yourself"))
		return
	}

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	if error := repository.Follow(userID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// UnfollowUser ppermits unfollow a user
func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, error := authentication.GetUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["userID"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("Its not valid unfollow yourself"))
		return
	}

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	if error := repository.UnFollow(userID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// GetFollowers get all folowers of a user
func GetFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["userID"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	followers, error := repository.GetFollowers(userID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, followers)

}

// GetFollowing get all following users
func GetFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["userID"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	following, error := repository.GetFollowing(userID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, following)
}

// UpdatePassword update user password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	tokenUserID, error := authentication.GetUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["userID"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if tokenUserID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("Cannot alter password of others users"))
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	var password models.Password

	if error = json.Unmarshal(requestBody, &password); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	dbPassword, error := repository.GetPassword(userID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.CheckPassword(dbPassword, password.CurrentPassword); error != nil {
		responses.Error(w, http.StatusForbidden, errors.New("The password do not exist in database"))
		return
	}

	hashPassword, error := security.Hash(password.NewPassword)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = repository.UpdatePassword(string(hashPassword), userID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
