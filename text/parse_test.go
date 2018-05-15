package text

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const expectedLink = "http://www.oracle.com/technetwork/articles/java/index-jsp-135444.html"

func TestParseLinkWithLinkPrefix(t *testing.T) {
	text := fmt.Sprintf("Java provides the link:%s[`javadoc`] tool for generating Javadocs documentation from source code.", expectedLink)

	links := ParseLinks(text)

	assert.Len(t, links, 1)
	assert.Equal(t, expectedLink, links[0])
}

func TestParseLinkWithoutLinkPrefix(t *testing.T) {
	text := fmt.Sprintf("Java provides the %s[`javadoc`] tool for generating Javadocs documentation from source code.", expectedLink)

	links := ParseLinks(text)

	assert.Len(t, links, 1)
	assert.Equal(t, expectedLink, links[0])
}

func TestParseLinkInText(t *testing.T) {
	text := fmt.Sprintf("Java provides the %s tool for generating Javadocs documentation from source code.", expectedLink)

	links := ParseLinks(text)

	assert.Len(t, links, 1)
	assert.Equal(t, expectedLink, links[0])
}

func TestParseLinkWithPlaceholder(t *testing.T) {
	text := fmt.Sprintf("Java provides the %s tool for generating Javadocs documentation from source code.", "http://${host}:8080/")

	links := ParseLinks(text)

	assert.Len(t, links, 0)
}

func TestParseLinkWithLocalhostUrl(t *testing.T) {
	text := fmt.Sprintf("Java provides the %s tool for generating Javadocs documentation from source code.", "http://localhost:8080/")

	links := ParseLinks(text)

	assert.Len(t, links, 0)
}

func TestParseLinkWithMailtoLink(t *testing.T) {
	text := fmt.Sprintf("Send me a mail to %s.", "mailto:benjamin.muschko@gmail.com")

	links := ParseLinks(text)

	assert.Len(t, links, 0)
}

func TestLinkInCode(t *testing.T) {
	text := fmt.Sprintf(`if (true) {
        return "%s";
    } else {
        return null;
    }`, expectedLink)

	links := ParseLinks(text)

	assert.Len(t, links, 1)
	assert.Equal(t, expectedLink, links[0])
}

func TestSanatizeLinkDescription(t *testing.T) {
	text := "http://www.javadoc.io/[javadoc.io]"
	links := ParseLinks(text)

	assert.Len(t, links, 1)
	assert.Equal(t, "http://www.javadoc.io/", links[0])
}

func TestSanatizeFragementInLink(t *testing.T) {
	text := "https://docs.gradle.org/current/userguide/java_plugin.html#sec:javadoc"
	links := ParseLinks(text)

	assert.Len(t, links, 1)
	assert.Equal(t, "https://docs.gradle.org/current/userguide/java_plugin.html", links[0])
}
