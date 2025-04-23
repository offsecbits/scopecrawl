
---

# scopecrawl Documentation

## Usage Table

| Flag           | Description                                                                | Default  |
|----------------|----------------------------------------------------------------------------|----------|
| `-u`           | Single URL to validate.                                                    |   None   |
| `-f`           | Path to the input file containing a list of URLs (FQDNs) or domains.       |   None   |
| `-o`           | Output format (`txt` or `json`).                                           |   `txt`  |
| `-v`           | Show version information.                                                  |   None   |


---

## Project Structure

The **scopecrawl** project is organized as follows:

- **`go.mod`**: This is the Go module file that contains dependencies and other metadata for the project.
  
- **`go.sum`**: This file contains checksums for the dependencies in your `go.mod` to ensure integrity and consistency.

- **`main.go`**: The entry point for the **scopecrawl** CLI tool. It handles command-line input and orchestrates the functionality of all the utilities.

- **`utils/`**: A folder containing all the utility modules. These are responsible for various tasks like input validation, URL normalization, deduplication, link extraction, output handling, and error management.

    - **`inputvalidator/`**: This module handles reading and validating input URLs or domain names (from a file or direct input).
    
    - **`urlnormalize/`**: This module ensures that URLs are properly normalized (e.g., adding `https://`, ensuring valid syntax).
    
    - **`dedupe/`**: This module removes duplicate URLs from a list to ensure that only unique links are processed.
    
    - **`linkextractor/`**: Extracts all the internal and external links from the HTML of a web page.
    
    - **`outputhandler/`**: This module handles the generation of output files in various formats like `txt` or `json`.
    
    - **`stderr/`**: Contains functions for consistent error handling and reporting to standard error.
  
---
