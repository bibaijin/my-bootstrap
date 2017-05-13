package main

import (
	"net/http"
	"os"
	"os/exec"
)

func init() {
	toInstallSoftwares = []string{
		"fish",
		"git",
		"vim",
		"emacs",
	}

	infof("Installing brew...")
	if err := installBrew(); err != nil {
		fatalf("Install brew failed, error: %s.\n", err)
	} else {
		infof("Install brew done.\n")
	}
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

	cmd := exec.Command("brew", "install", software)
	if err := runCmd(cmd); err != nil {
		return err
	}

	return nil
}
