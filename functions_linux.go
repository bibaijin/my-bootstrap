package main

import (
	"errors"
	"os"
	"os/exec"
	"os/user"
)

func init() {
	toInstallPrograms = []string{
		"fish",
		"vim",
		"emacs",
	}
	toInstallLibs = []string{
		"aspell-en",
	}
}

func installProgram(program string) error {
	_, err := which(program)
	if err == nil {
		warnf("\t%s is installed, skipped.", program)
		return nil
	}

	c := exec.Command("pacman", "-S", "--noconfirm", program)
	if err := runCmd(c); err != nil {
		return err
	}

	return nil
}

func installLib(lib string) error {
	if os.Geteuid() != 0 {
		return errors.New("sudo priviledge is required")
	}

	c := exec.Command("pacman", "-S", "--noconfirm", lib)
	if err := runCmd(c); err != nil {
		return err
	}

	return nil
}

func getHomeDir() (string, error) {
	if os.Geteuid() == 0 {
		return "/home/" + os.Getenv("SUDO_USER"), nil
	}

	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return u.HomeDir, nil
}
