#!/bin/sh
# All four quality gates, in the order that fails cheapest first.
#
# Lives in a script rather than only in the Makefile so that the same commands
# run identically from make, from a container, and from CI, with no shell
# quoting differences between them.
#
# Usage: scripts/gates.sh [--with-lint]
#   --with-lint also runs golangci-lint, which must already be on PATH. The
#   Makefile runs it in its own container, so this script skips it by default.

set -eu

GO=${GO:-go}
status=0

echo "== format =="
unformatted=$(gofmt -l ./cmd ./internal)
if [ -n "$unformatted" ]; then
    echo "gofmt would rewrite:"
    echo "$unformatted"
    status=1
else
    echo "ok"
fi

echo
echo "== build =="
if $GO build ./...; then
    echo "ok"
else
    status=1
fi

echo
echo "== vet =="
if $GO vet ./...; then
    echo "ok"
else
    status=1
fi

echo
echo "== test =="
# -count=1 defeats the test cache, so a green run means it actually ran.
if ! $GO test ./... -count=1; then
    status=1
fi

if [ "${1:-}" = "--with-lint" ]; then
    echo
    echo "== lint =="
    if golangci-lint run ./...; then
        echo "ok"
    else
        status=1
    fi
fi

echo
if [ "$status" -eq 0 ]; then
    echo "all gates passed"
else
    echo "GATES FAILED"
fi
exit "$status"
