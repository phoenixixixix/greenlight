package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// data to be encoded in json (using custom envelope type to wrap data under comon key)
	env := envelope{
		"status": "available",
		"sys_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Something went wrong. Can't process your request", http.StatusInternalServerError)
	}
}
