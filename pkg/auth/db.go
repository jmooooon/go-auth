package auth

import (
	"database/sql"
	// ...
	_ "github.com/lib/pq"
)

var (
	// DB ...
	DB *sql.DB
)

// NewDB ...
func NewDB(dbURL string) (error) {
	DB, err := sql.Open("postgres", dbURL)

	if err != nil {
		return err
	}

	if err = DB.Ping(); err!=nil {
		return err
	}

	return nil
}
