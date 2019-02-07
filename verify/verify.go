package verify

import (
	"fmt"
	"github.com/bmuschko/link-verifier/file"
	"github.com/bmuschko/link-verifier/http"
	"github.com/bmuschko/link-verifier/stat"
	"github.com/bmuschko/link-verifier/text"
	"os"
	"strconv"
	"strings"
)

// Resolve resolves text-based files for a given directories.
// Returns resolved text-based files.
func Resolve(rootDirs []string, includePatterns []string) []string {
	textBasedFiles := []string{}

	for _, rootDir := range rootDirs {
		_, err := os.Stat(rootDir)

		if os.IsNotExist(err) {
			fmt.Println(fmt.Errorf("Provided root directory '%s' does not exist!", rootDir))
			os.Exit(1)
		}

		foundFiles := file.FindTextBasedFiles(rootDir, includePatterns)
		textBasedFiles = append(textBasedFiles, foundFiles...)
	}

	fmt.Println("Searching text-based files in directories:", strings.Join(rootDirs, ", "))
	return textBasedFiles
}

// Process processes text-based files by verifying each parsed links by emitting a HTTP call.
// Prints out a summary of successful and failed links.
// By default fails the program if at least one link could not be resolved.
func Process(files []string, fail bool) {
	aggregateSummary := []stat.Summary{}

	for _, textBasedFile := range files {
		fmt.Println()
		fmt.Println("-> Verifying file:", textBasedFile)
		content := file.ReadFile(textBasedFile)
		summary := parseLinks(content)
		aggregateSummary = append(aggregateSummary, summary)
	}

	if len(aggregateSummary) > 0 {
		successCount := stat.SumSuccesses(aggregateSummary)
		failureCount := stat.SumFailures(aggregateSummary)
		errorCount := stat.SumErrors(aggregateSummary)
		stats := fmt.Sprintf("SUCCESSFUL: %s, FAILED: %s, ERRORED: %s", strconv.Itoa(successCount), strconv.Itoa(failureCount), strconv.Itoa(errorCount))
		fmt.Println()
		fmt.Println(calculateSeparator(stats))
		fmt.Println(stats)

		if (failureCount > 0 || errorCount > 0) && fail {
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

	ch := make(chan string)

	for _, link := range links {
		go validateLink(link, &summary, ch)
	}

	for range links {
		fmt.Println(<-ch)
	}

	return summary
}

func validateLink(link string, summary *stat.Summary, ch chan<- string) {
	// Try HEAD request first
	response := http.Head(link)

	// Fall back to GET request
	if !response.Success {
		response = http.Get(link)
	}

	if response.Error != nil {
		summary.Errored++
		ch <- fmt.Sprintf("[ERROR] %s (%s)", link, response.Error.Error())
	} else if response.Success {
		summary.Successful++
		ch <- fmt.Sprintf("[OK] %s", link)
	} else {
		summary.Failed++
		ch <- fmt.Sprintf("[FAILED] %s (HTTP %d)", link, response.StatusCode)
	}
}

func calculateSeparator(stats string) string {
	var separator = ""

	for i := 0; i < len(stats); i++ {
		separator += "-"
	}

	return separator
}
