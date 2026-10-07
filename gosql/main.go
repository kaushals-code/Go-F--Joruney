package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("The Weeknd is the GOAT bro")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	dbString := os.Getenv("PSQL_STRING")
	fmt.Println(dbString != "")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbString)
	if err != nil {
		log.Fatal("There is some error with the DB")
		return
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL")
}
