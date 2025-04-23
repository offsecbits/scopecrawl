package dedupe

// RemoveDuplicates takes a slice of URLs and returns a deduplicated list.
func RemoveDuplicates(urls []string) []string {
	unique := make([]string, 0)
	seen := make(map[string]struct{})

	for _, url := range urls {
		if _, found := seen[url]; !found {
			unique = append(unique, url)
			seen[url] = struct{}{}
		}
	}

	return unique
}
