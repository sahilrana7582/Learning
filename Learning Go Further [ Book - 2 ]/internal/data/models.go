package data

import "database/sql"

type Models struct {
	Movies MovieModal
	User   UserModel
}

func NewModal(db *sql.DB) *Models {
	return &Models{
		Movies: MovieModal{DB: db},
		User:   UserModel{DB: db},
	}
}
