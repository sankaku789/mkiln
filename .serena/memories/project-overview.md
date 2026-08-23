# Project Overview

## 目的
mkiln is a thin Go CLI around Pandoc. Its primary path converts Markdown with LaTeX math to standalone HTML using Pandoc defaults, CSS, and KaTeX. `--typst` emits plain Typst source; `--pdf` emits DSDocTemplate-based Typst source and compiles a sibling PDF.

## 技術スタック
- Go 1.24.4, standard library only.
- External runtime executable: Pandoc.
- Embedded defaults and CSS via `go:embed`.

## 主要構成
- `cmd/mkiln/main.go` — executable entry point and build-time version.
- `internal/cli.go`, `internal/run.go` — exact v1 CLI and explicit HTML/Typst dispatch.
- `internal/pandoc.go` — build resolution, exact Pandoc arguments, subprocess execution.
- `internal/config.go`, `internal/assets` — non-overwriting user config expansion.
- `internal/platform.go`, `internal/install.go` — setup platform/package-manager handling.
- `internal/mkiln_test.go` — contract tests.

## よく使うコマンド
- Format: `gofmt -w cmd internal`
- Test: `go test ./...`
- Static checks: `go vet ./...`
- Build: `go build ./cmd/mkiln`

## 重要な境界・規約
- `-t`/`--typst` is a bool; never add generic targets or expose other Pandoc writers.
- HTML and Typst remain separate concrete flows. Typst never reads or creates mkiln HTML config.
- Pandoc owns Markdown/YAML/rendering behavior; no parser dependencies.
- Existing config is never overwritten. Setup never invokes sudo.
- Assets live at `internal/assets`, not root `assets`, because Go embed patterns cannot traverse from `internal/assets.go` to a parent directory.

See `mem:core` for the memory graph and durable invariants.