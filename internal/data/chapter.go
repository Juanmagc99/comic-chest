package data

import (
	"context"
	"database/sql"
	"time"

	"juanmagc99.comic-chest/internal/validator"
)

type Chapter struct {
	ID        int64     `json:"-"`
	GnovelID  int64     `json:"gnovel_id"`
	Number    int       `json:"number"`
	FilePath  string    `json:"-"`
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
