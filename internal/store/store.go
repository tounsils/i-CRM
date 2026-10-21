// Package store owns persistence. It holds the SQLite handle, applies the
// embedded migrations and exposes one method per thing the web layer needs to
// read or write. SQL does not leave this package.
package store

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"sort"
	"time"

	_ "modernc.org/sqlite" // pure-Go driver, so the binary stays CGO-free
)

//go:embed migrations/*.sql
var migrationFS embed.FS

// ErrNotFound is returned instead of sql.ErrNoRows so that callers outside this
// package never need to import database/sql.
var ErrNotFound = errors.New("store: not found")

// maxOpenConns is 1 on purpose, which serialises every query including reads.
//
// SQLite permits one writer at a time, so a pool larger than 1 trades
// SQLITE_BUSY retries and pragma-per-connection bookkeeping for throughput this
// product does not need: a self-hosted CRM serves the staff of one business,
// where request concurrency is single digits. This is a correctness-for-
// throughput trade made deliberately and not a measured figure. If a deployment
// ever outgrows it, raise it after measuring, and move journal_mode and
// foreign_keys handling to a connection hook at the same time.
const maxOpenConns = 1

// Store is the database handle plus the queries over it.
type Store struct {
	db *sql.DB
}

// Open connects to the SQLite file at path, applies pragmas and runs any
// migrations that have not been applied yet.
//
// foreign_keys is set explicitly because SQLite defaults it to OFF, which would
// silently accept a line pointing at a deleted document.
func Open(path string) (*Store, error) {
	dsn := fmt.Sprintf(
		"file:%s?_pragma=foreign_keys(1)&_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)",
		path,
	)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("store: opening %s: %w", path, err)
	}
	db.SetMaxOpenConns(maxOpenConns)

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("store: connecting to %s: %w", path, err)
	}

	s := &Store{db: db}
	if err := s.migrate(context.Background()); err != nil {
		_ = db.Close()
		return nil, err
	}
	return s, nil
}

// Close releases the handle.
func (s *Store) Close() error {
	return s.db.Close()
}

// migrate applies every embedded migration that is not already recorded, in
// filename order, each in its own transaction.
func (s *Store) migrate(ctx context.Context) error {
	const createLedger = `CREATE TABLE IF NOT EXISTS schema_migrations (
		name       TEXT PRIMARY KEY,
		applied_at TEXT NOT NULL
	)`
	if _, err := s.db.ExecContext(ctx, createLedger); err != nil {
		return fmt.Errorf("store: creating migration ledger: %w", err)
	}

	names, err := fs.Glob(migrationFS, "migrations/*.sql")
	if err != nil {
		return fmt.Errorf("store: listing migrations: %w", err)
	}
	sort.Strings(names)

	for _, name := range names {
		applied, err := s.migrationApplied(ctx, name)
		if err != nil {
			return err
		}
		if applied {
			continue
		}
		if err := s.applyMigration(ctx, name); err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) migrationApplied(ctx context.Context, name string) (bool, error) {
	var count int
	row := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM schema_migrations WHERE name = ?`, name)
	if err := row.Scan(&count); err != nil {
		return false, fmt.Errorf("store: checking migration %s: %w", name, err)
	}
	return count > 0, nil
}

func (s *Store) applyMigration(ctx context.Context, name string) error {
	body, err := migrationFS.ReadFile(name)
	if err != nil {
		return fmt.Errorf("store: reading migration %s: %w", name, err)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("store: beginning migration %s: %w", name, err)
	}
	defer tx.Rollback() //nolint:errcheck // no-op once Commit has succeeded

	if _, err := tx.ExecContext(ctx, string(body)); err != nil {
		return fmt.Errorf("store: applying migration %s: %w", name, err)
	}
	const record = `INSERT INTO schema_migrations (name, applied_at) VALUES (?, ?)`
	if _, err := tx.ExecContext(ctx, record, name, formatTime(time.Now())); err != nil {
		return fmt.Errorf("store: recording migration %s: %w", name, err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("store: committing migration %s: %w", name, err)
	}
	return nil
}

// formatTime renders a timestamp the way every TEXT time column stores it.
func formatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

// parseTime reads a timestamp back out of a TEXT time column.
func parseTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// parseNullTime reads a nullable TEXT time column.
func parseNullTime(ns sql.NullString) (*time.Time, error) {
	if !ns.Valid || ns.String == "" {
		return nil, nil
	}
	t, err := parseTime(ns.String)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// nullTime renders an optional timestamp for storage.
func nullTime(t *time.Time) any {
	if t == nil {
		return nil
	}
	return formatTime(*t)
}

// nullID renders an optional foreign key for storage.
func nullID(id *int64) any {
	if id == nil {
		return nil
	}
	return *id
}

// idFromNull reads a nullable foreign key.
func idFromNull(n sql.NullInt64) *int64 {
	if !n.Valid {
		return nil
	}
	v := n.Int64
	return &v
}
