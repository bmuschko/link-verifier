package main

import (
	"github.com/bmuschko/link-verifier/cmd"
	"github.com/bmuschko/link-verifier/verify"
)

func main() {
	var cmdOptions = cmd.ParseOptions()
	files := verify.Resolve(cmdOptions.SourceDir)
	verify.Process(files, cmdOptions.Fail)
}
