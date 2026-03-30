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

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer pool.Close()

	code := os.Getenv("PRIORITY_CODE")
	if code == "" {
		code = fmt.Sprintf("DBG_%d", time.Now().Unix())
	}
	createdBy := os.Getenv("PRIORITY_CREATED_BY")
	if createdBy == "" {
		createdBy = "admin"
	}
	desc := os.Getenv("PRIORITY_DESC")
	if desc == "" {
		desc = "debug insert"
	}

	// Mirror notification-service/internal/infrastructure/repository/priority_repository.go
	q := `
		INSERT INTO notification_priority_master (priority_code, description, created_at, created_by, is_active, version)
		VALUES ($1, $2, NOW(), $3, true, 1)
		RETURNING priority_id, COALESCE(priority_code, ''), COALESCE(description, ''), created_at,
		          COALESCE(created_by, ''), COALESCE(is_active, false), COALESCE(version, 0)
	`

	fmt.Println("Inserting priority_code:", code)

	var (
		priorityID  int16
		outCode     string
		outDesc     string
		createdAt   time.Time
		outCreated  string
		isActive    bool
		outVersion  int
	)

	err = pool.
		QueryRow(ctx, q, code, desc, createdBy).
		Scan(&priorityID, &outCode, &outDesc, &createdAt, &outCreated, &isActive, &outVersion)
	if err != nil {
		// Print the raw pg error message for debugging.
		fmt.Println("Insert failed:", err)
		os.Exit(1)
	}

	fmt.Printf("Insert OK: priority_id=%d code=%s version=%d is_active=%v created_by=%s\n",
		priorityID, outCode, outVersion, isActive, outCreated)
}

