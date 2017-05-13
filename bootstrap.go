package main

import (
	"os"
	"os/exec"
	"strings"
)

var (
	toInstallPrograms []string
	toInstallLibs     []string
	xdgPrograms       []string // 支持 XDG Base Directory 的软件
	dotfiles          []string // $HOME/ 下的点文件
)

func init() {
	xdgPrograms = []string{
		"fish",
		"git",
	}
	dotfiles = []string{
		"spacemacs",
	}
}

func bootstrap(os string) {
	infof("Bootstraping %s...\n", os)

	infof("Installing programs...")
	for _, p := range toInstallPrograms {
		infof("\tInstalling %s...", p)
		if err := installProgram(p); err != nil {
			fatalf("\tInstall %s failed, error: %s.", p, err)
		} else {
			infof("\t%s installed.", p)
		}
	}
	infof("Programs installed.\n")

	infof("Installing librarys...")
	for _, l := range toInstallLibs {
		infof("\tInstalling %s...", l)
		if err := installLib(l); err != nil {
			fatalf("\tInstall %s failed, error: %s.", l, err)
		} else {
			infof("\t%s installed.", l)
		}
	}
	infof("Librarys installed.\n")

	infof("chsh -s $(which fish)...")
	if err := chsh(); err != nil {
		fatalf("chsh -s $(which fish) failed, error: %s.\n", err)
	}
	infof("chsh -s $(which fish) done.\n")

	infof("Configuring packages support XDG base directory...")
	for _, p := range xdgPrograms {
		infof("\tConfiguring %s...", p)
		if err := config(p); err != nil {
			fatalf("Configure %s failed, error: %s.\n", p, err)
		}
		infof("\t%s configured.", p)
	}
	infof("Packages support XDG base directory configured.\n")

	infof("Linking dotfiles...")
	for _, dotfile := range dotfiles {
		infof("\tLinking %s...", dotfile)
		if err := symlink("/dotfiles/"+dotfile, "/."+dotfile); err != nil {
			fatalf("Link %s failed, error: %s.", dotfile, err)
		}
		infof("\t%s linked.", dotfile)
	}
	infof("Dotfiles linked.\n")

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

func config(program string) error {
	if err := symlink("/config/"+program, "/.config/"+program); err != nil {
		return err
	}

	return nil
}
