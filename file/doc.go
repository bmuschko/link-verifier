package file

import "strings"

var AsciiDoc = Document{"AsciiDoc", []string{"*.adoc", "*.asciidoc", "*.asc"}}
var Markdown = Document{"Markdown", []string{"*.md", "*.markdown", "*.mdown"}}
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
