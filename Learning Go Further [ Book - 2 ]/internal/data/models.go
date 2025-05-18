package data

import "database/sql"

type Models struct {
	Movies MovieModal
}

func NewModal(db *sql.DB) *Models {
	return &Models{
		Movies: MovieModal{DB: db},
	}
}
