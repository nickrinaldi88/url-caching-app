package cache

// Implementation of the URL cache
import (
	"sync"
	"time"
)

// struct represents the metadata with a URL string + metadat payload
type URLdata struct {
	URL      string
	MetaData string
}

// struct represents the cache
type cachedURL struct {
	URLdata 
	expireAtTimestamp int64
}

func NewCache() *cachedURL{
	stop chan struct{}
	cache map[string]URLdata

	
}

func (c *cachedURL) AddURL(url string, metadata URLdata) {

	mu := sync.Mutex
	// acquire lock for goroutine -> you must use .Unlock for other goroutines to access
	c.mu.Lock()
	//
	defer c.mu.Unlock()
}

// GetURLdata
// Add Add+URL + get metadata to main.go
