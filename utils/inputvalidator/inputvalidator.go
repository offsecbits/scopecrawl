package inputvalidator

import (
	"bufio"
	"net/url"
	"fmt"
	"os"
	"regexp"
	"strings"
	"github.com/offsecbits/scopecrawl/utils/urlnormalize"
)

// ValidateInputFile validates and normalizes the URLs from the given file.
func ValidateInputFile(filePath string) ([]string, []string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var goodURLs, badURLs []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()


		// Normalize the URL if valid
		normalizedURL, err := urlnormalize.Normalize(line)
		if err != nil || !isValidURL(normalizedURL) {
			badURLs = append(badURLs, line)
		} else {
			goodURLs = append(goodURLs, normalizedURL)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return goodURLs, badURLs, nil
}

// ValidateSingleURL checks a single URL for validity and normalization.
func ValidateSingleURL(raw string) (string, error) {

		normalized, err := urlnormalize.Normalize(raw)
		if err != nil || !isValidURL(normalized) {
    		return "", fmt.Errorf("invalid URL: %s", raw)
		}
		return normalized, nil
}

// HandleBadURLs will prompt the user with bad URLs and ask if they want to continue.
func HandleBadURLs(badURLs []string) bool {
	if len(badURLs) > 0 {
		fmt.Println("\nWarning: The following invalid URLs were skipped during processing:")
		for _, bad := range badURLs {
			fmt.Println("  -", bad)
		}

		fmt.Print("\nDo you want to continue with the valid URLs? (y/n): ")
		var answer string
		fmt.Scanln(&answer)

		return strings.ToLower(answer) == "y"
	}
	return true
}

// isValidURL checks if a URL is valid based on basic rules like structure and host.
func isValidURL(raw string) bool {
	// Check for empty or URL containing spaces
	raw = strings.TrimSpace(raw)
	if raw == "" || strings.Contains(raw, " ") {
		return false
	}

	parsed, err := url.Parse(raw)
	if err != nil || parsed.Host == "" {
		return false
	}

	// Validate the hostname only
	host := parsed.Hostname()
	re := regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)
	matches := re.FindStringSubmatch(host)
	if len(matches) < 1 {
		return false
	}

	return true
}
