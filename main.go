package main

import (
	"runtime"
)

func main() {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		bootstrap(runtime.GOOS)
	} else {
		errorf("Unsupported OS: %s.", runtime.GOOS)
	}
}
