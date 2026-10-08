package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/jackc/pgx/v5/pgxpool"
// 	"github.com/joho/godotenv"
// )

// func checkDatabase(ctx context.Context, pool *pgxpool.Pool) error {
// 	var result int

// 	err := pool.QueryRow(ctx, "select 1").Scan(&result)
// 	if err != nil {
// 		return err
// 	}

// 	if result != 1 {
// 		return fmt.Errorf("Unexpected db result: %d", result)
// 	}

// 	return nil
// }

// func getContent(ctx context.Context, pool *pgxpool.Pool) any {
// 	var result any

// 	err := pool.QueryRow(ctx, "select name from customer").Scan(&result)
// 	if err != nil {
// 		return err
// 	}

// 	return result
// }

// func main() {
// 	fmt.Println("The Weeknd is the GOAT bro")

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	dbString := os.Getenv("PSQL_STRING")
// 	ctx, cancel := context.WithTimeout(
// 		context.Background(),
// 		5*time.Second,
// 	)
// 	defer cancel()

// 	// parse the config first
// 	config, err := pgxpool.ParseConfig(dbString)
// 	if err != nil {
// 		log.Fatal("There is some error with the DB")
// 		return
// 	}

// 	config.MaxConns = 10
// 	config.MaxConnLifetime = 5 * time.Minute

// 	pool, err := pgxpool.NewWithConfig(ctx, config)
// 	if err != nil {
// 		log.Fatal("There is some error with the DB")
// 		return
// 	}
// 	defer pool.Close()

// 	err = pool.Ping(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	if err := checkDatabase(ctx, pool); err != nil {
// 		log.Fatal(err)
// 	}

// 	content := getContent(ctx, pool)
// 	fmt.Println(content)

// 	fmt.Println("Connected to PostgreSQL")
// }
