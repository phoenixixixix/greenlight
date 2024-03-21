package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

// Wrapper for all models
type Models struct {
	Movies MovieModel
}

// Convenient function to initialize all models
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
