package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"database/sql"
	"encoding/json"
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

	db, error := setDatabase(w)

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

	db, error := setDatabase(w)

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

	db, error := setDatabase(w)

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

// UpdateUser atualiza um usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, error := strconv.ParseUint(params["userId"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
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

	db, error := setDatabase(w)

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

	db, error := setDatabase(w)

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

func setDatabase(w http.ResponseWriter) (*sql.DB, error) {

	db, error := database.Connect()

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return nil, error
	}

	return db, nil
}
