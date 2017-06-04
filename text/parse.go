package text

import (
    "github.com/mvdan/xurls"
    "regexp"
)

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
