package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Gnovels  GnovelModel
	Chapters ChapterModel
	Users    UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Gnovels:  GnovelModel{DB: db},
		Chapters: ChapterModel{DB: db},
		Users:    UserModel{DB: db},
	}
}
