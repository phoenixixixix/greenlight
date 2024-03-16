package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// ! Some methods here doesn't use any dependencies from app, but for consistency
// ! its good for setting up all handlers and helpers so that they are methods
// ! on application struct (+ in future I don't need to rewrite use places
// ! if I want to use app dependencis in the method)

// Helper method to reuse when I need to retrieve id from request params
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

// envelop is custom type, defined to wrap data under (common) top-level key name
type envelope map[string]interface{}

// Reuseable logic to form JSON response
func (app *application) writeJSON(w http.ResponseWriter, code int, data envelope, headers http.Header) error {
	// encode provided data to json Marshal(), with indentation MarshalIndent()
	// MarshalIndent() is much more slower that Marshl()
	js, err := json.MarshalIndent(data, "", "\t")
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

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.As(err, &invalidUnmarshalError):
			// It's good practice to panic some times: in this case when this error ocures
			// it's not because of bad user input raser it's developer mistake on providing
			// Destination in json.Decode function. So it's better to panic and catch this
			// error on early (dev) stage.
			panic(err)
		default:
			return err
		}
	}

	return nil
}
