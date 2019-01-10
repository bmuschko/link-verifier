package main

import (
	"github.com/bmuschko/link-verifier/cmd"
)

var (
	version = "undefined"
)

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
