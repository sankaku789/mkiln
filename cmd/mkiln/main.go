package main

import (
	"os"

	mkiln "github.com/sankaku789/mkiln/internal"
)

var version = "dev"

func main() {
	os.Exit(mkiln.Run(os.Args[1:], version, os.Stdout, os.Stderr))
}
