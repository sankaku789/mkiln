package mkiln

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("MKILN_FAKE_COMMAND") == "1" {
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func setTestConfigHome(t *testing.T, dir string) {
	t.Helper()
	t.Setenv("XDG_CONFIG_HOME", dir)
	t.Setenv("APPDATA", dir)
}

func writeFakeCommand(t *testing.T, dir, name string) string {
	t.Helper()
	if runtime.GOOS == "windows" {
		name += ".exe"
	}

	srcPath, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	src, err := os.Open(srcPath)
	if err != nil {
		t.Fatal(err)
	}
	defer src.Close()

	dstPath := filepath.Join(dir, name)
	dst, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o755)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := io.Copy(dst, src); err != nil {
		dst.Close()
		t.Fatal(err)
	}
	if err := dst.Close(); err != nil {
		t.Fatal(err)
	}
	return dstPath
}

func TestParseArgsAllowsOptionsAfterInput(t *testing.T) {
	input := filepath.Join(t.TempDir(), "note.md")
	if err := os.WriteFile(input, []byte("# Note"), 0o644); err != nil {
		t.Fatal(err)
	}
	opts, err := parseArgs([]string{input, "-o", "out.html", "--style", "novel"})
	if err != nil {
		t.Fatal(err)
	}
	want := Options{Input: input, Output: "out.html", Style: "novel"}
	if !reflect.DeepEqual(opts, want) {
		t.Fatalf("parseArgs() = %#v, want %#v", opts, want)
	}
}

