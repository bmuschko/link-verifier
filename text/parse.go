// Package text uses the library "github.com/mvdan/xurls" to parse URLs from a string.
// The library is distributed with the license BSD 3-clause.
// Copyright (c) 2015, Daniel Martí. All rights reserved.
package text

import (
	"github.com/mvdan/xurls"
	"regexp"
)

// ParseLinks parses a given text and extracts all links. None of the links is further modified except for the rules listed below.
// Sanitizes founds like based on the following logic:
//   - Removes link description e.g. [...] if extraction logic couldn't remove it.
//   - Remove fragments in URLs e.g. #sec:news - they are a browser-only concept.
// Does not included links based on the following logic:
//   - URLs that contains String interpolation with ${...}.
// Returns a slice of links.
func ParseLinks(content string) []string {
	links := xurls.Strict.FindAllString(content, -1)

	for i, link := range links {
		sanatizedLink := sanitizeLink(link)
		links[i] = sanatizedLink

		if skipLink(links[i]) {
			links[i] = links[len(links)-1]
		}
	}

	return links
}

func sanitizeLink(link string) string {
	// remove link description if extraction logic couldn't remove it
	unremovedLinkDescription := regexp.MustCompile("\\[.*]")
	sanitizeLink := unremovedLinkDescription.ReplaceAllString(link, "")
	// remove fragments - they are a browser-only concept
	fragment := regexp.MustCompile("#.*")
	sanitizeLink = fragment.ReplaceAllString(sanitizeLink, "")
	return sanitizeLink
}

func skipLink(link string) bool {
	re := regexp.MustCompile("\\$\\{.*}")
	return re.MatchString(link)
}
