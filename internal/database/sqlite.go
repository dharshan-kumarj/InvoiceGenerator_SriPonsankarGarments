// internal/database/sqlite.go
package database

import (
    "database/sql"
)

func InitDB(db *sql.DB) error {
    // Create users table if it doesn't exist
    _, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`)
    return err
}