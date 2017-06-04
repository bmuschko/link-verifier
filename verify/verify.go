package verify

import (
	"fmt"
	"github.com/bmuschko/asciidoc-link-verifier/file"
	"github.com/bmuschko/asciidoc-link-verifier/http"
	"github.com/bmuschko/asciidoc-link-verifier/stat"
	"github.com/bmuschko/asciidoc-link-verifier/text"
	"os"
	"strconv"
)

// Resolve resolves AsciiDoc files for a given directory.
// Returns resolved AsciiDoc files.
func Resolve(sourceDir string) []string {
	_, err := os.Stat(sourceDir)

	if os.IsNotExist(err) {
		fmt.Printf("Provided source directory '%s' does not exist!", sourceDir)
		os.Exit(1)
	}

	fmt.Println("Searching AsciiDoc files in directory:", sourceDir)
	return file.FindAsciiDocFiles(sourceDir)
}

// Process processes AsciiDoc files by verifying each parsed links by emitting a HTTP call.
// Prints out a summary of successful and failed links.
// By default fails the program if at least one link could not be resolved.
func Process(asciiDocFiles []string, fail bool) {
	aggregateSummary := []stat.Summary{}

	for _, asciiDocFile := range asciiDocFiles {
		fmt.Println()
		fmt.Println("-> Verifying file:", asciiDocFile)
		content := file.ReadFile(asciiDocFile)
		summary := parseLinks(content)
		aggregateSummary = append(aggregateSummary, summary)
	}

	if len(aggregateSummary) > 0 {
		successCount := stat.SumSuccesses(aggregateSummary)
		failureCount := stat.SumFailures(aggregateSummary)
		stats := fmt.Sprintf("SUCCESSFUL: %s, FAILED: %s", strconv.Itoa(successCount), strconv.Itoa(failureCount))
		fmt.Println()
		fmt.Println(calculateSeparator(stats))
		fmt.Println(stats)

		if failureCount > 0 && !fail {
			os.Exit(1)
		}
	}
}

func parseLinks(content string) stat.Summary {
	links := text.ParseLinks(content)
	summary := stat.Summary{Successful: 0, Failed: 0}

	if len(links) == 0 {
		fmt.Println("No links found.")
	}

	for _, link := range links {
		validateLink(link, &summary)
	}

	return summary
}

func validateLink(link string, summary *stat.Summary) {
	response := http.Get(link)

	if response.Success {
		summary.Successful++
		fmt.Println("[OK] " + link)
	} else {
		summary.Failed++
		fmt.Println("[FAILED] " + link + " (" + response.Status + ")")
	}
}

func calculateSeparator(stats string) string {
	var separator = ""

	for i := 0; i < len(stats); i++ {
		separator += "-"
	}

	return separator
}
