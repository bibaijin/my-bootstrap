package main

import (
	"runtime"
)

func main() {
	if runtime.GOOS == "linux" {
		infof("Bootstraping linux...\n")
		bootstrapLinux()
		infof("Bootstrap linux done.")
	} else if runtime.GOOS == "darwin" {
		infof("Bootstraping darwin...")
		bootstrapDarwin()
		infof("Bootstrap darwin done.")
	} else {
		errorf("Unsupported OS: %s.", runtime.GOOS)
	}
}
