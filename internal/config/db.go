package config

import "database/sql"

type Database struct {
	db *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}
