package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func infof(format string, a ...interface{}) {
	_printf('2', format, a...)
}

func warnf(format string, a ...interface{}) {
	_printf('3', format, a...)
}

func errorf(format string, a ...interface{}) {
	_printf('1', format, a...)
}

func fatalf(format string, a ...interface{}) {
	errorf(format, a...)
	os.Exit(1)
}

func _printf(level byte, format string, a ...interface{}) {
	var buf bytes.Buffer
	buf.WriteString("\033[0;3")
	buf.WriteByte(level)
	buf.WriteByte('m')
	buf.WriteString(format)
	buf.WriteString("\033[0m\n")
	fmt.Printf(buf.String(), a...)
}

func which(executable string) (string, error) {
	cmd := exec.Command("which", executable)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func runCmd(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func symlink(oldname, newname string) error {
	homeDir, err := getHomeDir()
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Symlink(cwd+oldname, homeDir+newname)
	if err != nil {
		if os.IsExist(err) {
			warnf("\t~%s already exists, skipped.", newname)
			return nil
		}

		return err
	}

	return nil
}
