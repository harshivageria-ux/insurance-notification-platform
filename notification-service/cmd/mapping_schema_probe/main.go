package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer pool.Close()

	// Mapping tables used by gin-mapping-api repository.
	tables := []string{
		"notification_category_channel_master",
		"channel_provider_master_map",
		"template_channel_language_master_map",
	}

	for _, table := range tables {
		fmt.Printf("\nTABLE %s\n", table)

		// Print whether table exists
		var exists int
		err := pool.QueryRow(ctx, `
			SELECT COUNT(*)::int
			FROM information_schema.tables
			WHERE table_schema = 'public' AND table_name = $1
		`, table).Scan(&exists)
		if err != nil {
			log.Fatalf("table exists check: %v", err)
		}
		if exists == 0 {
			fmt.Println("  <missing>")
			continue
		}

		rows, err := pool.Query(ctx, `
			SELECT column_name, data_type
			FROM information_schema.columns
			WHERE table_schema = 'public' AND table_name = $1
			ORDER BY ordinal_position
		`, table)
		if err != nil {
			log.Fatalf("columns query for %s: %v", table, err)
		}

		found := 0
		for rows.Next() {
			var name string
			var typ string
			if err := rows.Scan(&name, &typ); err != nil {
				log.Fatalf("scan columns for %s: %v", table, err)
			}
			found++
			fmt.Printf("  %s :: %s\n", name, typ)
		}
		rows.Close()

		if found == 0 {
			fmt.Println("  <no columns>")
		}
	}
}

