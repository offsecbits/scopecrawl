package activescanner

import (
//    "fmt"
    "net/url"

    "sync"

    "github.com/offsecbits/scopecrawl/utils/dedupe"
    "github.com/offsecbits/scopecrawl/utils/filters"   // Import the filter package
    "github.com/offsecbits/scopecrawl/utils/linkextractor"
    "github.com/offsecbits/scopecrawl/utils/ratelimiter"
)

// Crawl performs active crawling up to the specified depth
func Crawl(startURL string, depth int, limiter *ratelimiter.Limiter) []string {
    var (
        allLinks []string
        visited  = make(map[string]bool)
        mu       sync.Mutex
        wg       sync.WaitGroup
    )

    currentLevel := []string{startURL}

    for d := 0; d <= depth; d++ {
        var nextLevel []string

        for _, currentURL := range currentLevel {
            mu.Lock()
            if visited[currentURL] {
                mu.Unlock()
                continue
            }
            visited[currentURL] = true
            mu.Unlock()

            // Check if the URL is an HTML page before fetching
            if !filters.IsHTML(currentURL) {
                continue // Skip non-HTML URLs
            }

            wg.Add(1)
            urlToFetch := currentURL
            limiter.Submit(urlToFetch, func(urlToFetch string) {
                defer wg.Done()

//		fmt.Println("Fetching:", urlToFetch)

                html, err := linkextractor.FetchHTML(urlToFetch)
                if err != nil {
                    // You can remove this print if you're handling errors elsewhere
                    return
                }

                // Extract the links from the HTML
                links, err := linkextractor.ExtractLinks(html, urlToFetch)
                if err != nil {
                    // Handle this error elsewhere, avoid printing here
                    return
                }

                mu.Lock()
                allLinks = append(allLinks, links...)
                mu.Unlock()
            })
        }

        wg.Wait()

        // Prepare for next depth level, removing duplicates
        allLinks = dedupe.RemoveDuplicates(allLinks)

        for _, link := range allLinks {
            mu.Lock()
            if !visited[link] {
                parsed, err := url.Parse(link)
                if err == nil && parsed.Scheme != "" && parsed.Host != "" {
                    nextLevel = append(nextLevel, link)
                }
            }
            mu.Unlock()
        }

        currentLevel = dedupe.RemoveDuplicates(nextLevel)
    }

    // Final dedupe before returning
    finalLinks := dedupe.RemoveDuplicates(allLinks)

    // Apply filter: Remove non-HTML resources
//    finalLinks = filters.FilterNonHTML(finalLinks)

    // Sort the final links
    finalLinks = filters.SortLinks(finalLinks)

    return finalLinks
}
