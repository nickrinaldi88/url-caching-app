// Main entry point of application

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// set up cli to take input for user to add URL to fetch and catche

func main() {

	// Receive input

	for {
		fmt.Print("Enter URL ('type exit to quit'): ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if input == "exit" {
			break
		}
		// check for specific url format
		if err != nil {
			fmt.Println("An error occured: Not a valid URL. Please try again")
			return
		}
		input = strings.TrimSuffix(input, "\n")
		fmt.Println(input)

		cache := cache.makecache

	}

	// remove new line delimiter

	// check if url in cache, if not, fetch content and store (call fetch + cache)

	// use goroutne for concurrent fetching

	// fetch multiple urls at a time?
}

// error handling and logic
// tests
