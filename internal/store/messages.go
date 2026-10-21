package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/tounsils/i-CRM/internal/domain"
)

// CreateMessage appends an entry to a customer's correspondence timeline.
func (s *Store) CreateMessage(ctx context.Context, m Message) (int64, error) {
	const q = `INSERT INTO messages
	           (customer_id, document_id, author_id, direction, subject, body, created_at)
	           VALUES (?, ?, ?, ?, ?, ?, ?)`
	res, err := s.db.ExecContext(ctx, q, m.CustomerID, nullID(m.DocumentID), m.AuthorID,
		string(m.Direction), m.Subject, m.Body, formatTime(time.Now()))
	if err != nil {
		return 0, fmt.Errorf("store: creating message: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("store: creating message: %w", err)
	}
	return id, nil
}

// MessagesForCustomer returns the timeline newest first, with author names.
func (s *Store) MessagesForCustomer(ctx context.Context, customerID int64) ([]Message, error) {
	const q = `SELECT m.id, m.customer_id, m.document_id, m.author_id, m.direction,
	                  m.subject, m.body, m.created_at, u.name
	           FROM messages m
	           JOIN users u ON u.id = m.author_id
	           WHERE m.customer_id = ?
	           ORDER BY m.created_at DESC, m.id DESC`
	rows, err := s.db.QueryContext(ctx, q, customerID)
	if err != nil {
		return nil, fmt.Errorf("store: listing messages for customer %d: %w", customerID, err)
	}
	defer rows.Close()

	var out []Message
	for rows.Next() {
		m, err := scanMessage(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("store: listing messages for customer %d: %w", customerID, err)
	}
	return out, nil
}

func scanMessage(row rowScanner) (Message, error) {
	var (
		m         Message
		docID     sql.NullInt64
		direction string
		createdAt string
	)
	err := row.Scan(&m.ID, &m.CustomerID, &docID, &m.AuthorID, &direction,
		&m.Subject, &m.Body, &createdAt, &m.AuthorName)
	if err != nil {
		return Message{}, fmt.Errorf("store: reading message: %w", err)
	}
	m.DocumentID = idFromNull(docID)
	m.Direction = domain.MessageDirection(direction)
	if m.CreatedAt, err = parseTime(createdAt); err != nil {
		return Message{}, fmt.Errorf("store: reading message timestamp: %w", err)
	}
	return m, nil
}
