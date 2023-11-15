package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	ExitSuccess = 0
	ExitFailure = 1
)

const (
	OSServerPort = 8080
	OSServerHost = "localhost"
	DBHost       = "localhost"
	DBPort       = 3306
	DBUser       = "user"
	DBPassword   = "password"
	DBName       = "database"
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	return listen(ctx, OSServerPort)
}

func listen(ctx context.Context, port int) int {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database?charset=utf8&parseTime=true")
	if err != nil {
		log.Printf("Failed to connect to database: %s", err)
		return ExitFailure
	}
	log.Println("Connected to database")
	defer db.Close()

	if err := db.PingContext(context.Background()); err != nil {
		log.Printf("Failed to ping database: %s", err)
		return ExitFailure
	}
	return 0
}
