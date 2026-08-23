package mkiln

import (
	"errors"
	"fmt"
	"io"
)

func Run(args []string, version string, stdout, stderr io.Writer) int {
	if len(args) == 1 && args[0] == "setup" {
		if err := setupPandoc(); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		return 0
	}
	opts, err := parseArgs(args)
	if errors.Is(err, errHelp) {
		fmt.Fprint(stdout, usage)
		return 0
	}
	if errors.Is(err, errVersion) {
		fmt.Fprintln(stdout, version)
		return 0
	}
	if err != nil {
		fmt.Fprintln(stderr, err)
		fmt.Fprint(stderr, usage)
		return 2
	}
	pandoc, err := findPandoc()
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	if opts.Typst {
		build, err := resolveTypst(opts)
		if err == nil {
			err = runTypstPandoc(pandoc, build)
		}
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		return 0
	}
	if opts.PDF {
		if err := ensureTypstTemplate(); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		typst, err := findTypst()
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		build, err := resolvePDF(opts)
		if err == nil {
			err = runPDFPandoc(pandoc, build)
		}
		if err == nil {
			err = runTypstCompile(typst, build)
		}
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		return 0
	}
	if err := ensureUserConfig(); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	build, err := resolveHTML(opts)
	if err == nil {
		err = runHTMLPandoc(pandoc, build)
	}
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	return 0
}
