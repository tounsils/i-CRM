-- Money is stored in minor units (INTEGER), quantities in thousandths and tax
-- rates in basis points, matching domain.Money, domain.Quantity and
-- domain.Rate. No column in this schema holds a REAL, because no total in this
-- product may ever pass through binary floating point.
--
-- Timestamps are RFC3339 TEXT in UTC. SQLite has no date type, and text in a
-- fixed format sorts correctly and reads correctly in a shell.

CREATE TABLE users (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    email         TEXT    NOT NULL UNIQUE,
    name          TEXT    NOT NULL,
    password_hash TEXT    NOT NULL,
    role          TEXT    NOT NULL DEFAULT 'staff',
    created_at    TEXT    NOT NULL
);

CREATE TABLE customers (
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    name         TEXT    NOT NULL,
    contact_name TEXT    NOT NULL DEFAULT '',
    email        TEXT    NOT NULL DEFAULT '',
    phone        TEXT    NOT NULL DEFAULT '',
    address      TEXT    NOT NULL DEFAULT '',
    -- lead | active | archived, validated by domain.ParseCustomerStatus
    status       TEXT    NOT NULL DEFAULT 'lead',
    currency     TEXT    NOT NULL DEFAULT 'USD',
    notes        TEXT    NOT NULL DEFAULT '',
    created_at   TEXT    NOT NULL
);

CREATE INDEX idx_customers_status ON customers (status);

CREATE TABLE taxes (
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    name     TEXT    NOT NULL,
    rate_bp  INTEGER NOT NULL,
    archived INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE items (
    id               INTEGER PRIMARY KEY AUTOINCREMENT,
    name             TEXT    NOT NULL,
    description      TEXT    NOT NULL DEFAULT '',
    unit_price_minor INTEGER NOT NULL DEFAULT 0,
    tax_id           INTEGER REFERENCES taxes (id) ON DELETE SET NULL,
    archived         INTEGER NOT NULL DEFAULT 0
);

-- Proposals, estimates and invoices share this table. They differ by `type`
-- and by which statuses their lifecycle permits, which is enforced in
-- domain.Transition rather than here: SQLite cannot express a per-type status
-- machine, and splitting the rule across both places would give it two owners.
CREATE TABLE documents (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    type            TEXT    NOT NULL,
    number          TEXT    NOT NULL,
    customer_id     INTEGER NOT NULL REFERENCES customers (id),
    status          TEXT    NOT NULL DEFAULT 'draft',
    currency        TEXT    NOT NULL DEFAULT 'USD',
    issue_date      TEXT    NOT NULL,
    due_date        TEXT,
    notes           TEXT    NOT NULL DEFAULT '',
    -- Cached totals. document_lines is the source of truth; these are written
    -- only by store.recomputeTotals, in the same transaction as any line
    -- change, so that list screens can show a total without loading lines.
    subtotal_minor  INTEGER NOT NULL DEFAULT 0,
    tax_minor       INTEGER NOT NULL DEFAULT 0,
    total_minor     INTEGER NOT NULL DEFAULT 0,
    -- Set when this document has been converted onward, so the chain from
    -- proposal to estimate to invoice stays walkable.
    converted_to_id INTEGER REFERENCES documents (id) ON DELETE SET NULL,
    created_at      TEXT    NOT NULL,
    UNIQUE (type, number)
);

CREATE INDEX idx_documents_customer ON documents (customer_id);
CREATE INDEX idx_documents_type_status ON documents (type, status);

-- One sequence per document type, so numbering is a counter rather than a
-- parse of the highest existing number. Deriving the next number by reading
-- the last one means a single hand-edited reference breaks numbering for good;
-- a counter cannot be confused by one, and UNIQUE (type, number) above stays
-- as the backstop.
CREATE TABLE document_counters (
    type     TEXT    PRIMARY KEY,
    next_seq INTEGER NOT NULL DEFAULT 1
);

CREATE TABLE document_lines (
    id               INTEGER PRIMARY KEY AUTOINCREMENT,
    document_id      INTEGER NOT NULL REFERENCES documents (id) ON DELETE CASCADE,
    position         INTEGER NOT NULL DEFAULT 0,
    description      TEXT    NOT NULL,
    quantity_milli   INTEGER NOT NULL DEFAULT 1000,
    unit_price_minor INTEGER NOT NULL DEFAULT 0,
    -- The tax name and rate are copied onto the line rather than referenced,
    -- so that editing a tax rate next year cannot restate an invoice that was
    -- already sent.
    tax_name         TEXT    NOT NULL DEFAULT '',
    tax_rate_bp      INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_document_lines_document ON document_lines (document_id, position);

CREATE TABLE payments (
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    document_id  INTEGER NOT NULL REFERENCES documents (id) ON DELETE CASCADE,
    amount_minor INTEGER NOT NULL,
    method       TEXT    NOT NULL DEFAULT '',
    reference    TEXT    NOT NULL DEFAULT '',
    paid_on      TEXT    NOT NULL,
    created_at   TEXT    NOT NULL
);

CREATE INDEX idx_payments_document ON payments (document_id);

CREATE TABLE messages (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    customer_id INTEGER NOT NULL REFERENCES customers (id) ON DELETE CASCADE,
    document_id INTEGER REFERENCES documents (id) ON DELETE SET NULL,
    author_id   INTEGER NOT NULL REFERENCES users (id),
    -- outbound | inbound | note
    direction   TEXT    NOT NULL DEFAULT 'note',
    subject     TEXT    NOT NULL DEFAULT '',
    body        TEXT    NOT NULL,
    created_at  TEXT    NOT NULL
);

CREATE INDEX idx_messages_customer ON messages (customer_id, created_at);
