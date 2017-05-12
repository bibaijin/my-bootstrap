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

func bootstrap() {
	infof("Bootstraping linux...\n")

	infof("Installing packages...")
	for _, software := range softwares {
		infof("\tInstalling %s...", software)
		if err := install(software); err != nil {
			fatalf("\tInstall %s failed, error: %s.", software, err)
		} else {
			infof("\tInstall %s done.", software)
		}
	}
	infof("Packages installed.\n")

	infof("chsh -s $(which fish)...")
	if err := chsh(); err != nil {
		fatalf("chsh -s $(which fish) failed, error: %s.\n", err)
	} else {
		infof("chsh -s $(which fish) done.\n")
	}

	infof("Configuring fish...")
	if err := configFish(); err != nil {
		fatalf("Configure fish failed, error: %s.\n", err)
	} else {
		infof("Configure fish done.\n")
	}

	infof("Bootstrap linux done.")
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
