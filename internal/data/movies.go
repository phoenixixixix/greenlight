package data

import "time"

type Movie struct {
	// 1. Fields should start with capital letter to be exported.
	//    This way they are visible to encoding/json packege
	// 2. `json:"field"` is a Struct Tag which defines how fields
	//    should be parsed (to json in this case)
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // - (hyphen) copletly hide this field (ex. I don't want to show this field to end user)
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"` // omitempty hides field if it has its zero value
	Runtime   int32     `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
