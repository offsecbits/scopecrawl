package filters

import (
    "net/url"
    "sort"
    "strings"
)

// FilterNonHTML filters out non-HTML resources like images, .js, .css, etc.
func FilterNonHTML(links []string) []string {
    var filteredLinks []string

    for _, link := range links {
        if IsHTML(link) {
            filteredLinks = append(filteredLinks, link)
        }
    }
    return filteredLinks
}

// isHTML checks if the URL points to an HTML resource.
func IsHTML(link string) bool {
    // Filter out common file types (images, JavaScript, CSS, etc.)
    nonHTMLTypes := []string{".jpg", ".jpeg", ".png", ".webp", ".gif", ".bmp", ".js", ".css", ".svg", ".ico", ".woff", ".ttf", ".pdf", ".json", ".xml", ".php"}
    for _, ext := range nonHTMLTypes {
        if strings.HasSuffix(strings.ToLower(link), ext) {
            return false
        }
    }

    // If the URL doesn't end with non-HTML file types, consider it an HTML resource
    return true
}

// FilterByDomain filters links based on a specific domain.
func FilterByDomain(links []string, domain string) []string {
    var filteredLinks []string
    for _, link := range links {
        parsedURL, err := url.Parse(link)
        if err != nil {
            continue
        }

        // Only add the links with the matching domain
        if parsedURL.Host == domain {
            filteredLinks = append(filteredLinks, link)
        }
    }
    return filteredLinks
}

// SortLinks sorts the links in lexicographical (alphabetical) order.
func SortLinks(links []string) []string {
    sort.Strings(links)
    return links
}

// FilterByKeyword filters links containing a specific keyword.
func FilterByKeyword(links []string, keyword string) []string {
    var filteredLinks []string
    for _, link := range links {
        if strings.Contains(link, keyword) {
            filteredLinks = append(filteredLinks, link)
        }
    }
    return filteredLinks
}

// FilterByPath filters links based on the path (e.g., "/about" or "/contact").
func FilterByPath(links []string, path string) []string {
    var filteredLinks []string
    for _, link := range links {
        parsedURL, err := url.Parse(link)
        if err != nil {
            continue
        }

        if strings.Contains(parsedURL.Path, path) {
            filteredLinks = append(filteredLinks, link)
        }
    }
    return filteredLinks
}
