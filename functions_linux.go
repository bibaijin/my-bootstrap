package main

import (
	"os/exec"
)

func init() {
	toInstallSoftwares = []string{
		"fish",
		"git",
		"vim",
		"emacs",
	}
}

func install(software string) error {
	_, err := which(software)
	if err == nil {
		warnf("\t%s is installed, skipped.", software)
		return nil
	}

	cmd := exec.Command("pacman", "-S", "--noconfirm", software)
	if err := runCmd(cmd); err != nil {
		return err
	}

	return nil
}
