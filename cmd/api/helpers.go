package main

import (
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
