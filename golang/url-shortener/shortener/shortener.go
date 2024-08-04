package shortener

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
)

var (
	urlStore = make(map[string]string)
	mu       sync.Mutex
)

func GenerateShortURL() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func StoreURL(shortURL, originalURL string) {
	mu.Lock()
	defer mu.Unlock()
	urlStore[shortURL] = originalURL
}

func GetOriginalURL(shortURL string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	originalURL, exists := urlStore[shortURL]
	return originalURL, exists
}
