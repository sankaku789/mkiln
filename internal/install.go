package mkiln

import (
	"fmt"
	"os"
	"os/exec"
)

type Installer struct{ Name string }
type Command struct {
	Name string
	Args []string
}

func setupPandoc() error {
	if _, err := exec.LookPath("pandoc"); err == nil {
		return nil
	}
	p, err := detectPlatform()
	if err != nil {
		return err
	}
	i, err := detectInstaller(p)
	if err != nil {
		return err
	}
	if err := requireSetupPrivileges(p, i); err != nil {
		return err
	}
	c, err := pandocInstallCommand(i)
	if err != nil {
		return err
	}
	return runInstaller(c)
}

func detectInstaller(p Platform) (Installer, error) {
	candidates := []string{}
	switch p.OS {
	case "darwin":
		candidates = []string{"brew"}
	case "windows":
		candidates = []string{"winget"}
	case "linux":
		family := append([]string{p.Distro}, p.IDLike...)
		for _, id := range family {
			switch id {
			case "debian", "ubuntu":
				candidates = append(candidates, "apt-get")
			case "arch":
				candidates = append(candidates, "pacman")
			case "fedora", "rhel":
				candidates = append(candidates, "dnf")
			}
		}
	}
	for _, name := range candidates {
		if _, err := exec.LookPath(name); err == nil {
			return Installer{Name: name}, nil
		}
	}
	return Installer{}, fmt.Errorf("no supported package manager found; install Pandoc manually")
}

func requireSetupPrivileges(p Platform, i Installer) error {
	if p.OS == "linux" && os.Geteuid() != 0 {
		return fmt.Errorf("%s requires root privileges; run `sudo mkiln setup`", i.Name)
	}
	return nil
}

func pandocInstallCommand(i Installer) (Command, error) {
	switch i.Name {
	case "apt-get":
		return Command{i.Name, []string{"install", "-y", "pandoc"}}, nil
	case "pacman":
		return Command{i.Name, []string{"-S", "--noconfirm", "pandoc"}}, nil
	case "dnf":
		return Command{i.Name, []string{"install", "-y", "pandoc"}}, nil
	case "brew":
		return Command{i.Name, []string{"install", "pandoc"}}, nil
	case "winget":
		return Command{i.Name, []string{"install", "--id", "JohnMacFarlane.Pandoc", "--exact"}}, nil
	default:
		return Command{}, fmt.Errorf("unsupported installer %q", i.Name)
	}
}

func runInstaller(c Command) error { return runCommand(c.Name, c.Args) }
