# i-CRM

Self-hostable CRM for small firms that invoice: leads and customers, proposals,
estimates, invoices, line items, tax rates, payments and a correspondence
timeline.

One static binary and one SQLite file. No runtime to install, no database
server to administer, no outbound network calls, and no JavaScript. The
migrations, HTML templates and stylesheet are compiled into the binary, so a
deployment is a file and a data directory.

## Status

This is a working application, not a scaffold. It is also young, and the list
of what it does not do is as important as the list of what it does.

**Implemented and covered by tests**

- First-run administrator setup, bcrypt sign-in, signed session cookies
- Leads and customers as one record with a stage, filterable
- Per-customer correspondence timeline: sent, received and internal notes
- Proposals, estimates and invoices sharing one lifecycle and one numbering scheme
- Line items with fractional quantities, per-line tax, and totals that cannot
  disagree with their lines
- Named tax rates, and a reusable item catalogue for filling lines
- A status machine that refuses illegal moves, so a paid invoice cannot be
  walked back to draft and edited
- Conversion from proposal to estimate to invoice, carrying the lines across
- Payments against sent invoices, with paid and partly paid derived from the
  payments on file rather than chosen by a click
- A dashboard showing what is outstanding and what moved recently

**Not implemented**

- No PDF rendering. Documents are screens, not files, for now.
- No email is sent. The correspondence timeline records what was said; it does
  not deliver it.
- One administrator account. There is no user management screen and no roles.
- No recurring invoices, credit notes, or currency conversion. A document holds
  a currency code and arithmetic stays within it.
- No reporting beyond the dashboard, no attachments, and no HTTP API.

## Running it

### Docker, which is the supported path

```sh
cp .env.example .env
# Put a long random value in ICRM_SESSION_KEY. For example:
#   openssl rand -base64 48
docker compose up --build -d
```

Open <http://localhost:8080>. The first visit offers to create the
administrator account, and that page closes permanently once it succeeds.

### From source

```sh
go build -o bin/icrm ./cmd/icrm
./bin/icrm
```

Go 1.26 or newer. There is nothing else to install: the SQLite driver is pure
Go, so the build needs no C toolchain and `CGO_ENABLED=0` produces a static
binary.

## Configuration

Everything is an environment variable, and `internal/config` is the only place
a default is written down.

| Variable | Default | What it does |
| --- | --- | --- |
| `ICRM_SESSION_KEY` | none | Signs session cookies. At least 32 characters. Required when `ICRM_ENV=production`; otherwise a throwaway key is generated at boot and everyone is signed out on restart. |
| `ICRM_ENV` | `development` | `production` marks session cookies `Secure`. See the warning below. |
| `ICRM_ADDR` | `:8080` | Listen address. |
| `ICRM_DB` | `data/icrm.db` | Path to the SQLite file. The container sets `/data/icrm.db`. |
| `ICRM_COMPANY_NAME` | `i-CRM` | Shown in the title and top bar. |
| `ICRM_COMPANY_ADDRESS` | empty | Your postal address, printed on documents. Use `\n` between lines, since a `.env` file cannot hold a real newline. |
| `ICRM_CURRENCY` | `USD` | Default currency for new records. |
| `ICRM_PAYMENT_TERMS_DAYS` | `30` | Added to an invoice's issue date to get its due date. |

**Set `ICRM_ENV=production` only once you are behind HTTPS.** A `Secure` cookie
is never sent over plain HTTP, so turning this on while serving over HTTP makes
sign-in silently impossible. That is why the default is `development` even
though it is the less hardened of the two.

There is no TLS in this binary. Put it behind a reverse proxy that terminates
TLS before exposing it to anything other than a trusted network.

## How it is put together

```
cmd/icrm          entrypoint: config, store, server, graceful shutdown
internal/domain   business rules. No database, no HTTP, no templates.
internal/store    SQLite, embedded migrations, one method per query
internal/web      routing, sessions, CSRF, handlers, embedded templates
```

`internal/domain` is where the rules that cost money live, and it imports
nothing from the rest of the application, so they can be tested directly.

Four decisions worth knowing before changing anything:

**Money never touches a float.** Amounts are `int64` minor units, quantities are
thousandths of a unit, and tax rates are basis points. Binary floating point
cannot represent 0.10, and the error compounds once a rate and a quantity are
applied to it.

**Tax is rounded per line, then summed.** Taxing the subtotal once gives a
different answer, by up to one minor unit per line. Per line is the rule here
because the tax printed next to a line has to be the tax that was added to the
total: a customer recomputing the invoice by hand must get the same number. The
test that pins this is `TestTaxIsRoundedPerLineNotOnTheSubtotal`.

**The router denies by default.** Public routes are enumerated explicitly and
everything else falls through the authentication gate, which also enforces CSRF
on every write. A new route is therefore private unless someone deliberately
publishes it, so the cost of forgetting is a locked door rather than an open
one. `TestEveryPrivateRouteRefusesAnUnauthenticatedRequest` asserts this over
every private route, and `TestUnknownPathsArePrivateByDefault` covers the paths
that are not listed anywhere.

**A tax rate is copied onto a line, not referenced.** Editing a rate next year
must not restate an invoice that was already sent.

## Development

```sh
make gates         # format check, build, lint, test
make docker-gates  # the same, inside a container, if you have no Go toolchain
make test
make lint
make run
```

All four gates pass on `main`. The test suite reports no skips; if you add a
skip, say so, because a skipped test is a failing test nobody can see.

## Backing it up

Everything is in the one SQLite file named by `ICRM_DB`, plus its `-wal`
sidecar. With the container running:

```sh
docker compose exec icrm sh -c 'cp /data/icrm.db /data/backup.db'
docker compose cp icrm:/data/backup.db ./icrm-backup.db
```

That file is every customer record and every invoice. Treat it accordingly.

## Licence

MIT. See [LICENSE](LICENSE).
