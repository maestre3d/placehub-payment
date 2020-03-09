package rdbms

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresPool(username, password, host, port, database string) *sql.DB {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(8 * time.Second)
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(1)

	return db
}
