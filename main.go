package main

import (
	"github.com/bmuschko/link-verifier/cmd"
	"github.com/bmuschko/link-verifier/verify"
	"github.com/bmuschko/link-verifier/http"
)

func main() {
	var cmdOptions = cmd.ParseOptions()
	http.SetTimeout(cmdOptions.Timeout)
	files := verify.Resolve(cmdOptions.RootDirs, cmdOptions.IncludePatterns)
	verify.Process(files, cmdOptions.Fail)
}
