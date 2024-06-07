package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func MustInit() {
	url, set := os.LookupEnv("POSTGRES_URL")
	if !set {
		panic("Environment variable 'POSTGRES_URL' not set!")
	}
	var err error
	pool, err = pgxpool.New(context.Background(), url)
	if err != nil {
		panic(err)
	}
	log.Printf("Intialized db %q", url)
}

func Close() {
	pool.Close()
	log.Println("Closed db")
}
