package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"juanmagc99.comic-chest/internal/validator"
)

type Chapter struct {
	ID        int64     `json:"-"`
	GnovelID  int64     `json:"gnovel_id"`
	Number    int       `json:"number"`
	FilePath  string    `json:"filepath"`
	CreatedAt time.Time `json:"-"`
}

type ChapterModel struct {
	DB *sql.DB
}

func ValidateChapter(v *validator.Validator, chp *Chapter, gm GnovelModel) {

	_, err := gm.Get(chp.GnovelID)

	v.Check(err == nil, "gnovel_id", "This graphic novel doesnt exists")

	v.Check(chp.Number >= 0, "number", "Number of chapter cant be negative")

	//v.Check(chp.PageCount >= 1, "page_count", "A chapter must have at least one page")
}

func (m ChapterModel) Insert(chapter *Chapter) error {
	query := `
		INSERT INTO chapters (gnovelid, number, filepath)
		VALUES ($1,$2,$3)
		RETURNING id, created_at
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	args := []any{chapter.GnovelID, chapter.Number, chapter.FilePath}

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&chapter.ID, &chapter.CreatedAt)
}

func (m ChapterModel) Get(id int64, number int) (*Chapter, error) {
	//Id already checked before this function

	query := `
		SELECT *
		FROM chapters
		WHERE gnovelid = $1 AND number = $2
	`

	var chapter Chapter

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	args := []any{id, number}

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&chapter.ID,
		&chapter.GnovelID,
		&chapter.Number,
		&chapter.FilePath,
		&chapter.CreatedAt,
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
	return &chapter, nil
}

func (m ChapterModel) Delete(id int64, number int) error {

	query := `
		DELETE FROM chapters
		WHERE gnovelid = $1 AND number = $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	args := []any{id, number}

	result, err := m.DB.ExecContext(ctx, query, args...)
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
