# Build and run i-CRM as one static binary.
#
# CGO is off and the SQLite driver is pure Go, so the result has no shared
# library dependencies at all. That is what makes the self-hosting claim real:
# the image carries a binary, and the binary carries its own migrations,
# templates and stylesheet.

FROM golang:1.26-alpine AS build

WORKDIR /src

# Dependencies first, so editing application code does not re-download them.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Fail the build rather than ship an unverified image.
RUN go vet ./... && go test ./... -count=1

# -trimpath keeps build paths out of the binary; -s -w drop the symbol and DWARF
# tables, which is about a third of the size and nothing a production deployment
# reads.
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/icrm ./cmd/icrm


# Alpine rather than scratch: the binary needs nothing from it, but an operator
# debugging their own install wants a shell, and the healthcheck below wants
# wget. The cost is about 8 MB.
FROM alpine:3.21

RUN adduser -S -u 10001 icrm \
    && mkdir -p /data \
    && chown icrm /data

COPY --from=build /out/icrm /usr/local/bin/icrm

USER icrm

ENV ICRM_ADDR=:8080 \
    ICRM_DB=/data/icrm.db

EXPOSE 8080
VOLUME /data

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget -q -O- http://127.0.0.1:8080/healthz || exit 1

ENTRYPOINT ["/usr/local/bin/icrm"]
