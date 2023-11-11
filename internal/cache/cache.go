package cache

// Implementation of the URL cache
import (
	"sync"
)

// struct represents the metadata with a URL string + metadat payload
type URLMetadata struct {
	URL      string
	MetaData string
}

// struct represents the cache
type URLCache struct {
	cache map[string]URLMetadata
	mu    sync.RWMutex
}

func NewCache() *URLCache {

	return &URLCache{
		cache: make(map[string]URLMetadata),
	}
}

func (c *URLCache) AddURL(url string, metadata URLMetadata) {
	c.mu.Lock()
	// some logic
	defer c.mu.Unlock()
}

// GetURLMetadata
// Add Add+URL + get metadata to main.go
