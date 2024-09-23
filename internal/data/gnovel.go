package data

import (
	"database/sql"
	"time"

	"juanmagc99.comic-chest/internal/validator"
)

var gntypes = []string{"Manga", "Manhwa", "Manhua", "Comic"}
var genres = []string{"Action", "Adventure", "Comedy", "Drama", "Fantasy", "Historical",
	"Horror", "Isekai", "Magic", "Martial Arts", "Mecha", "Military", "Mystery", "Psychological",
	"Romance", "School Life", "Sci-Fi", "Seinen", "Shoujo", "Shounen", "Slice of Life", "Sports",
	"Supernatural", "Thriller", "Tragedy", "Vampire"}
var status = []string{"ongoing", "ended", ""}

type Gnovel struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	GNType      string    `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Genres      []string  `json:"genres"`
	Status      string    `json:"status"`
	NChapers    int       `json:"nchapters"`
	Author      string    `json:"author"`
	Year        int32     `json:"year"`
}

type MovieModel struct {
	DB *sql.DB
}

func ValidateMovie(v *validator.Validator, gnovel *Gnovel) {
	v.Check(gnovel.Title != "", "title", "must be provided")
	v.Check(len(gnovel.Title) <= 300, "title", "must not be more than 300 bytes long")

	v.Check(gnovel.Description != "", "description", "must be provided")
	v.Check(len(gnovel.Description) <= 5000, "description", "must not be more than 500 bytes long")

	v.Check(gnovel.Author != "", "author", "must be provided or be 'Anonymous'")

	v.Check(gnovel.Year != 0, "year", "must be provided")
	v.Check(gnovel.Year >= 1900, "year", "must be greater than 1900")
	v.Check(gnovel.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(gnovel.Genres != nil, "type", "must be provided")
	v.Check(validator.PermittedValue(gnovel.GNType, gntypes...), "type", "type must be one of the valid options")

	v.Check(gnovel.Genres != nil, "genres", "must be provided")
	v.Check(len(gnovel.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(gnovel.Genres) <= 10, "genres", "must not contain more than 10 genres")
	v.Check(validator.Unique(gnovel.Genres), "genres", "must not contain duplicate values")

	for _, genre := range genres {
		v.Check(validator.PermittedValue(genre, genres...), "genres", "genres must be valid options")
	}

	v.Check(validator.PermittedValue(gnovel.Status), "status", "status mus be a valid option")

	v.Check(gnovel.NChapers >= 0, "nchapters", "chapters must be equal or greater than 0")

}
