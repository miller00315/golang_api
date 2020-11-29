package controllers

import (
	"api/src/authentication"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePublication
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.GetUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	var publication models.Publication

	publication.AuthorID = userID

	if error = publication.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = json.Unmarshal(requestBody, &publication); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	repository := repositories.NewPublicationRepository(db)

	publication.ID, error = repository.CreatePublication(publication)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, publication.ID)
}

// ListPublications
func ListPublications(w http.ResponseWriter, r *http.Request) {}

// GetPublication
func GetPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Pega todos os parâmetros da rota

	publicationID, error := strconv.ParseUint(params["publicationId"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := SetDatabase(w)

	if error != nil {
		return
	}

	defer db.Close()

	repository := repositories.NewPublicationRepository(db)

	publication, error := repository.GetPublication(publicationID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, publication)
}

// UpdatePublication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {}

// DeletePublication
func DeletePublication(w http.ResponseWriter, r *http.Request) {}
