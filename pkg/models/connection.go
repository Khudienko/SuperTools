package models

import "github.com/jmoiron/sqlx"

type Connection struct {
	DB *sqlx.DB
}

