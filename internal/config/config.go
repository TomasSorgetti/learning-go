package config

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
    "log"
    "os"
)

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using default environment variables")
    }
}

func GetDB() *sql.DB {
    connStr := os.Getenv("DATABASE_URL")
    if connStr == "" {
        log.Fatal("DATABASE_URL not set")
    }

    db, err := sql.Open("mysql", connStr)
    if err != nil {
        log.Fatal("Failed to connect to MySQL:", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal("Failed to ping MySQL:", err)
    }

    return db
}