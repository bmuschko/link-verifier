package main

import (
	"github.com/bmuschko/asciidoc-link-verifier/cmd"
	"github.com/bmuschko/asciidoc-link-verifier/verify"
)

func main() {
	var cmdOptions = cmd.ParseOptions()
	asciiDocFiles := verify.Resolve(cmdOptions.SourceDir)
	verify.Process(asciiDocFiles, cmdOptions.Fail)
}
