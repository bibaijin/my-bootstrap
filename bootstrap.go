package main

import (
	"os"
	"os/exec"
	"strings"
)

var (
	toInstallSoftwares []string
	xdgSoftwares       []string // 支持 XDG Base Directory
)

func init() {
	xdgSoftwares = []string{
		"fish",
		"git",
	}
}

func bootstrap(os string) {
	infof("Bootstraping %s...\n", os)

	infof("Installing packages...")
	for _, software := range toInstallSoftwares {
		infof("\tInstalling %s...", software)
		if err := install(software); err != nil {
			fatalf("\tInstall %s failed, error: %s.", software, err)
		} else {
			infof("\t%s installed.", software)
		}
	}
	infof("Packages installed.\n")

	infof("chsh -s $(which fish)...")
	if err := chsh(); err != nil {
		fatalf("chsh -s $(which fish) failed, error: %s.\n", err)
	} else {
		infof("chsh -s $(which fish) done.\n")
	}

	infof("Configuring packages support XDG base directory...")
	for _, software := range xdgSoftwares {
		infof("Configuring %s...", software)
		if err := config(software); err != nil {
			fatalf("Configure %s failed, error: %s.\n", software, err)
		} else {
			infof("%s configured.\n", software)
		}
	}
	infof("Packages support XDG base directory configured.")

	infof("%s bootstraped.", os)
}

func chsh() error {
	fish, err := which("fish")
	if err != nil {
		return err
	}

	if strings.Contains(os.Getenv("SHELL"), "fish") {
		warnf("\tfish is already the login shell, skipped.")
		return nil
	}

	cmd := exec.Command("chsh", "-s", fish)
	if err := runCmd(cmd); err != nil {
		return err
	}

	return nil
}

func config(software string) error {
	if err := symlink("/config/fish", "/.config/"+software); err != nil {
		return err
	}

	return nil
}
