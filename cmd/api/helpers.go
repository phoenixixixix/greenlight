package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Helper method to reuse when I need to retrieve id from request params
// ! this method doesn't use any dependencies from app, but for consistency
// ! its good for setting up all handlers and helpers so that they are methods
// ! on application struct (+ in future I don't need to rewrite use places if
// ! if I want to use app dependencis in the method)
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	// The value returned by ByName() is always a string,
	//  so we try to convert it to abase 10 integer (with a bit size of 64).
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// Reuseable logic to form JSON response
func (app *application) writeJSON(w http.ResponseWriter, code int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data) // encode provided data to json
	if err != nil {
		return err
	}
	js = append(js, '\n')

	// Add all previous headers
	for k, v := range headers {
		w.Header()[k] = v // Header is map so we can set headers this way
	}
	// Add new headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// and write response
	w.Write([]byte(js))

	return nil
}
