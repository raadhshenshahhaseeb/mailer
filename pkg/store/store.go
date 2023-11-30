package store

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

func New() (Store, error) {
	sqlDB, err := initSQLite()
	if err != nil {
		return nil, err
	}

	return &store{
		sqlDB:    sqlDB,
		sqlMutex: &sync.Mutex{},
	}, nil
}

func initSQLite() (*sql.DB, error) {
	// Create the directory if it does not exist
	dbPath := "data/db"
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err := os.MkdirAll(dbPath, 0755) // Create the directory with appropriate permissions
		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite3", "./data/db/av.db")
	if err != nil {
		return nil, err
	}

	// Use a ping to ensure the database is accessible
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = migrations(db)
	if err != nil {
		return nil, err
	}

	return db, err
}

func migrations(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS emails (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

type store struct {
	sqlDB    *sql.DB
	sqlMutex *sync.Mutex
}

type Store interface {
	InsertRecord(email string) error
	Get() (*[]Emails, error)
}
