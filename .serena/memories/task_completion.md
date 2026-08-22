# Task completion

Run:

1. `gofmt -w cmd internal`
2. `go test ./...`
3. `go build ./cmd/mkiln`

For CLI behavior changes, also run the affected command with a controlled fake or installed Pandoc as appropriate.