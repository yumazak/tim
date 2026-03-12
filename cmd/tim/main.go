package main

import (
	"os"

	"github.com/yumazak/tim/internal"
)

func main() {
	os.Exit(internal.Run(os.Args[1:]))
}
