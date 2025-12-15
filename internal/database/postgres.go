package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectPostgres() *pgxpool.Pool {
	_ = godotenv.Load()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("❌ DATABASE_URL not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal("❌ pgx parse config error:", err)
	}

	// optional tuning
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = time.Minute * 5

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal("❌ pgx pool error:", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("❌ pgx ping error:", err)
	}

	log.Println("✅ PostgreSQL connected using pgx")
	return pool
}
