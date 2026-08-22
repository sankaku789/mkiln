package mkiln

import (
	"path/filepath"
	"strings"
)

func resolveOutput(input, output string, typst bool) string {
	if output != "" {
		return output
	}
	ext := ".html"
	if typst {
		ext = ".typ"
	}
	inputExt := filepath.Ext(input)
	return strings.TrimSuffix(input, inputExt) + ext
}
