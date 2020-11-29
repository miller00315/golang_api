package controllers

import (
	"api/src/database"
	"api/src/responses"
	"database/sql"
	"net/http"
)

func SetDatabase(w http.ResponseWriter) (*sql.DB, error) {

	db, error := database.Connect()

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return nil, error
	}

	return db, nil
}