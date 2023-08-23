package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TntraParas/jet-demo/repo"
	_ "github.com/lib/pq"
)

// postgresql://user:password@localhost:5433/jet?sslmode=disable

func main() {
	pool, err := ConnectToPG()
	if err != nil {
		log.Fatalf("main could not connect to DB: %s", err)
	}
	repo.CreateUser(pool)
}

func getDBURL() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s connect_timeout=%d",
		"localhost", 5433, "user", "password", "jet", "disable", 5,
	)
}

func ConnectToPG() (*sql.DB, error) {
	dbURL := getDBURL()

	pool, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	err = pool.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to db : %w", err)
	}
	log.Printf("database connected successfully")
	return pool, nil
}
