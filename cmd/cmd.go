package cmd

import (
	"flag"
	"strings"
)

// ParseOptions parses the command line options.
// Returns the parsed options.
func ParseOptions() Options {
	var rootDirs = flag.String("dirs", ".", "The root source directories used to search for files")
	var includePatterns = flag.String("include", "*.md,*.adoc", "The file inclusion patterns use to search for files")
	var fail = flag.Bool("fail", true, "Fails the program if at least one issue was found")
	var timeout = flag.Int("timeout", 5, "The timeout in seconds used when calling the URL")
	flag.Parse()
	return Options{RootDirs: strings.Split(*rootDirs, ","), IncludePatterns: strings.Split(*includePatterns, ","), Fail: *fail, Timeout: *timeout}
}

// Options represent command line options exposed by this program.
type Options struct {
	RootDirs        []string
	IncludePatterns []string
	Fail            bool
	Timeout         int
}
