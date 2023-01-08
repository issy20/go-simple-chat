package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Conn struct {
	DB *sqlx.DB
}

func NewDB() (*Conn, error) {
	db, err := sqlx.Connect("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		return nil, fmt.Errorf("failed to open MYSQL: %w", err)
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to db ping : %w", err)
	}

	return &Conn{DB: db}, nil
}
