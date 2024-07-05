package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"strings"
)

const url = "https://publicsuffix.org/list/public_suffix_list.dat"

// fetchList fetches the public suffix list from the given URL
func fetchList(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lines []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// filterLinesByComment filters lines based on the specified comment substring
func filterLinesByComment(lines []string, comment string) []string {
	var endpoints []string
	var captureEndpoints bool

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Check if the line is a comment
		if strings.HasPrefix(trimmedLine, "//") {
			if strings.Contains(trimmedLine, comment) {
				captureEndpoints = true
			}
			continue
		}

		// Capture endpoints if in the relevant comment block
		if captureEndpoints && trimmedLine != "" {
			endpoints = append(endpoints, trimmedLine)
		}

		// Stop capturing when a blank line is encountered
		if trimmedLine == "" {
			captureEndpoints = false
		}
	}

	return endpoints
}

// filterLinesByDomain filters lines based on the specified domain substring
func filterLinesByDomain(lines []string, domain string) []string {
	var endpoints []string
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmedLine, "//") && strings.Contains(trimmedLine, domain) {
			endpoints = append(endpoints, trimmedLine)
		}
	}
	return endpoints
}

func main() {
	// Define command line flags
	commentPtr := flag.String("c", "", "Comment substring to search for")
	domainPtr := flag.String("d", "", "Domain substring to search for")
	flag.Parse()

	if *commentPtr == "" && *domainPtr == "" {
		fmt.Println("Please provide a comment (-c) or domain (-d) to search for")
		return
	}

	lines, err := fetchList(url)
	if err != nil {
		fmt.Println("Error fetching list:", err)
		return
	}

	var endpoints []string
	if *commentPtr != "" {
		endpoints = filterLinesByComment(lines, *commentPtr)
	} else if *domainPtr != "" {
		endpoints = filterLinesByDomain(lines, *domainPtr)
	}

	for _, endpoint := range endpoints {
		fmt.Println(endpoint)
	}
}
