package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Database holds the database connection
type Database struct {
	*sql.DB
}

// NewDatabase creates a new database connection
func NewDatabase() (*Database, error) {
	// For demo purposes, using a simple connection string
	// In production, these should be environment variables
	dsn := "root:Pw@#$234@tcp(localhost:3306)/restuarant?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	log.Println("Successfully connected to database")
	return &Database{db}, nil
}

// Close closes the database connection
func (db *Database) Close() error {
	return db.DB.Close()
}
