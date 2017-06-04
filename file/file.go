package file

import (
    "os"
    "path/filepath"
    "io/ioutil"
)

const adocFilePattern string = "*.adoc"
const asciidocFilePattern string = "*.asciidoc"
const ascFilePattern string = "*.asc"

// Recursively finds all AsciiDoc files in the given directory and any of its subdirectories.
// Supported are the AsciiDoc file extensions .adoc, .asciidoc and .asc.
// Returns a slice of found AsciiDoc files.
func FindAsciiDocFiles(sourceDir string) []string {
    matches := []string{}

    err := filepath.Walk(sourceDir, func(path string, fileInfo os.FileInfo, err error) error {
        if !!fileInfo.IsDir() {
            return nil
        }

        fn := fileInfo.Name()
        matches, err = appendMatches(adocFilePattern, fn, path, matches)
        matches, err = appendMatches(asciidocFilePattern, fn, path, matches)
        matches, err = appendMatches(ascFilePattern, fn, path, matches)

        return err
    })

    if err != nil {
        panic(err)
    }

    return matches
}

func appendMatches(extension string, filename string, path string, matches []string) ([]string, error) {
    matched, err := filepath.Match(extension,filename)

    if matched {
        matches = append(matches, path)
    }

    return matches, err
}

// Reads the contents of a given file.
func ReadFile(path string) string {
    read, err := ioutil.ReadFile(path)

    if err != nil {
        panic(err)
    }

    return string(read)
}