// Main entry point of application

package main

import (
	"fmt"
	"regexp"
	"url-caching-app/internal/fetcher"
)

// set up cli to take input for user to add URL to fetch and catche

func main() {

	// // Receive input

	// for {
	// 	fmt.Print("Enter URL ('type exit to quit'): ")
	// 	reader := bufio.NewReader(os.Stdin)
	// 	input, _ := reader.ReadString('\n')
	// 	trimmedInput := strings.TrimSpace(input)
	// 	if trimmedInput == "exit" {
	// 		break
	// 	}
	// 	// check for specific url format

	// 	if isValidURL(trimmedInput) {
	// 		fmt.Println("Valid URL: ", trimmedInput)
	// 		// cache.updateCache(trimmedInput)
	// 	} else {
	// 		fmt.Println("Invalid URL. Please try again. ")
	// 	}

	// 	// call fetcher for URL Metdata

	// 	fmt.Println(trimmedInput)

	// 	cache := cache.NewCache()

	// 	cache.AddURL(trimmedInput)

	// }

	// fetcher logic

	url := "https://seatgeek.com/"
	resp, err := fetcher.FetchMetadata(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Status Code:", resp.Status)

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
