package file

import "strings"

// AsciiDoc represents a document using the AsciiDoc markup language.
var AsciiDoc = Document{"AsciiDoc", []string{"*.adoc", "*.asciidoc", "*.asc"}}
// Markdown represents a document using the Markdown markup language.
var Markdown = Document{"Markdown", []string{"*.md", "*.markdown", "*.mdown"}}
// AsciiDoc represents a plain-text document.
var PlainText = Document{"Plain Text", []string{"*.txt"}}

type Document struct {
	name string
	ext  []string
}

// JoinedExt joins all registered file extensions of a document.
// Returns the joined extensions of a document separated by comma.
func (d Document) JoinedExt() string {
	return strings.Join(d.ext, ",")
}
