package main

import (
	"flag"
	"fmt"

	"github.com/offsecbits/scopecrawl/utils/aesthetics"
	"github.com/offsecbits/scopecrawl/utils/dedupe"
	"github.com/offsecbits/scopecrawl/utils/inputvalidator"
	"github.com/offsecbits/scopecrawl/utils/linkextractor"
	"github.com/offsecbits/scopecrawl/utils/outputhandler"
	"github.com/offsecbits/scopecrawl/utils/stderr"
)

const version = "v1.0.1"

func main() {

	aesthetics.PrintBanner()
	// Flags
	inputFile := flag.String("f", "", "Path to the input file containing FQDNs or URLs")
	singleURL := flag.String("u", "", "Single URL to validate")
	showVersion := flag.Bool("v", false, "Show version information and exit")
	outputFormat := flag.String("o", "", "Output format: txt or json. Leave blank to skip saving output.")
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

	// Determine if output should be saved
	writeOutput := false
	format := "txt"

	if *outputFormat != "" {
		writeOutput = true
		if *outputFormat == "json" {
			format = "json"
		}
	}

	// Output
	aesthetics.PrintInfo("\nEnumerating valid target(s):")
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

		// Output Logic

		// Save links to file only if -o flag was used
		if writeOutput {
			err = outputhandler.SaveLinks(url, uniqueLinks, format)
			if err != nil {
				stderr.PrintOutputError(url, err)
			}
		} else {
			// Print to terminal if -o flag is missing
			fmt.Println("\nExtracted links from", url)
			for _, link := range uniqueLinks {
				fmt.Println(" -", link)
			}
		}
		fmt.Printf("Total unique links extracted from %s: %d\n", url, len(uniqueLinks))
	
	}
}
