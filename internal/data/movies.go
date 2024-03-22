package data

import (
	"database/sql"
	"time"

	"github.com/lib/pq"

	"github.com/phoenixixixix/greenlight/internal/validator"
)

type Movie struct {
	// 1. Fields should start with capital letter to be exported.
	//    This way they are visible to encoding/json packege
	// 2. `json:"field"` is a Struct Tag which defines how fields
	//    should be parsed (to json in this case)
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // - (hyphen) copletly hide this field (ex. I don't want to show this field to end user)
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"` // omitempty hides field if it has its zero value
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

// Validation of required fields to prevent mistakes in user input
func ValidateMovie(v *validator.Validator, m *Movie) {
	v.Check(m.Title != "", "title", "must be provided")
	v.Check(len(m.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(m.Year != 0, "year", "must be provided")
	v.Check(m.Year >= 1888, "year", "must be greater than 1888")
	v.Check(m.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(m.Runtime != 0, "runtime", "must be provided")
	v.Check(m.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(m.Genres != nil, "genres", "must be provided")
	v.Check(len(m.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(m.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(m.Genres), "genres", "must not conain duplicate values")
}

type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(movie *Movie) error {
	query := `
  INSERT INTO movies(title, year, runtime, genres)
  VALUES ($1, $2, $3, $4)
  RETURNING id, created_at, version`

	// It's optional to daclare this slice but it makes nice and clear what values
	// should be in placeholders
	args := []interface{}{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	// retrurning error from Scan if any and simultaneously Scan populates system-generated
	// field in passed movie object
	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m MovieModel) Get(movie *Movie) (*Movie, error) {
	return nil, nil
}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m MovieModel) Delete(movie *Movie) error {
	return nil
}
