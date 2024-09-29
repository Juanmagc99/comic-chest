package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"juanmagc99.comic-chest/internal/validator"
)

var gntypes = []string{"Manga", "Manhwa", "Manhua", "Comic"}
var genres = []string{"Action", "Adventure", "Comedy", "Drama", "Fantasy", "Historical",
	"Horror", "Isekai", "Magic", "Martial Arts", "Mecha", "Military", "Mystery", "Psychological",
	"Romance", "School Life", "Sci-Fi", "Seinen", "Shoujo", "Shounen", "Slice of Life", "Sports",
	"Supernatural", "Thriller", "Tragedy", "Vampire"}
var status = []string{"ongoing", "ended", ""}

type Gnovel struct {
	ID          int64     `json:"-"`
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

type GnovelModel struct {
	DB *sql.DB
}

func ValidateGnovel(v *validator.Validator, gnovel *Gnovel) {
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

	v.Check(validator.PermittedValue(gnovel.Status, status...), "status", "status mus be a valid option")

	v.Check(gnovel.NChapers >= 0, "nchapters", "chapters must be equal or greater than 0")

}

func (m GnovelModel) Insert(gnovel *Gnovel) error {
	query := `
		INSERT INTO gnovels (gn_type, title, description, genres, nchapters, author, year, status)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		RETURNING id, created_at
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	args := []any{gnovel.GNType, gnovel.Title, gnovel.Description, pq.Array(gnovel.Genres),
		0, gnovel.Author, gnovel.Year, gnovel.Status}

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&gnovel.ID, &gnovel.CreatedAt)
}

func (m GnovelModel) Get(id int64) (*Gnovel, error) {
	//Id already checked before this function

	query := `
		SELECT *
		FROM gnovels
		WHERE id = $1
	`

	var gnovel Gnovel

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&gnovel.ID,
		&gnovel.CreatedAt,
		&gnovel.GNType,
		&gnovel.Title,
		&gnovel.Description,
		pq.Array(&gnovel.Genres),
		&gnovel.Status,
		&gnovel.NChapers,
		&gnovel.Author,
		&gnovel.Year,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	// Otherwise, return a pointer to the Movie struct.
	return &gnovel, nil

}

func (m GnovelModel) Delete(id int64) error {

	query := `
		DELETE FROM gnovels
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m GnovelModel) Update(gnovel *Gnovel) error {
	query := `
		UPDATE gnovels
		SET gn_type = $1, title = $2, description = $3, genres = $4, 
		author = $5, year = $6, status = $7, 
		WHERE id = $8
		RETURNING ID`

	args := []any{
		gnovel.GNType,
		gnovel.Title,
		gnovel.Description,
		pq.Array(gnovel.Genres),
		gnovel.Author,
		gnovel.Year,
		gnovel.Status,
		gnovel.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&gnovel.ID)
	if err != nil {
		return err
	}

	return nil
}
