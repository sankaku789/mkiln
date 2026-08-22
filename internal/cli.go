package mkiln

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	Input  string
	Output string
	Style  string
	Typst  bool
}

var (
	errHelp    = errors.New("help requested")
	errVersion = errors.New("version requested")
)

func parseArgs(args []string) (Options, error) {
	var opts Options
	var inputs []string
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-h", "--help":
			return Options{}, errHelp
		case "-V", "--version":
			return Options{}, errVersion
		case "-t", "--typst":
			opts.Typst = true
		case "-o", "--output", "-s", "--style":
			if i+1 >= len(args) {
				return Options{}, fmt.Errorf("option %s requires a value", arg)
			}
			i++
			if arg == "-o" || arg == "--output" {
				opts.Output = args[i]
			} else {
				opts.Style = args[i]
			}
		default:
			if value, ok := strings.CutPrefix(arg, "--output="); ok {
				opts.Output = value
			} else if value, ok := strings.CutPrefix(arg, "--style="); ok {
				opts.Style = value
			} else if strings.HasPrefix(arg, "-") {
				return Options{}, fmt.Errorf("unknown option %s", arg)
			} else {
				inputs = append(inputs, arg)
			}
		}
	}
	if len(inputs) != 1 {
		return Options{}, fmt.Errorf("expected exactly one input file")
	}
	opts.Input = inputs[0]
	if opts.Typst && opts.Style != "" {
		return Options{}, fmt.Errorf("style cannot be used with Typst conversion")
	}
	if info, err := os.Stat(opts.Input); err != nil {
		return Options{}, fmt.Errorf("input %q: %w", opts.Input, err)
	} else if info.IsDir() {
		return Options{}, fmt.Errorf("input %q is not a file", opts.Input)
	}
	return opts, nil
}

const usage = `Usage:
  mkiln FILE [-o PATH] [-s NAME]
  mkiln FILE --typst [-o PATH]
  mkiln setup

Options:
  -o, --output PATH  output path
  -s, --style NAME   HTML CSS style
  -t, --typst        convert to plain Typst source
  -h, --help         show help
  -V, --version      show version
`
