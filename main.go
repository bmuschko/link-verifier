package main

import (
	"github.com/bmuschko/link-verifier/cmd"
	"github.com/bmuschko/link-verifier/verify"
)

func main() {
	var cmdOptions = cmd.ParseOptions()
	files := verify.Resolve(cmdOptions.RootDirs, cmdOptions.IncludePatterns)
	verify.Process(files, cmdOptions.Fail)
}
