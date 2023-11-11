// Main entry point of application

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// set up cli to take input for user to add URL to fetch and catche

func main() {

	// Receive input

	for {
		fmt.Print("Enter URL ('type exit to quit'): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "exit" {
			break
		}
		// check for specific url format

		if isValidURL(trimmedInput) {
			fmt.Println("Valid URL: ", trimmedInput)
			// cache.updateCache(trimmedInput)
		} else {
			fmt.Println("Invalid URL. Please try again. ")
		}
		fmt.Println(trimmedInput)

		// cache := cache.makecache

	}

	// remove new line delimiter

	// check if url in cache, if not, fetch content and store (call fetch + cache)

	// use goroutne for concurrent fetching

	// fetch multiple urls at a time?
}

func isValidURL(url string) bool {

	urlRegex := `^(http|https)://`

	pattern := regexp.MustCompile(urlRegex)

	return pattern.MatchString(url)
}
