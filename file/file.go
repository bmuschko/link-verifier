package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const adocFilePattern string = "*.adoc"
const asciidocFilePattern string = "*.asciidoc"
const ascFilePattern string = "*.asc"
const mdFilePattern string = "*.md"
const markdownFilePattern string = "*.markdown"
const mdownFilePattern string = "*.mdown"
const txtFilePattern string = "*.txt"

// FindTextBasedFiles recursively finds all text-based files in the given directory and any of its subdirectories.
// Supported are the file extensions:
// - AsciiDoc: .adoc, .asciidoc and .asc
// - Markdown: .md, .markdown and .mdown
// - Plain text: .txt
// Returns a slice of found text-based files.
func FindTextBasedFiles(sourceDir string) []string {
	matches := []string{}

	err := filepath.Walk(sourceDir, func(path string, fileInfo os.FileInfo, err error) error {
		if !!fileInfo.IsDir() {
			return nil
		}

		fn := fileInfo.Name()
		matches, err = appendMatches(adocFilePattern, fn, path, matches)
		matches, err = appendMatches(asciidocFilePattern, fn, path, matches)
		matches, err = appendMatches(ascFilePattern, fn, path, matches)
		matches, err = appendMatches(mdFilePattern, fn, path, matches)
		matches, err = appendMatches(markdownFilePattern, fn, path, matches)
		matches, err = appendMatches(mdownFilePattern, fn, path, matches)
		matches, err = appendMatches(txtFilePattern, fn, path, matches)

		return err
	})

	if err != nil {
		panic(err)
	}

	return matches
}

func appendMatches(extension string, filename string, path string, matches []string) ([]string, error) {
	matched, err := filepath.Match(extension, filename)

	if matched {
		matches = append(matches, path)
	}

	return matches, err
}

// ReadFile reads the contents of a given file.
func ReadFile(path string) string {
	read, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(read)
}
