// Package text uses the library "github.com/mvdan/xurls" to parse URLs from a string.
// The library is distributed with the license BSD 3-clause.
// Copyright (c) 2015, Daniel Martí. All rights reserved.
package text

import (
	"github.com/mvdan/xurls"
	"regexp"
	"strings"
)

// ParseLinks parses a given text and extracts all links. None of the links is further modified except for the rules listed below.
// Sanitizes founds like based on the following logic:
//   - Removes link description e.g. [...] if extraction logic couldn't remove it.
//   - Remove fragments in URLs e.g. #sec:news - they are a browser-only concept.
// Does not included links based on the following logic:
//   - URLs that contains String interpolation with ${...}.
// Returns a slice of links.
func ParseLinks(content string) []string {
	uniqueLinks := make(map[string]bool)
	links := xurls.Strict.FindAllString(content, -1)

	for _, link := range links {
		sanatizedLink := sanitizeLink(link)

		if !uniqueLinks[sanatizedLink] && !skipLink(sanatizedLink) {
			uniqueLinks[sanatizedLink] = true
		}
	}

	return keysInMap(uniqueLinks)
}

func keysInMap(data map[string]bool) []string {
	keys := make([]string, 0, len(data))

	for k := range data {
		keys = append(keys, k)
	}

	return keys
}

func sanitizeLink(link string) string {
	// remove link description if extraction logic couldn't remove it
	unremovedLinkDescription := regexp.MustCompile(`\[.*]`)
	sanitizeLink := unremovedLinkDescription.ReplaceAllString(link, "")
	// remove fragments - they are a browser-only concept
	fragment := regexp.MustCompile("#.*")
	sanitizeLink = fragment.ReplaceAllString(sanitizeLink, "")
	return sanitizeLink
}

func skipLink(link string) bool {
	// localhost URLs
	localhostUrl := strings.Contains(link, "localhost")
	// mailto links
	mailtoLink := strings.Contains(link, "mailto:")
	// placeholder in link e.g. http://${host}/path
	re := regexp.MustCompile(`\$\{.*}`)
	return re.MatchString(link) || localhostUrl || mailtoLink
}
