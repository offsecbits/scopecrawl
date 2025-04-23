package main

import (
	"flag"
	"fmt"

	"github.com/offsecbits/scopecrawl/utils/dedupe"
	"github.com/offsecbits/scopecrawl/utils/inputvalidator"
	"github.com/offsecbits/scopecrawl/utils/linkextractor"
	"github.com/offsecbits/scopecrawl/utils/outputhandler"
	"github.com/offsecbits/scopecrawl/utils/stderr"

)

const version = "v1.0.0"

func main() {
	// Flags
	inputFile := flag.String("f", "", "Path to the input file containing FQDNs or URLs")
	singleURL := flag.String("u", "", "Single URL to validate")
	showVersion := flag.Bool("v", false, "Show version information and exit")
	outputFormat := flag.String("o", "txt", "Output format: txt or json")
	flag.Usage = func() {
		stderr.Usage(version)
	}
	flag.Parse()

	// Version
	if *showVersion {
		stderr.PrintVersion(version)
		return
	}


	var goodURLs, badURLs []string
	var err error

	// Case 1: Process single URL
	if *singleURL != "" {
		normalizedURL, err := inputvalidator.ValidateSingleURL(*singleURL)
		if err != nil {
			stderr.PrintHintedError(err)
			return
		}

		goodURLs = []string{normalizedURL}
		badURLs = []string{} // No need to handle prompt for single URL

	} else if *inputFile != "" {
		// Case 2: Process URL file using the inputvalidator
		goodURLs, badURLs, err = inputvalidator.ValidateInputFile(*inputFile)
		if err != nil {
			stderr.PrintFileError(err)
			return
		}
	} else {
		stderr.PromptMissingInput(version)
		return
	}

	// Handle bad URLs
	if len(badURLs) > 0 && !inputvalidator.HandleBadURLs(badURLs) {
		fmt.Println("Exiting...")
		return
	}

	// Deduplicate URLs
	uniqueURLs := dedupe.RemoveDuplicates(goodURLs)

	// Output
	fmt.Println("\nEnumerating valid target(s):")
	for _, url := range uniqueURLs {
		fmt.Println(" -", url)

		// Fetch the HTML content (passive scan by default)
		htmlContent, err := linkextractor.FetchHTML(url)
		if err != nil {
			stderr.PrintFetchError(url, err)
			continue
		}

		// Extract links from the HTML content
		links, err := linkextractor.ExtractLinks(htmlContent, url)
		if err != nil {
			stderr.PrintExtractError(url, err)
			continue
		}

		// Deduplicate extracted links
		uniqueLinks := dedupe.RemoveDuplicates(links)

		// Output extracted links
		fmt.Println("\nExtracted links from", url)
		for _, link := range uniqueLinks {
			fmt.Println(" -", link)
		}
		fmt.Printf("Total unique links extracted from %s: %d\n", url, len(uniqueLinks))
        	// Save links to file
        	err = outputhandler.SaveLinks(url, uniqueLinks, *outputFormat)
        	if err != nil {
            	stderr.PrintOutputError(url, err)
        	}
	}
}
