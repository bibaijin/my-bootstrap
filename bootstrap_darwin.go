package main

import (
	"net/http"
	"os"
	"os/exec"
)

var (
	softwares []string
)

func init() {
	softwares = []string{
		"brew",
		"fish",
		"git",
		"vim",
		"emacs",
	}
}

func bootstrap() {
	infof("Bootstraping darwin...\n")

	infof("Installing brew...")
	if err := installBrew(); err != nil {
		fatalf("Install brew failed, error: %s.\n", err)
	} else {
		infof("Install brew done.\n")
	}

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

	infof("Bootstrap darwin done.")
}

func installBrew() error {
	_, err := which("brew")
	if err == nil {
		warnf("brew is installed, skipped.")
		return nil
	}

	resp, err := http.Get("https://raw.githubusercontent.com/Homebrew/install/master/install")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	cmd := exec.Command("/usr/bin/ruby", "-e")
	cmd.Stdin = resp.Body
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return err
	}

	return nil
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
