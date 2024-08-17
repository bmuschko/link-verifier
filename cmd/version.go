package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number and exit",
	Run:   printVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func SetVersion(v string) {
	version = v
}

func printVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("link verifier %s\n", version)
}
