package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/tounsils/i-CRM/internal/domain"
)

const customerColumns = `id, name, contact_name, email, phone, address, status, currency, notes, created_at`

// CreateCustomer inserts a lead or customer and returns the new id.
func (s *Store) CreateCustomer(ctx context.Context, c Customer) (int64, error) {
	const q = `INSERT INTO customers
	           (name, contact_name, email, phone, address, status, currency, notes, created_at)
	           VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := s.db.ExecContext(ctx, q,
		c.Name, c.ContactName, c.Email, c.Phone, c.Address,
		string(c.Status), c.Currency, c.Notes, formatTime(time.Now()))
	if err != nil {
		return 0, fmt.Errorf("store: creating customer: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("store: creating customer: %w", err)
	}
	return id, nil
}

// UpdateCustomer overwrites the editable fields of an existing record.
func (s *Store) UpdateCustomer(ctx context.Context, c Customer) error {
	const q = `UPDATE customers
	           SET name = ?, contact_name = ?, email = ?, phone = ?, address = ?,
	               status = ?, currency = ?, notes = ?
	           WHERE id = ?`
	res, err := s.db.ExecContext(ctx, q,
		c.Name, c.ContactName, c.Email, c.Phone, c.Address,
		string(c.Status), c.Currency, c.Notes, c.ID)
	if err != nil {
		return fmt.Errorf("store: updating customer %d: %w", c.ID, err)
	}
	return requireOneRow(res, fmt.Sprintf("customer %d", c.ID))
}

// CustomerByID loads one record.
func (s *Store) CustomerByID(ctx context.Context, id int64) (Customer, error) {
	q := `SELECT ` + customerColumns + ` FROM customers WHERE id = ?`
	return scanCustomer(s.db.QueryRowContext(ctx, q, id))
}

// ListCustomers returns records ordered by name. An empty status returns every
// record, so the same method serves the all-customers screen and each pipeline
// filter without a second query.
func (s *Store) ListCustomers(ctx context.Context, status domain.CustomerStatus) ([]Customer, error) {
	q := `SELECT ` + customerColumns + ` FROM customers`
	args := []any{}
	if status != "" {
		q += ` WHERE status = ?`
		args = append(args, string(status))
	}
	q += ` ORDER BY name COLLATE NOCASE`

	rows, err := s.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("store: listing customers: %w", err)
	}
	defer rows.Close()

	var out []Customer
	for rows.Next() {
		c, err := scanCustomerRows(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("store: listing customers: %w", err)
	}
	return out, nil
}

// CountCustomersByStatus powers the dashboard tiles in one query rather than
// one per status.
func (s *Store) CountCustomersByStatus(ctx context.Context) (map[domain.CustomerStatus]int, error) {
	const q = `SELECT status, COUNT(*) FROM customers GROUP BY status`
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("store: counting customers: %w", err)
	}
	defer rows.Close()

	out := map[domain.CustomerStatus]int{}
	for rows.Next() {
		var (
			status string
			count  int
		)
		if err := rows.Scan(&status, &count); err != nil {
			return nil, fmt.Errorf("store: counting customers: %w", err)
		}
		out[domain.CustomerStatus(status)] = count
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("store: counting customers: %w", err)
	}
	return out, nil
}

// rowScanner is satisfied by both *sql.Row and *sql.Rows, so one scan function
// serves a single lookup and a list.
type rowScanner interface {
	Scan(dest ...any) error
}

func scanCustomer(row *sql.Row) (Customer, error) {
	c, err := scanCustomerRows(row)
	if errors.Is(err, sql.ErrNoRows) {
		return Customer{}, ErrNotFound
	}
	return c, err
}

func scanCustomerRows(row rowScanner) (Customer, error) {
	var (
		c         Customer
		status    string
		createdAt string
	)
	err := row.Scan(&c.ID, &c.Name, &c.ContactName, &c.Email, &c.Phone,
		&c.Address, &status, &c.Currency, &c.Notes, &createdAt)
	if errors.Is(err, sql.ErrNoRows) {
		return Customer{}, sql.ErrNoRows
	}
	if err != nil {
		return Customer{}, fmt.Errorf("store: reading customer: %w", err)
	}
	c.Status = domain.CustomerStatus(status)
	if c.CreatedAt, err = parseTime(createdAt); err != nil {
		return Customer{}, fmt.Errorf("store: reading customer timestamp: %w", err)
	}
	return c, nil
}

// requireOneRow turns a write that matched nothing into ErrNotFound, so an
// update against a deleted id reports a missing record instead of succeeding
// silently.
func requireOneRow(res sql.Result, what string) error {
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("store: checking write to %s: %w", what, err)
	}
	if affected == 0 {
		return fmt.Errorf("%w: %s", ErrNotFound, what)
	}
	return nil
}
