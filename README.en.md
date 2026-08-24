# mkiln

[日本語](README.md) | [简体中文](README.cn.md)

mkiln is a thin Go CLI that uses Pandoc to convert Markdown into self-contained HTML or plain Typst source. Markdown parsing, math, syntax highlighting, and writers are delegated to Pandoc.

## Purpose

I wanted a tool that simply converts Markdown to HTML or Typst, but Pandoc has so many features that configuring its options can be somewhat cumbersome. I therefore created a thin wrapper that provides straightforward Markdown-to-HTML and Markdown-to-Typst commands.

## Features

- Converts Markdown to standalone HTML or Typst.
- Standalone HTML uses MathML for math and provides a simple viewer.

## Requirements

- Go 1.24.4 or later when building from source
- Pandoc

## Installation

```bash
go install github.com/sankaku789/mkiln/cmd/mkiln@latest
```

To build from the repository:

```bash
go build ./cmd/mkiln
```

## Usage

### Setup

If Pandoc is not installed, you can install it with the `setup` command.

```bash
mkiln setup
```

Installation uses the operating system's default package manager, such as winget, apt, or pacman.

### Conversion

```bash
mkiln note.md [-o PATH] [-s NAME]
```

This generates `note.html`. CSS and the outline are embedded in the HTML, so the resulting single file can be moved, shared, and viewed offline.

The `-o`, `--output` option specifies the output file name and path.

The `-s`, `--style` option selects a CSS file. When omitted, the provided default CSS is used.

```bash
mkiln FILE --typst [-o PATH]
```

The `-t`, `--typst` option outputs Typst source.

## Configuration

On the first HTML build, mkiln creates the following under the user's configuration directory:

```text
mkiln/
├── default.yaml
└── styles/
    └── default.css
```

- `default.yaml` is a Pandoc defaults file, not a custom mkiln format.
- Select `styles/NAME.css` with `--style NAME`.
- Existing files are not overwritten automatically.
- `--typst` does not create the configuration directory.

## Development

```bash
gofmt -w cmd internal
go test ./...
go vet ./...
go build ./cmd/mkiln
```

## References

The CSS layout is based on the [Digital Agency Design System](https://design.digital.go.jp/dads/) and reuses selected components from its [example implementation](https://github.com/digital-go-jp/design-system-example-components-html).

## License

Provided under the MIT License.
