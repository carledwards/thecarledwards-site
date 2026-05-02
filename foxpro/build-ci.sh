#!/bin/sh
# build-ci.sh — installs Go 1.22 on demand and runs the wasm build.
#
# Cloudflare Pages (and similar CI environments) ship a Go that's
# older than what our go.mod requires. Go's automatic toolchain
# fetcher would try to download 1.22 from `dl.google.com` but is
# blocked in many build sandboxes, so we install it ourselves to
# /tmp/go and put it on PATH.
#
# Locally (where Go 1.22+ is already installed) this script is a
# no-op — the version check passes and we go straight to `make build`.

set -e

GO_REQUIRED_MAJOR=1
GO_REQUIRED_MINOR=22
GO_INSTALL_VERSION=1.22.12
GO_INSTALL_DIR=/tmp/go
GO_TARBALL_URL="https://go.dev/dl/go${GO_INSTALL_VERSION}.linux-amd64.tar.gz"

need_install=1
if command -v go >/dev/null 2>&1; then
    ver=$(go version | awk '{print $3}' | sed 's/^go//')
    major=$(echo "$ver" | cut -d. -f1)
    minor=$(echo "$ver" | cut -d. -f2)
    if [ "$major" -gt "$GO_REQUIRED_MAJOR" ] || \
       { [ "$major" -eq "$GO_REQUIRED_MAJOR" ] && [ "$minor" -ge "$GO_REQUIRED_MINOR" ]; }; then
        need_install=0
    fi
fi

if [ "$need_install" = "1" ]; then
    if [ ! -x "$GO_INSTALL_DIR/bin/go" ]; then
        echo "Installing Go ${GO_INSTALL_VERSION} to ${GO_INSTALL_DIR}…"
        rm -rf "$GO_INSTALL_DIR"
        curl -fsSL "$GO_TARBALL_URL" | tar -xz -C /tmp
    else
        echo "Reusing cached Go at ${GO_INSTALL_DIR}/bin"
    fi
    export PATH="${GO_INSTALL_DIR}/bin:$PATH"
fi

echo "Using: $(go version)"

cd "$(dirname "$0")"
exec make build
