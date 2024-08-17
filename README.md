# Link Verifier [![Build and Publish](https://github.com/bmuschko/link-verifier/actions/workflows/build-publish.yml/badge.svg)](https://github.com/bmuschko/link-verifier/actions/workflows/build-publish.yml)

A tool for verifying links in text-based files. Written in Go, available as executable.

![Logo](https://user-images.githubusercontent.com/440872/27007990-1184b292-4e34-11e7-8417-fc62542250b5.jpg)

## Motivation

Ambitious tech writers use plain-text mark-up formats like [AsciiDoc](http://asciidoc.org/) and
[Markdown](https://daringfireball.net/projects/markdown/) for turning text into properly formatted blog posts, web
pages and documentation. Incorporating URLs to refer to external resources is part of the process. Before publishing
content it's important to verify that linked URLs are valid and can be resolved. Nothing feels more unprofessional than broken
links in a carefully crafted document.

Link verifier to the rescue! Automatically discover all mark-up files in a given directory and verify all links
found in the documents before publishing them.

## Installation

### Prebuilt libraries

[Prebuilt libraries for various platforms](https://github.com/bmuschko/link-verifier/releases) are available on GitHub.
Just download, exact the archive and execute the binary.

### Building from source

Building from source requires that you have [Go](https://golang.org/doc/install) installed on your machine.

```
$ go get -u -v github.com/bmuschko/link-verifier
```

Run the program with `link-verifier` from anywhere in the file system.

## Command line options

| Option                | Description | Default Value |
| --------------------- | ----------- | ------------- |
| `dirs`                | The comma-separated root directories used to recursively search for files. | [.] |
| `includes`            | The comma-separated include patterns used to search for files. | [\*.md,*.adoc] |
| `ignore-status-codes` | The comma-separated HTTP response status codes that will be ignored for validation. For example, some sites block crawlers and return with a specific response code. | [] |
| `fail`                | Fails the program if at least one discovered link cannot be resolved. | true |
| `timeout`             | The timeout in seconds used when calling the URL. | 5 |

**Example**:

```
$ link-verifier --dirs data,content --include *.html,*.yml --fail=false
```

## Usage on CI

Integrating the tool with a build on CI is a breeze and can used as validation step before publishing mark-up files.

## GitHub Actions

The following `.github/workflows/verify-links.yml` demonstrate the use of the prebuilt binary version on Linux 64 bit:

``` yaml
- name: Verify Links
  run: |
    wget https://github.com/bmuschko/link-verifier/releases/download/v0.1/link-verifier-0.1-linux64.tar.gz -O /tmp/link-verifier.tar.gz
    tar -xvf /tmp/link-verifier.tar.gz
    export PATH=$PATH:$PWD/link-verifier/
    ./link-verifier
```

## Contribute!

It's easy to contribute to this project. Install link:https://golang.org/doc/install[Go] >= 1.11. Then run the following commands to get the source code, resolve external dependencies
and build the project.

```
$ git clone https://github.com/bmuschko/link-verifier.git
$ cd link-verifier
$ go build
```