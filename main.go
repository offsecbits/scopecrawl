package main

import (
        "flag"
        "fmt"


        "github.com/offsecbits/scopecrawl/utils/aesthetics"
        "github.com/offsecbits/scopecrawl/utils/dedupe"
        "github.com/offsecbits/scopecrawl/utils/inputvalidator"
//        "github.com/offsecbits/scopecrawl/utils/linkextractor"
        "github.com/offsecbits/scopecrawl/utils/ratelimiter"
        "github.com/offsecbits/scopecrawl/utils/outputhandler"
        "github.com/offsecbits/scopecrawl/utils/stderr"
        "github.com/offsecbits/scopecrawl/utils/activescanner"

)

const version = "v1.1.0"


func main() {
    aesthetics.PrintBanner()
    // Flags
    inputFile := flag.String("f", "", "Path to the input file containing FQDNs or URLs")
    singleURL := flag.String("u", "", "Single URL to validate")
    showVersion := flag.Bool("v", false, "Show version information and exit")
    outputFormat := flag.String("o", "", "Output format: txt or json. Leave blank to skip saving output.")
    depth := flag.Int("d", 0, "Depth for recursive crawling (0 for passive scanning)")
    rate := flag.Int("r", 2, "Requests per second (rate limiting)")
    threads := flag.Int("t", 2, "Number of concurrent threads for crawling")

    flag.Usage = func() {
        stderr.Usage(version)
    }
    flag.Parse()

    // Call your validator functions
    inputvalidator.ValidatePositiveInt("t", *threads)
    inputvalidator.ValidatePositiveInt("r", *rate)
    inputvalidator.ValidateNonNegativeInt("d", *depth)

    // ✅ Safe to continue
//    fmt.Println(aesthetics.Blue + "Starting with | Threads:",*threads, "| Rate:",*rate, "| Depth:",*depth, "|" + aesthetics.Reset)

    // Version
    if *showVersion {
        stderr.PrintVersion(version)
        return
    }

    var goodURLs, badURLs []string
    var err error

    // Process the input URLs
    if *singleURL != "" {
        normalizedURL, err := inputvalidator.ValidateSingleURL(*singleURL)
        if err != nil {
            stderr.PrintHintedError(err)
            return
        }
        goodURLs = []string{normalizedURL}
        badURLs = []string{} // No need to handle prompt for single URL
    } else if *inputFile != "" {
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

    // Initialize rate limiter
    limiter := ratelimiter.NewLimiter(*threads, *rate)
    limiter.Start()

// Determine if output should be saved
writeOutput := false
format := "txt"
if *outputFormat != "" {
    writeOutput = true
    if *outputFormat == "json" {
        format = "json"
    }
}

// Output and processing of URLs
spinner := aesthetics.StartSpinner()
aesthetics.PrintInfo("\nEnumerating valid target(s):")
for _, url := range uniqueURLs {
    fmt.Println("\n+", url) 

    // Unified crawl (handles both active and passive internally)
    links := activescanner.Crawl(url, *depth, limiter)

    // Deduplicate links
    uniqueLinks := dedupe.RemoveDuplicates(links)

    // Unified output
    err = outputhandler.SaveLinks(url, uniqueLinks, format, writeOutput)
    if err != nil {
        stderr.PrintOutputError(url, err)
    }

    if writeOutput {
        // Notify user where the output is saved
        fmt.Printf(aesthetics.Green + "\nEnumeration completed for %s and Output is saved to " + aesthetics.Bold + "output/%s.%s\n" + aesthetics.Reset, url, outputhandler.SanitizeFilename(url), format + aesthetics.Reset)
    }
}

spinner.Stop()
// Stop the limiter after crawling is done
limiter.Stop()


}
