package main

import (
	"fmt"
	"runtime"
)

var version = "dev"

func main() {
	fmt.Printf("Hello World from NPM! (%s, %s, %s)\n", version, runtime.GOOS, runtime.GOARCH)
}
