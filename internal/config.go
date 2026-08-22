package mkiln

import (
	"fmt"
	"os"
	"path/filepath"
)

func configDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("resolve user config directory: %w", err)
	}
	return filepath.Join(dir, "mkiln"), nil
}

func ensureUserConfig() error {
	dir, err := configDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(dir, "styles"), 0o755); err != nil {
		return fmt.Errorf("create config directory: %w", err)
	}
	if err := writeEmbeddedFile("assets/default.yaml", filepath.Join(dir, "default.yaml")); err != nil {
		return err
	}
	return writeEmbeddedFile("assets/default.css", filepath.Join(dir, "styles", "default.css"))
}

func writeEmbeddedFile(src, dst string) error {
	data, err := assets.ReadFile(src)
	if err != nil {
		return fmt.Errorf("read embedded %q: %w", src, err)
	}
	f, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if errorsIsExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("create %q: %w", dst, err)
	}
	if _, err := f.Write(data); err != nil {
		f.Close()
		return fmt.Errorf("write %q: %w", dst, err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close %q: %w", dst, err)
	}
	return nil
}

func errorsIsExist(err error) bool { return os.IsExist(err) }
