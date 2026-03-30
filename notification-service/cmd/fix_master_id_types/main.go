package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// This command is a one-time local dev helper.
// It enforces INTEGER auto-increment IDs (SERIAL/BIGSERIAL) for the master tables
// that are still currently UUID-based in the target PostgreSQL instance.
//
// NOTE: This drops and recreates the affected master tables.
func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	// Load SQL migration file from notification-service/migrations.
	sqlPath := filepath.Join("migrations", "001_init.sql")
	sqlBytes, err := os.ReadFile(sqlPath)
	if err != nil {
		log.Fatalf("read migration SQL (%s): %v", sqlPath, err)
	}
	migrationSQL := string(sqlBytes)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer pool.Close()

	log.Println("Applying master ID type fix (drop/recreate)...")

	// Tables that are still uuid-based in the DB (per schema_probe output).
	// Drop order matters only insofar as foreign keys exist; CASCADE handles it.
	toDrop := []string{
		"channel_provider_settings_master",
		"provider_routing_rules_master",
		"notification_templates_master",
		"template_groups_master",
		"channel_providers_master",
	}

	for _, t := range toDrop {
		stmt := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", t)
		if _, err := pool.Exec(ctx, stmt); err != nil {
			log.Fatalf("drop %s: %v", t, err)
		}
		log.Printf("Dropped %s", t)
	}

	// Execute the full init migration file; CREATE TABLE IF NOT EXISTS ensures
	// we only create what is missing (our drops above).
	//
	// We split by ';' for simplicity since this migration file doesn't contain
	// any procedural SQL blocks.
	stmts := strings.Split(migrationSQL, ";")
	executed := 0
	for _, raw := range stmts {
		stmt := strings.TrimSpace(raw)
		if stmt == "" {
			continue
		}

		if _, err := pool.Exec(ctx, stmt); err != nil {
			log.Fatalf("exec failed after splitting (%q...): %v", stmt[:min(60, len(stmt))], err)
		}
		executed++
	}

	log.Printf("Migration executed statements: %d", executed)
	log.Println("Done.")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

