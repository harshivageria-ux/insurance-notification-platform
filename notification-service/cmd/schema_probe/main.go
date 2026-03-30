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

	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer db.Close()

	tables := []string{
		"languages",
		"priorities",
		"statuses",
		"schedule_types",
		"categories",
		"channels",
		"channel_providers",
		"provider_settings",
		"template_groups",
		"templates",
		"routing_rules",
		"languages_master",
		"notification_priority_master",
		"notification_statuses_master",
		"notification_schedule_type_master",
		"notification_categories_master",
		"notification_channels_master",
		"channel_providers_master",
		"channel_provider_settings_master",
		"template_groups_master",
		"notification_templates_master",
		"provider_routing_rules_master",
	}

	for _, table := range tables {
		fmt.Printf("TABLE %s\n", table)
		rows, err := db.Query(ctx, `
			SELECT column_name, data_type
			FROM information_schema.columns
			WHERE table_schema = 'public' AND table_name = $1
			ORDER BY ordinal_position
		`, table)
		if err != nil {
			log.Fatalf("query columns for %s: %v", table, err)
		}

		count := 0
		for rows.Next() {
			var name string
			var dataType string
			if err := rows.Scan(&name, &dataType); err != nil {
				rows.Close()
				log.Fatalf("scan columns for %s: %v", table, err)
			}
			count++
			fmt.Printf("  %s :: %s\n", name, dataType)
		}
		rows.Close()

		if count == 0 {
			fmt.Println("  <missing>")
		}
	}

	fmt.Println("PUBLIC TABLES")
	rows, err := db.Query(ctx, `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
		ORDER BY table_name
	`)
	if err != nil {
		log.Fatalf("query public tables: %v", err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("scan public tables: %v", err)
		}
		count++
		fmt.Printf("  %s\n", tableName)
	}

	if count == 0 {
		fmt.Println("  <none>")
	}

	fmt.Println("PROBE QUERIES")
	runProbe(ctx, db, "languages", `SELECT id::text, name, code, created_by, COALESCE(updated_by, ''), is_active::text, version::text, created_at::text, updated_at::text FROM languages_master LIMIT 1`)
	runProbe(ctx, db, "priorities", `SELECT priority_id::text, priority_code, description, created_at::text, created_by, is_active::text, version::text, ''::text, ''::text FROM notification_priority_master LIMIT 1`)
}

func runProbe(ctx context.Context, db *pgxpool.Pool, label, query string) {
	row := db.QueryRow(ctx, query)
	var a, b, c, d, e, f, g, h, i string
	if err := row.Scan(&a, &b, &c, &d, &e, &f, &g, &h, &i); err != nil {
		fmt.Printf("  %s ERROR: %v\n", label, err)
		return
	}

	fmt.Printf("  %s OK\n", label)
}
