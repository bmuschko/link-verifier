package cmd

import (
	"github.com/bmuschko/link-verifier/verify"
	"github.com/spf13/cobra"
)

func doVerifyCmd(cmd *cobra.Command, args []string) {
	files := verify.Resolve(parsedOptions.RootDirs, parsedOptions.IncludePatterns)
	verify.Process(files, parsedOptions.Timeout, parsedOptions.IgnoreStatusCodes, parsedOptions.Fail)
}
