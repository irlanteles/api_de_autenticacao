package db

import (
    "database/sql"
    "fmt"
    _"github.com/lib/pq" // Driver do PostgreSQL
    "os"
    "github.com/joho/godotenv"
    "log"
)

var DB *sql.DB

func Connect() error {

    if err := godotenv.Load(); err != nil {
        log.Println("Não foi possível carregar .env, usando variáveis do sistema")
    }
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

	var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        return err
    }

    return DB.Ping()
}