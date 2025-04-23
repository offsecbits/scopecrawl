package linkextractor

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"github.com/offsecbits/scopecrawl/utils/dedupe"

	"golang.org/x/net/html"
)

// FetchHTML retrieves the raw HTML content from a URL
func FetchHTML(targetURL string) (string, error) {
	// Send GET request
	resp, err := http.Get(targetURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP 200 OK
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Read the entire HTML body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(bodyBytes), nil
}

// ExtractLinks parses HTML content and extracts only same-origin (internal) links
func ExtractLinks(htmlContent, baseURL string) ([]string, error) {
	var links []string

	// Parse the base URL
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %v", err)
	}

	// Parse the HTML content
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	// Recursive function to walk through the DOM tree
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "a", "link":
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						link := strings.TrimSpace(attr.Val)
						resolved := resolveURL(base, link)
						if isSameOrigin(base, resolved) {
							links = append(links, resolved.String())
						}
					}
				}
			case "script", "img":
				for _, attr := range n.Attr {
					if attr.Key == "src" {
						link := strings.TrimSpace(attr.Val)
						resolved := resolveURL(base, link)
						if isSameOrigin(base, resolved) {
							links = append(links, resolved.String())
						}
					}
				}
			}
		}

		// Recurse
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)

	// Deduplicate final list of internal links
	cleanLinks := dedupe.RemoveDuplicates(links)

	return cleanLinks, nil
}

// resolveURL makes relative URLs absolute based on base URL
func resolveURL(base *url.URL, href string) *url.URL {
	parsed, err := url.Parse(href)
	if err != nil {
		return base
	}
	return base.ResolveReference(parsed)
}

// isSameOrigin checks if two URLs have the same scheme and host
func isSameOrigin(base, target *url.URL) bool {
	return base.Scheme == target.Scheme && base.Host == target.Host
}
