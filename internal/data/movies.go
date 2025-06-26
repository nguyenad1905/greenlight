package data

import (
	"encoding/json"
	"fmt"
	"time"
)

type Movie struct {
    ID        int64     `json:"id"`
    CreatedAt time.Time `json:"-"` // Use the - directive
    Title     string    `json:"title"`
    Year      int32     `json:"year,omitempty"`    // Add the omitempty directive
    Runtime   int32     `json:"-"`
    Genres    []string  `json:"genres,omitempty"`  // Add the omitempty directive
    Version   int32     `json:"version"`
}

// Implement a MarshalJSON() method on the Movie struct, so that it satisfies the
// json.Marshaler interface.
func (m Movie) MarshalJSON() ([]byte, error) {
    // Declare a variable to hold the custom runtime string (this will be the empty 
    // string "" by default).
    var runtime string

    // If the value of the Runtime field is not zero, set the runtime variable to be a
    // string in the format "<runtime> mins".
    if m.Runtime != 0 {
        runtime = fmt.Sprintf("%d mins", m.Runtime)
    }

    type MovieAlias Movie

    // Create an anonymous struct to hold the data for JSON encoding. This has exactly
    // the same fields, types and tags as our Movie struct, except that the Runtime
    // field here is a string, instead of an int32. Also notice that we don't include
    // a CreatedAt field at all (there's no point including one, because we don't want
    // it to appear in the JSON output).
    aux := struct {
        MovieAlias
        Runtime string `json:"runtime,omitempty"`
    }{
        MovieAlias: MovieAlias(m),
        Runtime:    runtime,
    }

    // Encode the anonymous struct to JSON, and return it.
    return json.Marshal(aux)
}