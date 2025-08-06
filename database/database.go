package database

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
)

const databaseName = "fox.db"
const appName = "CodeFox"

// GetConnection gets connection of sqlite
func GetConnection() (*sql.DB, error) {
	dbPath := getDirectory()
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDirectory() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	appDir := filepath.Join(dir, appName)
	err = os.MkdirAll(appDir, 0700)
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(appDir, databaseName)
	return dbPath
}
