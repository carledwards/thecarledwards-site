// Self-contained build module for the CodeLab wasm assets, mirroring
// foxpro/. It has no Go source of its own — it exists only to pin the
// versions the Makefile builds the two upstream wasm cmds from:
//
//	github.com/carledwards/go6sim/cmd/6502-core-wasm  -> core.wasm
//	github.com/carledwards/go6asm/cmd/go6asm-wasm      -> go6asm.wasm
//
// Hand-pinned (do NOT `go mod tidy` — nothing imports these, tidy would
// drop them). CI builds from these pinned tags; local dev may use a
// go.work / replace to iterate without pushing tags.
module github.com/carledwards/thecarledwards-site/codelab

go 1.26.2

require (
	github.com/carledwards/go6asm v0.1.1
	github.com/carledwards/go6sim v0.1.0
)