func TestParseArgsRejectsTypstStyle(t *testing.T) {
	input := filepath.Join(t.TempDir(), "note.md")
	if err := os.WriteFile(input, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := parseArgs([]string{input, "-t", "-s", "novel"})
	if err == nil || !strings.Contains(err.Error(), "style") {
		t.Fatalf("parseArgs() error = %v, want style error", err)
	}
}

func TestParseArgsRejectsTypstAndPDF(t *testing.T) {
	input := filepath.Join(t.TempDir(), "note.md")
	if err := os.WriteFile(input, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := parseArgs([]string{input, "--typst", "--pdf"})
	if err == nil || !strings.Contains(err.Error(), "cannot be used together") {
		t.Fatalf("parseArgs() error = %v, want mutual exclusion error", err)
	}
}

func TestPandocArgsRemainSeparate(t *testing.T) {
	html := HTMLBuild{"note.md", "note.html", "default.yaml", "default.css"}
	wantHTML := []string{"note.md", "--defaults", "default.yaml", "--include-in-header", "style.html", "-o", "note.html"}
	if got := htmlPandocArgs(html, "style.html"); !reflect.DeepEqual(got, wantHTML) {
		t.Fatalf("htmlPandocArgs() = %q, want %q", got, wantHTML)
	}
	typst := TypstBuild{"note.md", "note.typ"}
	wantTypst := []string{"note.md", "-t", "typst", "-o", "note.typ"}
	if got := typstPandocArgs(typst); !reflect.DeepEqual(got, wantTypst) {
		t.Fatalf("typstPandocArgs() = %q, want %q", got, wantTypst)
	}
	pdf := PDFBuild{"note.md", "note.typ", "note.pdf", "default.typ"}
	wantPDF := []string{"note.md", "-t", "typst", "--standalone", "--template", "default.typ", "-o", "note.typ"}
	if got := pdfPandocArgs(pdf); !reflect.DeepEqual(got, wantPDF) {
		t.Fatalf("pdfPandocArgs() = %q, want %q", got, wantPDF)
	}
}

func TestEmbeddedHTMLDefaultsUseMathML(t *testing.T) {
	defaults, err := assets.ReadFile("assets/default.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(defaults), "method: mathml") {
		t.Fatalf("default.yaml does not use MathML: %q", defaults)
	}
	if strings.Contains(string(defaults), "katex") || strings.Contains(string(defaults), "embed-resources") {
		t.Fatalf("default.yaml retains KaTeX resource configuration: %q", defaults)
	}
}

func TestEmbeddedDefaultCSSIncludesPrintProfile(t *testing.T) {
	data, err := assets.ReadFile("assets/default.css")
	if err != nil {
		t.Fatal(err)
	}
	css := string(data)
	for _, rule := range []string{
		"src/global.css",
		"src/components/link/link.css",
		"src/components/blockquote/blockquote.css",
		"src/components/image/image.css",
		"src/components/toc/toc.css",
		".storybook/prose.css",
		"@media print",
		"@page",
		"#TOC,",
		"break-after: avoid-page",
		"break-inside: avoid-page",
		"table-header-group",
		"tr {",
		"white-space: pre-wrap",
	} {
		if !strings.Contains(css, rule) {
			t.Errorf("default.css does not contain %q", rule)
		}
	}
}

func TestResolveTypstRejectsNonTypOutput(t *testing.T) {
	_, err := resolveTypst(Options{Input: "note.md", Output: "note.pdf", Typst: true})
	if err == nil || !strings.Contains(err.Error(), ".typ") {
		t.Fatalf("resolveTypst() error = %v, want .typ extension error", err)
	}
}

func TestEnsureUserConfigDoesNotOverwrite(t *testing.T) {
	setTestConfigHome(t, t.TempDir())
	if err := ensureUserConfig(); err != nil {
		t.Fatal(err)
	}
	dir, err := configDir()
	if err != nil {
		t.Fatal(err)
	}
	defaults := filepath.Join(dir, "default.yaml")
	if err := os.WriteFile(defaults, []byte("custom"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := ensureUserConfig(); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(defaults)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "custom" {
		t.Fatalf("existing defaults overwritten: %q", got)
	}
}

func TestRunTypstDoesNotCreateConfig(t *testing.T) {
	root := t.TempDir()
	configHome := filepath.Join(root, "config")
	binDir := filepath.Join(root, "bin")
	if err := os.Mkdir(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	writeFakeCommand(t, binDir, "pandoc")
	input := filepath.Join(root, "note.md")
	if err := os.WriteFile(input, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("MKILN_FAKE_COMMAND", "1")
	setTestConfigHome(t, configHome)
	var stdout, stderr bytes.Buffer
	if code := Run([]string{input, "--typst"}, "test", &stdout, &stderr); code != 0 {
		t.Fatalf("Run() = %d, stderr = %q", code, stderr.String())
	}
	if _, err := os.Stat(filepath.Join(configHome, "mkiln")); !os.IsNotExist(err) {
		t.Fatalf("Typst conversion created config directory: %v", err)
	}
}

func TestParseArgsRejectsDisabledPDF(t *testing.T) {
	input := filepath.Join(t.TempDir(), "note.md")
	if err := os.WriteFile(input, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := parseArgs([]string{input, "--pdf"})
	if err == nil || !strings.Contains(err.Error(), "currently disabled") {
		t.Fatalf("parseArgs() error = %v, want disabled error", err)
	}
}

func TestRunReportsMissingPandocWithoutInstalling(t *testing.T) {
	input := filepath.Join(t.TempDir(), "note.md")
	if err := os.WriteFile(input, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", t.TempDir())
	var stdout, stderr bytes.Buffer
	if code := Run([]string{input}, "test", &stdout, &stderr); code == 0 {
		t.Fatal("Run() succeeded without pandoc")
	}
	if want := "pandoc not found; run `mkiln setup`"; !strings.Contains(stderr.String(), want) {
		t.Fatalf("stderr = %q, want %q", stderr.String(), want)
	}
}

func TestReadOSRelease(t *testing.T) {
	path := filepath.Join(t.TempDir(), "os-release")
	data := "NAME=Example\nID=ubuntu\nID_LIKE=\"debian linux\"\n"
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	id, like, err := readOSRelease(path)
	if err != nil {
		t.Fatal(err)
	}
	if id != "ubuntu" || !reflect.DeepEqual(like, []string{"debian", "linux"}) {
		t.Fatalf("readOSRelease() = %q, %q", id, like)
	}
}
