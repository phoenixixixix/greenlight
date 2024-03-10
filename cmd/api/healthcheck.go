package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Making json like string, because in the end of the day json is just a text
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version) // papulating interpolated values

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}
