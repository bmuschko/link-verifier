package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// ParsedOptions only contains populated data after Execute() was called.
var parsedOptions = Options{}

// rootCmd represents the root CLI command.
var rootCmd = &cobra.Command{
	Use:   "link-verifier",
	Short: "Links verifier",
	Long:  `A tool for verifying links in text-based files.`,
	Run:   doVerifyCmd,
}

func init() {
	rootCmd.PersistentFlags().StringSliceVar(&parsedOptions.RootDirs, "dirs", []string{"."}, "comma-separated root directories used to recursively search for files")
	rootCmd.PersistentFlags().StringSliceVar(&parsedOptions.IncludePatterns, "include", []string{"*.md", "*.adoc"}, "comma-separated root directories used to recursively search for files")
	rootCmd.PersistentFlags().BoolVar(&parsedOptions.Fail, "fail", true, "fails the program if at least one discovered link cannot be resolved")
	rootCmd.PersistentFlags().IntVar(&parsedOptions.Timeout, "timeout", 5, "timeout in seconds used when calling the URL")
}

// Execute runs the CLI root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Options represent command line options exposed by this program.
type Options struct {
	RootDirs        []string
	IncludePatterns []string
	Fail            bool
	Timeout         int
}
