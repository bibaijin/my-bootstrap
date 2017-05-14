package main

import (
	"net/http"
	"os"
	"os/exec"
)

func init() {
	toInstallPrograms = []string{
		"fish",
		"vim",
		"emacs",
	}

	toInstallLibs = []string{}

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

func installProgram(program string) error {
	_, err := which(program)
	if err == nil {
		warnf("\t%s is installed, skipped.", program)
		return nil
	}

	cmd := exec.Command("brew", "install", program)
	if err := runCmd(cmd); err != nil {
		return err
	}

	return nil
}

func installLib(lib string) error {
	c := exec.Command("brew", "install", lib)
	if err := runCmd(c); err != nil {
		return err
	}

	return nil
}

func getHomeDir() (string, error) {
	if os.Geteuid() == 0 {
		return "/Users/" + os.Getenv("SUDO_USER"), nil
	}

	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return u.HomeDir, nil
}
