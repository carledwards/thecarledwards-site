#!/bin/sh
# build-ci.sh — installs Go 1.26 on demand and builds the CodeLab wasms.
#
# Mirrors foxpro/build-ci.sh, with two differences:
#   - go6sim/go6asm require Go 1.26 (foxpro only needs 1.22), so we
#     install a newer toolchain.
#   - we install to a SEPARATE dir and version-check the cached binary,
#     so we never reuse foxpro's 1.22 in /tmp/go (foxpro's script only
#     checks existence, not version).
#
# Cloudflare Pages ships an older Go and blocks Go's toolchain fetcher,
# so we fetch it ourselves. Locally (Go 1.26+ already on PATH) this is
# a no-op and we go straight to `make build`.

set -e

GO_REQUIRED_MAJOR=1
GO_REQUIRED_MINOR=26
GO_INSTALL_VERSION=1.26.2
GO_INSTALL_DIR=/tmp/go-1.26
GO_TARBALL_URL="https://go.dev/dl/go${GO_INSTALL_VERSION}.linux-amd64.tar.gz"

ge_required() { # $1=major $2=minor -> 0 if >= required
    [ "$1" -gt "$GO_REQUIRED_MAJOR" ] || \
    { [ "$1" -eq "$GO_REQUIRED_MAJOR" ] && [ "$2" -ge "$GO_REQUIRED_MINOR" ]; }
}

use_go() { # echo a usable `go` (>= required) or empty
    for go in go "$GO_INSTALL_DIR/bin/go"; do
        command -v "$go" >/dev/null 2>&1 || continue
        v=$("$go" version | awk '{print $3}' | sed 's/^go//')
        if ge_required "$(echo "$v" | cut -d. -f1)" "$(echo "$v" | cut -d. -f2)"; then
            echo "$go"; return 0
        fi
    done
    return 1
}

if ! GO_BIN=$(use_go); then
    if [ ! -x "$GO_INSTALL_DIR/bin/go" ]; then
        echo "Installing Go ${GO_INSTALL_VERSION} to ${GO_INSTALL_DIR}…"
        rm -rf "$GO_INSTALL_DIR"
        mkdir -p "$GO_INSTALL_DIR"
        curl -fsSL "$GO_TARBALL_URL" | tar -xz -C "$GO_INSTALL_DIR" --strip-components=1
    fi
    export PATH="${GO_INSTALL_DIR}/bin:$PATH"
fi

echo "Using: $(go version)"

cd "$(dirname "$0")"
exec make build
