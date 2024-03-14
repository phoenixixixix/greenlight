package main

import (
	"fmt"
	"net/http"
)

// Errors logic for whole main package.
// Methods that form output in console and JSON response

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) errorResponce(w http.ResponseWriter, r *http.Request, code int, message interface{}) {
	env := envelope{"error": message}
	err := app.writeJSON(w, code, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

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
