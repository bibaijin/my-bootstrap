package main

import (
	"os"
	"os/exec"
	"strings"
)

var (
	softwares []string
)

func init() {
	softwares = []string{
		"fish",
		"git",
		"vim",
		"emacs",
	}
}

func bootstrapLinux() {
	infof("Installing packages...")
	for _, software := range softwares {
		_, err := which(software)
		if err == nil {
			infof("\tSkip %s, because it is installed.", software)
			continue
		}

		infof("\tpacman -S --noconfirm %s...", software)
		if err := install(software); err != nil {
			errorf("\tpacman -S --noconfirm %s failed, error: %s.", software, err)
		} else {
			infof("\tpacman -S --noconfirm %s done.", software)
		}
	}
	infof("Packages installed.\n")

	infof("chsh -s $(which fish)...")
	if err := chsh(); err != nil {
		errorf("chsh -s $(which fish) failed, error: %s.\n", err)
	} else {
		infof("chsh -s $(which fish) done.\n")
	}
}

func which(executable string) (string, error) {
	cmd := exec.Command("which", executable)
	out, err := cmd.CombinedOutput()
	if err != nil {
		errorf("\t%s", out)
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func install(software string) error {
	cmd := exec.Command("pacman", "-S", "--noconfirm", software)
	if err := runCmd(cmd); err != nil {
		return err
	}

	return nil
}

func chsh() error {
	fish, err := which("fish")
	if err != nil {
		return err
	}

	cmd := exec.Command("chsh", "-s", fish)
	if err := runCmd(cmd); err != nil {
		return err
	}

	return nil
}
