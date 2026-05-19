# thecarledwards-site / codelab — build the two CodeLab wasms.
#
#   make build  — build core.wasm (go6sim) + go6asm.wasm (go6asm) at
#                 the versions pinned in go.mod, and copy wasm_exec.js,
#                 into ../static/codelab/  (served at /codelab/* by
#                 Astro, publicDir = ./static).
#   make clean  — remove built artifacts.
#
# Mirrors foxpro/Makefile. The two cmds are `package main` in the
# pinned dependency modules; we build them by import path from this
# module's build list (no clone, reproducible via go.sum).

OUT      := ../static/codelab
CORE     := $(OUT)/core.wasm
GOASM    := $(OUT)/go6asm.wasm
EXEC     := $(OUT)/wasm_exec.js

GOROOT   := $(shell go env GOROOT)
EXEC_SRC := $(firstword $(wildcard $(GOROOT)/lib/wasm/wasm_exec.js $(GOROOT)/misc/wasm/wasm_exec.js))

.PHONY: build clean
build: $(CORE) $(GOASM) $(EXEC)

$(CORE):
	@mkdir -p $(OUT)
	GOOS=js GOARCH=wasm go build -o $(CORE) github.com/carledwards/go6sim/cmd/6502-core-wasm
	@ls -lh $(CORE) | awk '{print "  built " $$NF " (" $$5 ")"}'

$(GOASM):
	@mkdir -p $(OUT)
	GOOS=js GOARCH=wasm go build -o $(GOASM) github.com/carledwards/go6asm/cmd/go6asm-wasm
	@ls -lh $(GOASM) | awk '{print "  built " $$NF " (" $$5 ")"}'

$(EXEC):
	@mkdir -p $(OUT)
	@if [ -z "$(EXEC_SRC)" ]; then \
	  echo "ERROR: wasm_exec.js not found under $(GOROOT)/{lib,misc}/wasm/"; \
	  exit 1; \
	fi
	cp $(EXEC_SRC) $@

clean:
	rm -f $(CORE) $(GOASM) $(EXEC)
