package stderr

import (
	"flag"
	"fmt"
	"os"
)

func Usage(version string) {
	fmt.Fprintf(os.Stderr, "\nScopeCrawl %s - A modern passive web crawler\n", version)
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  scopecrawl -u <URL> | -f <inputfile>\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nHint: Use -h for help or check the documentation.\n")
}

func PrintVersion(version string) {
	fmt.Printf("ScopeCrawl version %s\n", version)
}

func PrintErrorAndUsage(err error, version string) {
	fmt.Fprintf(os.Stderr, "\n[!] Error: %v\n", err)
	Usage(version)
}

func PromptMissingInput(version string) {
	fmt.Fprintln(os.Stderr, "\n[!] Please provide either a URL (-u) or an input file (-f).")
	Usage(version)
}

func PrintHintedError(err error) {
	fmt.Fprintf(os.Stderr, "\n[!] Error: %v\n", err)
	fmt.Fprintln(os.Stderr, "Hint: Use -h for help or check documentation.")
}

func PrintFileError(err error) {
	fmt.Fprintf(os.Stderr, "\n[!] Error reading and validating input file: %v\n", err)
}

func PrintFetchError(url string, err error) {
	fmt.Fprintf(os.Stderr, "  [!] Error fetching HTML for %s: %v\n", url, err)
}

func PrintExtractError(url string, err error) {
	fmt.Fprintf(os.Stderr, "  [!] Error extracting links from %s: %v\n", url, err)
}
// PrintOutputError prints an error that occurred while saving output.
func PrintOutputError(url string, err error) {
	fmt.Fprintf(os.Stderr, "Error writing output for %s: %v\n", url, err)
}
