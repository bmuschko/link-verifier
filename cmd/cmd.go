package cmd

import (
	"flag"
)

// ParseOptions parses the command line options.
// Returns the parsed options.
func ParseOptions() Options {
	var sourceDir = flag.String("sourceDir", "content", "The source directory to search for mark-up files")
	var fail = flag.Bool("fail", true, "Fails the program if at least one issue was found")
	flag.Parse()
	return Options{SourceDir: *sourceDir, Fail: *fail}
}

// Options represent command line options exposed by this program.
type Options struct {
	SourceDir string
	Fail      bool
}
