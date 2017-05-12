package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
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

func _printf(level byte, format string, a ...interface{}) {
	var buf bytes.Buffer
	buf.WriteString("\033[0;3")
	buf.WriteByte(level)
	buf.WriteByte('m')
	buf.WriteString(format)
	buf.WriteString("\033[0m\n")
	fmt.Printf(buf.String(), a...)
}

func runCmd(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
