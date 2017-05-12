package main

import (
	"os"
	"os/exec"
	"strings"
)

func chsh() error {
	fish, err := which("fish")
	if err != nil {
		return err
	}

	if strings.Contains(os.Getenv("SHELL"), "fish") {
		infof("fish is already the login shell, skipped.")
		return nil
	}

	cmd := exec.Command("chsh", "-s", fish)
	if err := runCmd(cmd); err != nil {
		return err
	}

	return nil
}

func configFish() error {
	homeDir, err := getHomeDir()
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Symlink(cwd+"/config/fish", homeDir+"/.config/fish")
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}
