package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-f0532417/backend/migrations"
	"github.com/jackc/pgx/v5/stdlib"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("open database: %v", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := migrate(ctx, db); err != nil {
		log.Fatalf("migrate: %v", err)
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("database ping: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		if err := db.PingContext(r.Context()); err != nil {
			http.Error(w, `{"error":{"code":"SERVICE_UNAVAILABLE","message":"Service unavailable."}}`, http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func migrate(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (name text PRIMARY KEY, checksum text NOT NULL, applied_at timestamptz NOT NULL DEFAULT now())`); err != nil {
		return err
	}
	entries, err := fs.ReadDir(migrations.Files, ".")
	if err != nil {
		return err
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".up.sql") {
			continue
		}
		body, err := migrations.Files.ReadFile(name)
		if err != nil {
			return err
		}
		checksum := fmt.Sprintf("%x", sha256.Sum256(body))
		var existing string
		err = db.QueryRowContext(ctx, "SELECT checksum FROM schema_migrations WHERE name = $1", name).Scan(&existing)
		if err == nil {
			if existing != checksum {
				return fmt.Errorf("migration %s checksum changed", name)
			}
			continue
		}
		if err != sql.ErrNoRows {
			return err
		}
		if strings.Contains(strings.ToUpper(string(body)), "CREATE INDEX CONCURRENTLY") {
			if _, err = db.ExecContext(ctx, string(body)); err == nil {
				_, err = db.ExecContext(ctx, "INSERT INTO schema_migrations (name, checksum) VALUES ($1, $2)", name, checksum)
			}
		} else {
			tx, txErr := db.BeginTx(ctx, nil)
			if txErr != nil { return txErr }
			if _, err = tx.ExecContext(ctx, string(body)); err == nil {
				_, err = tx.ExecContext(ctx, "INSERT INTO schema_migrations (name, checksum) VALUES ($1, $2)", name, checksum)
			}
			if err != nil { _ = tx.Rollback(); return err }
			err = tx.Commit()
		}
		if err != nil { return fmt.Errorf("apply %s: %w", name, err) }
	}
	return nil
}
