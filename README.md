# scopecrawl

**scopecrawl** is a powerful, easy-to-use CLI tool for web crawling, URL validation, and link extraction. Written in Go, **scopecrawl** helps you validate URLs, remove duplicates, and extract valuable links from web pages. This tool is ideal for security professionals, web developers, and anyone in need of efficient URL analysis and crawling.

## Features

- **Single URL validation**: Quickly validate and normalize any single URL.
- **Batch URL processing**: Read URLs from a file, validate, and process them.
- **Deduplication**: Automatically remove duplicate URLs for cleaner analysis.
- **Link extraction**: Extract all internal and external links from web pages.
- **Output support**: Export results in plain text or JSON format.
- **Customizable options**: Easily control the output format, input source, and more with CLI flags.
  
## Installation

### Prerequisites

Ensure you have Go installed on your system. If not, [install Go](https://golang.org/doc/install).

### Install `scopecrawl`

To install `scopecrawl`, use the following command:

```bash
go install github.com/yourusername/scopecrawl@latest
