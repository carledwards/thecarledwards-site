# thecarledwards-site / foxpro — build the FoxPro Mode wasm.
#
#   make build    — rebuild sim.wasm + copy wasm_exec.js + foxpro.js
#                   into ../static/foxpro/
#   make clean    — remove built artifacts
#
# Astro is configured (in astro.config.mjs) to use ./static as its
# publicDir, so files in ../static/ are served as-is at the site root.
# The wasm + JS shims land at /foxpro/{sim.wasm,wasm_exec.js,foxpro.js}
# once `astro dev` or `astro build` runs.

PKG    := github.com/carledwards/thecarledwards-site/foxpro
OUT    := ../static/foxpro
WASM   := $(OUT)/sim.wasm
EXEC   := $(OUT)/wasm_exec.js
FOXPRO := $(OUT)/foxpro.js

GOROOT := $(shell go env GOROOT)
EXEC_SRC := $(firstword $(wildcard $(GOROOT)/lib/wasm/wasm_exec.js $(GOROOT)/misc/wasm/wasm_exec.js))
.PHONY: build clean wasm

build: wasm $(EXEC) $(FOXPRO)

wasm:
	@mkdir -p $(OUT)
	GOOS=js GOARCH=wasm go build -o $(WASM) $(PKG)
	@ls -lh $(WASM) | awk '{print "  built " $$NF " (" $$5 ")"}'

$(EXEC):
	@mkdir -p $(OUT)
	@if [ -z "$(EXEC_SRC)" ]; then \
	  echo "ERROR: wasm_exec.js not found under $(GOROOT)/{lib,misc}/wasm/"; \
	  exit 1; \
	fi
	cp $(EXEC_SRC) $@

# foxpro-go ships the shared canvas-renderer + input bridge in
# wasm/foxpro.js; resolve from the module cache so we always pick
# up the version matching go.mod's pinned tag. Both `go mod download`
# and `go list -m` are run in the same shell so the cache is
# guaranteed populated before the path lookup. Doing this in a make
# variable (`:=` or `=`) failed on a fresh CI runner because the
# module wasn't yet in the cache when the variable expanded.
$(FOXPRO):
	@mkdir -p $(OUT)
	@set -e; \
	go mod download github.com/carledwards/foxpro-go; \
	src="$$(go list -m -f '{{.Dir}}' github.com/carledwards/foxpro-go)/wasm/foxpro.js"; \
	if [ ! -f "$$src" ]; then \
	  echo "ERROR: foxpro.js not found at $$src"; \
	  exit 1; \
	fi; \
	cp "$$src" $@

clean:
	rm -f $(WASM) $(EXEC) $(FOXPRO)
