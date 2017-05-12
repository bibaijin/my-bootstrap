package main

import (
	"runtime"
)

func main() {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		bootstrap()
	} else {
		errorf("Unsupported OS: %s.", runtime.GOOS)
	}
}
