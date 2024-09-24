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
	Gnovels GnovelModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Gnovels: GnovelModel{DB: db},
	}
}
