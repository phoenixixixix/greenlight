package main

import (
	"fmt"
	"net/http"
)

// Errors logic for whole main package.

// Mehod that form outputs error in console
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// Method that output errors in JSON
func (app *application) errorResponce(w http.ResponseWriter, r *http.Request, code int, message interface{}) {
	env := envelope{"error": message}
	err := app.writeJSON(w, code, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Below are methods, whose names correspond to specific error case

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponce(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the request resource could not be found"
	app.errorResponce(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponce(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponce(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponce(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponce(w, r, http.StatusUnprocessableEntity, errors)
}
