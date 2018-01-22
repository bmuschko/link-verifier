package main

import (
	"github.com/bmuschko/link-verifier/cmd"
	"github.com/bmuschko/link-verifier/http"
	"github.com/bmuschko/link-verifier/verify"
)

func main() {
	var cmdOptions = cmd.ParseOptions()
	http.SetTimeout(cmdOptions.Timeout)
	files := verify.Resolve(cmdOptions.RootDirs, cmdOptions.IncludePatterns)
	verify.Process(files, cmdOptions.Fail)
}
