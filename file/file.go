package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const AdocFilePattern string = "*.adoc"
const AsciidocFilePattern string = "*.asciidoc"
const AscFilePattern string = "*.asc"
const MdFilePattern string = "*.md"
const MarkdownFilePattern string = "*.markdown"
const MdownFilePattern string = "*.mdown"
const TxtFilePattern string = "*.txt"

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
		matches, err = appendMatches(AdocFilePattern, fn, path, matches)
		matches, err = appendMatches(AsciidocFilePattern, fn, path, matches)
		matches, err = appendMatches(AscFilePattern, fn, path, matches)
		matches, err = appendMatches(MdFilePattern, fn, path, matches)
		matches, err = appendMatches(MarkdownFilePattern, fn, path, matches)
		matches, err = appendMatches(MdownFilePattern, fn, path, matches)
		matches, err = appendMatches(TxtFilePattern, fn, path, matches)

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
