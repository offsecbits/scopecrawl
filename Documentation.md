
---

# scopecrawl Documentation

## Usage Table

| Flag           | Description                                                                | Default  |
|----------------|----------------------------------------------------------------------------|----------|
| `-u`           | Single URL to validate.                                                    |   None   |
| `-f`           | Path to the input file containing a list of URLs (FQDNs) or domains.       |   None   |
| `-o`           | Output format (`txt` or `json`).                                           |   None   |
| `-v`           | Show version information.                                                  |   None   |
| `-d`           | Depth of recursive scanning.                                               |     0    |
| `-r`           | Requests per second (rps) rate limit for active scanning.                  |     2    |
| `-t`           | Number of concurrent threads for active scanning.                          |     2    |



---

## Project Structure

The **scopecrawl** project is organized as follows:

- **`go.mod`**: Go module file containing dependencies.
- **`go.sum`**: Dependency integrity checksums.
- **`main.go`**: CLI entry point orchestrating the entire tool's workflow.
- **`utils/`**: Utility modules:
  
    - **`inputvalidator/`**: Reads and validates input URLs (single or file).
    - **`urlnormalize/`**: Normalizes URLs (e.g., adding schemes, correcting format).
    - **`dedupe/`**: Removes duplicate URLs.
    - **`linkextractor/`**: Extracts internal links from HTML.
    - **`outputhandler/`**: Handles output file creation (`txt`, `json`).
    - **`ratelimiter/`**: Manages concurrency and request-per-second limits.
    - **`stderr/`**: Consistent error reporting to standard error output.
    - **`aesthetics/`**: Handles colors, banner, and beautification of tool.
    - **`activescanner/`**: Recursive active scanning implementation.
    - **`filters/`**: Filters on non-html pages and keywords. Work in progress...
  
---
