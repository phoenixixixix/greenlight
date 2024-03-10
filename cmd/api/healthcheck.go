package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// data to be encoded in json
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Something went wrong. Can't process your request", http.StatusInternalServerError)
	}
}
