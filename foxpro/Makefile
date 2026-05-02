# thecarledwards-site / foxpro — build the FoxPro Mode wasm.
#
#   make build    — rebuild sim.wasm + copy wasm_exec.js into ../static/foxpro/
#   make clean    — remove built artifacts
#
# Astro is configured (in astro.config.mjs) to use ./static as its
# publicDir, so files in ../static/ are served as-is at the site root.
# The wasm and JS shim end up at /foxpro/sim.wasm and /foxpro/wasm_exec.js
# once `astro dev` or `astro build` runs.

PKG  := github.com/carledwards/thecarledwards-site/foxpro
OUT  := ../static/foxpro
WASM := $(OUT)/sim.wasm
EXEC := $(OUT)/wasm_exec.js

GOROOT := $(shell go env GOROOT)
EXEC_SRC := $(firstword $(wildcard $(GOROOT)/lib/wasm/wasm_exec.js $(GOROOT)/misc/wasm/wasm_exec.js))

.PHONY: build clean wasm

# `go build` does its own change detection, so always invoke it
# rather than mirroring Go's import graph here.
build: wasm $(EXEC)

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

clean:
	rm -f $(WASM) $(EXEC)
