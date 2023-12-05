// Implementation of URL fetching
package fetcher

import (
	"fmt"
	"net/http"
)

// func fetches content of a url and returns it

func FetchMetadata(url string) (*http.Response, error) {

	// responseMap := make(map[string]interface{})

	resp, err := http.Head(url)
	if err != nil {
		return nil, err
	}

	// fetch headers

	fmt.Println("Headers:\n")

	header := resp.Header
	fmt.Println(header)
	// fmt.Println("\n")
	// fmt.Println(reflect.TypeOf(header))

	defer resp.Body.Close()

	// create fields slice

	fields := []string{
		"Geo Country Code",
		"Cache Control",
		"Last Modified",
		"Etag",
		"ContentType",
		"ContentLength",
		"Status-Code to check if website is up or down",
		"Set-cookie",
	}

	// // loop to update map
	for key, _ := range header {

		found := containsString(fields, key)

		// fmt.Printf("%s: %s\n", key, value)
		// fmt.Println(reflect.TypeOf(key))
		if found {
			fmt.Println(key)
		}

	}
	// update map with
	//
	// Geo Country Code
	// Cache Control
	// Last Modified
	// Etag
	// ContentType
	// ContentLength
	// Status-Code to check if website is up or down
	// Set-cookie
	// return resp, err
	return resp, err
}

func containsString(slice []string, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

func main() {

	url := "https://seatgeek.com/"

	FetchMetadata(url)
}

// what kind of content would I want?

// Params to fetch
// Content-Type
// Date
// Server
// transfer-encoding
// content-encoding
//
