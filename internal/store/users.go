package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

// CreateUser inserts a staff member and returns the new id. The caller hashes
// the password; this package never sees a plaintext one.
func (s *Store) CreateUser(ctx context.Context, u User) (int64, error) {
	const q = `INSERT INTO users (email, name, password_hash, role, created_at)
	           VALUES (?, ?, ?, ?, ?)`
	email := strings.ToLower(strings.TrimSpace(u.Email))
	res, err := s.db.ExecContext(ctx, q, email, u.Name, u.PasswordHash, u.Role, formatTime(time.Now()))
	if err != nil {
		return 0, fmt.Errorf("store: creating user: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("store: creating user: %w", err)
	}
	return id, nil
}

// UserByEmail looks a user up for sign-in. The lookup is case-insensitive
// because an email that differs only in case is the same mailbox.
func (s *Store) UserByEmail(ctx context.Context, email string) (User, error) {
	const q = `SELECT id, email, name, password_hash, role, created_at
	           FROM users WHERE email = ?`
	row := s.db.QueryRowContext(ctx, q, strings.ToLower(strings.TrimSpace(email)))
	return scanUser(row)
}

// UserByID loads the signed-in user for a request.
func (s *Store) UserByID(ctx context.Context, id int64) (User, error) {
	const q = `SELECT id, email, name, password_hash, role, created_at
	           FROM users WHERE id = ?`
	return scanUser(s.db.QueryRowContext(ctx, q, id))
}

// CountUsers reports how many staff exist, which is how first-run setup decides
// whether to offer to create the first administrator.
func (s *Store) CountUsers(ctx context.Context) (int, error) {
	var n int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM users`).Scan(&n); err != nil {
		return 0, fmt.Errorf("store: counting users: %w", err)
	}
	return n, nil
}

func scanUser(row *sql.Row) (User, error) {
	var (
		u         User
		createdAt string
	)
	err := row.Scan(&u.ID, &u.Email, &u.Name, &u.PasswordHash, &u.Role, &createdAt)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, ErrNotFound
	}
	if err != nil {
		return User{}, fmt.Errorf("store: reading user: %w", err)
	}
	if u.CreatedAt, err = parseTime(createdAt); err != nil {
		return User{}, fmt.Errorf("store: reading user timestamp: %w", err)
	}
	return u, nil
}
