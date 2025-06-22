package crawler

import (
	"fmt"
	"net/url"
	"sync"
)

type Config struct {
	Pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *Config) AddPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.Pages[normalizedURL]; visited {
		cfg.Pages[normalizedURL]++
		return false
	}

	cfg.Pages[normalizedURL] = 1
	return true
}

func (cfg *Config) PagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.Pages)
}

func Configure(rawBaseURL string, maxConcurrency int, maxPages int) (*Config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &Config{
		Pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		Wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil
}
