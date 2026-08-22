package mkiln

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Platform struct {
	OS     string
	Distro string
	IDLike []string
}

func detectPlatform() (Platform, error) {
	p := Platform{OS: runtime.GOOS}
	if p.OS != "linux" {
		return p, nil
	}
	id, like, err := readOSRelease("/etc/os-release")
	if err != nil {
		return Platform{}, fmt.Errorf("detect Linux distribution: %w", err)
	}
	p.Distro, p.IDLike = id, like
	return p, nil
}

func readOSRelease(path string) (string, []string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()
	values := map[string]string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		key, value, ok := strings.Cut(s.Text(), "=")
		if ok && (key == "ID" || key == "ID_LIKE") {
			values[key] = strings.Trim(strings.TrimSpace(value), `"'`)
		}
	}
	if err := s.Err(); err != nil {
		return "", nil, err
	}
	return strings.ToLower(values["ID"]), strings.Fields(strings.ToLower(values["ID_LIKE"])), nil
}
