package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/tounsils/i-CRM/internal/domain"
)

// CreateTax adds a named rate.
func (s *Store) CreateTax(ctx context.Context, t Tax) (int64, error) {
	const q = `INSERT INTO taxes (name, rate_bp, archived) VALUES (?, ?, ?)`
	res, err := s.db.ExecContext(ctx, q, t.Name, int64(t.Rate), t.Archived)
	if err != nil {
		return 0, fmt.Errorf("store: creating tax: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("store: creating tax: %w", err)
	}
	return id, nil
}

// ListTaxes returns the rates available to put on a line. Archived rates are
// excluded from the picker but never deleted, because lines already issued
// reference them by name and rate.
func (s *Store) ListTaxes(ctx context.Context) ([]Tax, error) {
	const q = `SELECT id, name, rate_bp, archived FROM taxes
	           WHERE archived = 0 ORDER BY name COLLATE NOCASE`
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("store: listing taxes: %w", err)
	}
	defer rows.Close()

	var out []Tax
	for rows.Next() {
		var (
			t    Tax
			rate int64
		)
		if err := rows.Scan(&t.ID, &t.Name, &rate, &t.Archived); err != nil {
			return nil, fmt.Errorf("store: reading tax: %w", err)
		}
		t.Rate = domain.Rate(rate)
		out = append(out, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("store: listing taxes: %w", err)
	}
	return out, nil
}

// TaxByID loads one rate.
func (s *Store) TaxByID(ctx context.Context, id int64) (Tax, error) {
	const q = `SELECT id, name, rate_bp, archived FROM taxes WHERE id = ?`
	var (
		t    Tax
		rate int64
	)
	err := s.db.QueryRowContext(ctx, q, id).Scan(&t.ID, &t.Name, &rate, &t.Archived)
	if errors.Is(err, sql.ErrNoRows) {
		return Tax{}, ErrNotFound
	}
	if err != nil {
		return Tax{}, fmt.Errorf("store: reading tax %d: %w", id, err)
	}
	t.Rate = domain.Rate(rate)
	return t, nil
}

// CreateItem adds a catalogue entry.
func (s *Store) CreateItem(ctx context.Context, it Item) (int64, error) {
	const q = `INSERT INTO items (name, description, unit_price_minor, tax_id, archived)
	           VALUES (?, ?, ?, ?, ?)`
	res, err := s.db.ExecContext(ctx, q,
		it.Name, it.Description, int64(it.UnitPrice), nullID(it.TaxID), it.Archived)
	if err != nil {
		return 0, fmt.Errorf("store: creating item: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("store: creating item: %w", err)
	}
	return id, nil
}

// ListItems returns the catalogue entries available to fill a line.
func (s *Store) ListItems(ctx context.Context) ([]Item, error) {
	const q = `SELECT id, name, description, unit_price_minor, tax_id, archived
	           FROM items WHERE archived = 0 ORDER BY name COLLATE NOCASE`
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("store: listing items: %w", err)
	}
	defer rows.Close()

	var out []Item
	for rows.Next() {
		var (
			it    Item
			price int64
			taxID sql.NullInt64
		)
		if err := rows.Scan(&it.ID, &it.Name, &it.Description, &price, &taxID, &it.Archived); err != nil {
			return nil, fmt.Errorf("store: reading item: %w", err)
		}
		it.UnitPrice = domain.Money(price)
		it.TaxID = idFromNull(taxID)
		out = append(out, it)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("store: listing items: %w", err)
	}
	return out, nil
}
