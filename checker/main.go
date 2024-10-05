package main

import (
	"fmt"
	"net/http"
	"time"
)

type SiteConfig struct {
	URL             string
	AcceptableCodes []int
	Frequency       int
}

type Result struct {
	URL    string
	Up     bool
	Status int
}

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

// DefaultClient interits HttpClient

type DefaultClient struct {
}

func (c *DefaultClient) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

func check(config SiteConfig, client HttpClient, results chan<- Result) {
	resp, err := client.Get(config.URL)
	result := Result{
		URL:    config.URL,
		Up:     false,
		Status: resp.StatusCode,
	}

	if err != nil {
		results <- result
		return
	}

	defer resp.Body.Close()

	for _, code := range config.AcceptableCodes {
		if resp.StatusCode == code {
			result.Up = true
			break
		}
	}

	results <- result

}

func scheduleCheck(config SiteConfig, client HttpClient, results chan<- Result) {
	go func() {

		ticker := time.NewTicker(time.Duration(config.Frequency) * time.Second)

		for {
			select {
			case <-ticker.C:
				check(config, client, results)
			}
		}

	}()
}

func main() {

	sites := []SiteConfig{
		{
			URL:             "https://google.com",
			AcceptableCodes: []int{http.StatusOK},
			Frequency:       5,
		}, {
			URL:             "https://bing.com",
			AcceptableCodes: []int{http.StatusOK},
			Frequency:       4,
		}, {
			URL:             "https://go.dev",
			AcceptableCodes: []int{http.StatusOK},
			Frequency:       2,
		},
	}

	results := make(chan Result)

	client := &DefaultClient{}

	for _, site := range sites {
		scheduleCheck(site, client, results)
	}

	for result := range results {
		if result.Up {
			fmt.Println(result.URL, "is up with status code", result.Status)
		} else {
			fmt.Println(result.URL, "is down with status code", result.Status)
		}
	}

}
