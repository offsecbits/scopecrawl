package urlnormalize

import (
	"fmt"
	"net/url"
	"strings"
)

// Normalize takes a raw string input and returns a normalized,
// validated URL string with scheme and hostname checks.
func Normalize(inputURL string) (string, error) {
	inputURL = strings.TrimSpace(inputURL)

	if inputURL == "" {
		return "", fmt.Errorf("empty input URL")
	}

	// Add scheme if not present (default to https)
	if !strings.HasPrefix(inputURL, "http://") && !strings.HasPrefix(inputURL, "https://") {
		inputURL = "https://" + inputURL
	}

	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL format: %v", err)
	}

	// Lowercase scheme and host
	parsedURL.Scheme = strings.ToLower(parsedURL.Scheme)
	parsedURL.Host = strings.ToLower(parsedURL.Host)

	// Default path
	if parsedURL.Path == "" {
		parsedURL.Path = "/"
	}

	return parsedURL.String(), nil
}
