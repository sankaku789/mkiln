package mkiln

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type HTMLBuild struct {
	Input        string
	Output       string
	DefaultsPath string
	CSSPath      string
}

type TypstBuild struct {
	Input  string
	Output string
}

func findPandoc() (string, error) {
	path, err := exec.LookPath("pandoc")
	if err != nil {
		return "", fmt.Errorf("pandoc not found; run `mkiln setup`")
	}
	return path, nil
}

func resolveHTML(opts Options) (HTMLBuild, error) {
	defaults, err := resolveDefaults()
	if err != nil {
		return HTMLBuild{}, err
	}
	style := opts.Style
	if style == "" {
		style = "default"
	}
	css, err := resolveStyle(style)
	if err != nil {
		return HTMLBuild{}, err
	}
	return HTMLBuild{opts.Input, resolveOutput(opts.Input, opts.Output, false), defaults, css}, nil
}

func resolveDefaults() (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, "default.yaml")
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("defaults %q: %w", path, err)
	}
	return path, nil
}

func resolveStyle(name string) (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, "styles", name+".css")
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("style %q not found at %q: %w", name, path, err)
	}
	return path, nil
}

func htmlPandocArgs(b HTMLBuild) []string {
	return []string{b.Input, "--defaults", b.DefaultsPath, "--css", b.CSSPath, "-o", b.Output}
}

func runHTMLPandoc(pandocPath string, b HTMLBuild) error {
	return runCommand(pandocPath, htmlPandocArgs(b))
}

func resolveTypst(opts Options) (TypstBuild, error) {
	if opts.Style != "" {
		return TypstBuild{}, fmt.Errorf("style cannot be used with Typst conversion")
	}
	return TypstBuild{opts.Input, resolveOutput(opts.Input, opts.Output, true)}, nil
}

func typstPandocArgs(b TypstBuild) []string {
	return []string{b.Input, "-t", "typst", "-o", b.Output}
}

func runTypstPandoc(pandocPath string, b TypstBuild) error {
	return runCommand(pandocPath, typstPandocArgs(b))
}

func runCommand(name string, args []string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s failed: %w", filepath.Base(name), err)
	}
	return nil
}
