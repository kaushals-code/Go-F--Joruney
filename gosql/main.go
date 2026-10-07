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

func checkDatabase(ctx context.Context, pool *pgxpool.Pool) error {
	var result int

	err := pool.QueryRow(ctx, "select 1").Scan(&result)
	if err != nil {
		return err
	}

	if result != 1 {
		return fmt.Errorf("Unexpected db result: %d", result)
	}

	return nil
}

func getContent(ctx context.Context, pool *pgxpool.Pool) error {
	var result int

	err := pool.QueryRow(ctx, "select * from customer").Scan(&result)
	if err != nil {
		return err
	}

	if result != 1 {
		return fmt.Errorf("Unexpected db result: %d", result)
	}

	return nil
}

func main() {
	fmt.Println("The Weeknd is the GOAT bro")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	dbString := os.Getenv("PSQL_STRING")
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	// parse the config first
	config, err := pgxpool.ParseConfig(dbString)
	if err != nil {
		log.Fatal("There is some error with the DB")
		return
	}

	config.MaxConns = 10
	config.MaxConnLifetime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal("There is some error with the DB")
		return
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := checkDatabase(ctx, pool); err != nil {
		log.Fatal(err)
	}

	if err := 

	fmt.Println("Connected to PostgreSQL")
}
