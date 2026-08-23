package mkiln

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

type PDFBuild struct {
	Input        string
	Output       string
	PDFOutput    string
	TemplatePath string
}

func findPandoc() (string, error) {
	path, err := exec.LookPath("pandoc")
	if err != nil {
		return "", fmt.Errorf("pandoc not found; run `mkiln setup`")
	}
	return path, nil
}

func findTypst() (string, error) {
	path, err := exec.LookPath("typst")
	if err != nil {
		return "", fmt.Errorf("typst not found; install Typst to use `mkiln --pdf`")
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

func htmlPandocArgs(b HTMLBuild, styleHeaderPath string) []string {
	return []string{b.Input, "--defaults", b.DefaultsPath, "--include-in-header", styleHeaderPath, "-o", b.Output}
}

func runHTMLPandoc(pandocPath string, b HTMLBuild) error {
	css, err := os.ReadFile(b.CSSPath)
	if err != nil {
		return fmt.Errorf("read style %q: %w", b.CSSPath, err)
	}
	outline, err := assets.ReadFile("assets/outline.js")
	if err != nil {
		return fmt.Errorf("read outline script: %w", err)
	}
	header, err := os.CreateTemp("", "mkiln-style-*.html")
	if err != nil {
		return fmt.Errorf("create temporary style header: %w", err)
	}
	headerPath := header.Name()
	defer os.Remove(headerPath)
	if _, err := fmt.Fprintf(header, "<style>\n%s\n</style>\n<script>\n%s\n</script>\n", css, outline); err != nil {
		header.Close()
		return fmt.Errorf("write temporary style header: %w", err)
	}
	if err := header.Close(); err != nil {
		return fmt.Errorf("close temporary style header: %w", err)
	}
	return runCommand(pandocPath, htmlPandocArgs(b, headerPath))
}

func resolveTypst(opts Options) (TypstBuild, error) {
	if opts.Style != "" {
		return TypstBuild{}, fmt.Errorf("style cannot be used with Typst conversion")
	}
	output := resolveOutput(opts.Input, opts.Output, true)
	if filepath.Ext(output) != ".typ" {
		return TypstBuild{}, fmt.Errorf("Typst output must use the .typ extension: %q", output)
	}
	return TypstBuild{opts.Input, output}, nil
}

func resolvePDF(opts Options) (PDFBuild, error) {
	if opts.Style != "" {
		return PDFBuild{}, fmt.Errorf("style cannot be used with PDF conversion")
	}
	output := resolveOutput(opts.Input, opts.Output, true)
	if filepath.Ext(output) != ".typ" {
		return PDFBuild{}, fmt.Errorf("PDF source output must use the .typ extension: %q", output)
	}
	template, err := resolveTypstTemplate()
	if err != nil {
		return PDFBuild{}, err
	}
	pdfOutput := strings.TrimSuffix(output, filepath.Ext(output)) + ".pdf"
	return PDFBuild{opts.Input, output, pdfOutput, template}, nil
}

func typstPandocArgs(b TypstBuild) []string {
	return []string{b.Input, "-t", "typst", "-o", b.Output}
}

func pdfPandocArgs(b PDFBuild) []string {
	return []string{b.Input, "-t", "typst", "--standalone", "--template", b.TemplatePath, "-o", b.Output}
}

func runTypstPandoc(pandocPath string, b TypstBuild) error {
	return runCommand(pandocPath, typstPandocArgs(b))
}

func runPDFPandoc(pandocPath string, b PDFBuild) error {
	return runCommand(pandocPath, pdfPandocArgs(b))
}

func runTypstCompile(typstPath string, b PDFBuild) error {
	return runCommand(typstPath, []string{"compile", b.Output, b.PDFOutput})
}

func runCommand(name string, args []string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s failed: %w", filepath.Base(name), err)
	}
	return nil
}
