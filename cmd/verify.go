package cmd

import (
	"github.com/bmuschko/link-verifier/http"
	"github.com/bmuschko/link-verifier/verify"
	"github.com/spf13/cobra"
)

func doVerifyCmd(cmd *cobra.Command, args []string) {
	http.SetTimeout(parsedOptions.Timeout)
	files := verify.Resolve(parsedOptions.RootDirs, parsedOptions.IncludePatterns)
	verify.Process(files, parsedOptions.IgnoreStatusCodes, parsedOptions.Fail)
}
