package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/phoenixixixix/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "createMoview Handler")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Dune",
		Runtime:   148,
		Genres:    []string{"action", "mistery"},
		Version:   1,
	}

	// envelop is custom type, defined to envelop data under top-level key name
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem.", http.StatusInternalServerError)
	}
}
