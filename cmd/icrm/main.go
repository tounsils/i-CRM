// Command icrm is the i-CRM server: one binary that carries its own
// migrations, templates and stylesheet, and stores everything in one SQLite
// file. Configuration comes from the environment; see internal/config.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/tounsils/i-CRM/internal/config"
	"github.com/tounsils/i-CRM/internal/store"
	"github.com/tounsils/i-CRM/internal/web"
)

// Server timeouts. A self-hosted install is often reached over a slow link, so
// these are generous, but none of them is unbounded: an unbounded read timeout
// means one stalled client holds a connection for as long as it likes.
const (
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 30 * time.Second
	writeTimeout      = 60 * time.Second
	idleTimeout       = 2 * time.Minute

	// shutdownGrace is how long in-flight requests get to finish on SIGTERM.
	// Ten seconds comfortably covers the slowest request here, a bcrypt
	// verification plus a handful of SQLite queries.
	shutdownGrace = 10 * time.Second

	// databaseDirMode keeps the database directory owner-only. The file holds
	// every customer record and every invoice, so it is not world-readable.
	databaseDirMode = 0o750
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("i-CRM: %v", err)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	if dir := filepath.Dir(cfg.DatabasePath); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, databaseDirMode); err != nil {
			return fmt.Errorf("creating the database directory %s: %w", dir, err)
		}
	}

	db, err := store.Open(cfg.DatabasePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("i-CRM: closing the database: %v", err)
		}
	}()

	server, err := web.New(cfg, db)
	if err != nil {
		return err
	}

	if cfg.SessionKeyGenerated {
		log.Print("i-CRM: ICRM_SESSION_KEY is not set, so a one-off key was generated. " +
			"Sessions will not survive a restart. Set it before using this in earnest.")
	}

	httpServer := &http.Server{
		Addr:              cfg.Addr,
		Handler:           server.Routes(),
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	listenErr := make(chan error, 1)
	go func() {
		log.Printf("i-CRM: listening on %s, database at %s", cfg.Addr, cfg.DatabasePath)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			listenErr <- err
		}
	}()

	select {
	case err := <-listenErr:
		return fmt.Errorf("serving: %w", err)
	case <-ctx.Done():
	}

	log.Print("i-CRM: shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownGrace)
	defer cancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutting down: %w", err)
	}
	return nil
}
