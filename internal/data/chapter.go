package data

import (
	"context"
	"database/sql"
	"time"

	"juanmagc99.comic-chest/internal/validator"
)

type Chapter struct {
	ID        int64     `json:"id"`
	GnovelID  int64     `json:"gnovel_id"`
	Number    int       `json:"number"`
	PageCount int       `json:"page_count"`
	FilePath  string    `json:"file_path"`
	CreatedAt time.Time `json:"created_at"`
}

type ChapterModel struct {
	DB *sql.DB
}

func ValidateChapter(v *validator.Validator, chp *Chapter) {
	v.Check(chp.Number >= 0, "number", "Number of chapter cant be negative")

	v.Check(chp.PageCount >= 1, "page_count", "A chapter must have at least one page")
}

func (m ChapterModel) Insert(chapter *Chapter) error {
	query := `
		INSERT INTO chapters (gnovelid, number, pagecount, filpath)
		VALUES ($1,$2,$3,$4)
		RETURNING id, created_at
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	args := []any{chapter.GnovelID, chapter.Number, chapter.PageCount, chapter.FilePath}

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&chapter.ID, &chapter.CreatedAt)
}
