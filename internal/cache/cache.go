package cache

// Implementation of the URL cache

import (
	"bufio"
)

// add func to create a new cache using sync.Map
func makecache(url string) {
	cache := make(map[string]string)

	cache[url] = url

	fmt.Println("Here's the url", cache[url])

}

func readcache(items map) map{

	// read cache somehow
}


// make new cache using sync.map
// using a struct of url data
// add expiring on cached URL
// add func to check if url is in cache
// if url in cache, return message or code
// add func to add url and content to catch
// if url not in cache, add url and content
