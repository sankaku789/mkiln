package mkiln

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

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

func TestPandocArgsRemainSeparate(t *testing.T) {
	html := HTMLBuild{"note.md", "note.html", "default.yaml", "default.css"}
	wantHTML := []string{"note.md", "--defaults", "default.yaml", "--include-in-header", "style.html", "-o", "note.html"}
	if got := htmlPandocArgs(html, "style.html"); !reflect.DeepEqual(got, wantHTML) {
		t.Fatalf("htmlPandocArgs() = %q, want %q", got, wantHTML)
	}
	typst := TypstBuild{"note.md", "note.typ", "note.pdf", "default.typ"}
	wantTypst := []string{"note.md", "-t", "typst", "--standalone", "--template", "default.typ", "-o", "note.typ"}
	if got := typstPandocArgs(typst); !reflect.DeepEqual(got, wantTypst) {
		t.Fatalf("typstPandocArgs() = %q, want %q", got, wantTypst)
	}
}

func TestResolveTypstRejectsNonTypOutput(t *testing.T) {
	_, err := resolveTypst(Options{Input: "note.md", Output: "note.pdf", Typst: true})
	if err == nil || !strings.Contains(err.Error(), ".typ") {
		t.Fatalf("resolveTypst() error = %v, want .typ extension error", err)
	}
}

func TestEnsureUserConfigDoesNotOverwrite(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
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

func TestRunTypstCreatesTemplateConfig(t *testing.T) {
	root := t.TempDir()
	configHome := filepath.Join(root, "config")
	binDir := filepath.Join(root, "bin")
	if err := os.Mkdir(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	pandoc := filepath.Join(binDir, "pandoc")
	if err := os.WriteFile(pandoc, []byte("#!/bin/sh\nexit 0\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	typst := filepath.Join(binDir, "typst")
	if err := os.WriteFile(typst, []byte("#!/bin/sh\nexit 0\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	input := filepath.Join(root, "note.md")
	if err := os.WriteFile(input, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", binDir)
	t.Setenv("XDG_CONFIG_HOME", configHome)
	var stdout, stderr bytes.Buffer
	if code := Run([]string{input, "--typst"}, "test", &stdout, &stderr); code != 0 {
		t.Fatalf("Run() = %d, stderr = %q", code, stderr.String())
	}
	template := filepath.Join(configHome, "mkiln", "templates", "default.typ")
	if _, err := os.Stat(template); err != nil {
		t.Fatalf("Typst template was not created: %v", err)
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
