package data

import (
	"database/sql"
	"time"

	"juanmagc99.comic-chest/internal/validator"
)

type Chapter struct {
	ID        int64     `json:"id"`
	GnovelID  int64     `json:"gnovel_id"`
	Number    int       `json:"number"`
	PageCount int       `json:"page_count"`
	CreatedAt time.Time `json:"created_at"`
}

type ChapterModel struct {
	DB *sql.DB
}

func ValidateChapter(v *validator.Validator, chp *Chapter) {
	v.Check(chp.Number >= 0, "number", "Number of chapter cant be negative")

	v.Check(chp.PageCount >= 1, "page_count", "A chapter must have at least one page")
}
