package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type EchoResponse struct {
	Service string `json:"service"`
}

func main() {
	user_names := []string{"alice", "bob", "charlie", "dave", "eve"}
	user_orgs := []string{"org1", "org2", "org3", "org4", "org5"}
	baseURL := "http://localhost:10000"
	// wait for localhost:10000 to be available
	fmt.Printf("Waiting for localhost:10000 to be available...\n")
	for {
		resp, err := http.Get(baseURL)
		if err == nil && resp.StatusCode == http.StatusOK {
			fmt.Printf("localhost:10000 is available!\n")
			break
		}
		fmt.Printf(".")
	}

	var wg sync.WaitGroup
	for _, name := range user_names {
		for _, org := range user_orgs {
			fmt.Printf("Running %s %s\n", name, org)
			url := fmt.Sprintf("%s/get?user_id=%s&user_org=%s", baseURL, name, org)
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				runConsistentHashQueryParamTest(url)
			}(url)
			fmt.Printf("\n")
		}

	}
	wg.Wait()
}

func runConsistentHashQueryParamTest(url string) {
	var firstService string
	count := 0

	fmt.Printf("Starting requests to %s\n", url)

	for {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			continue
		}

		var echo EchoResponse
		if err := json.Unmarshal(body, &echo); err != nil {
			fmt.Printf("Warning: could not parse JSON response: %v\n", err)
			continue
		}

		currentService := echo.Service
		if currentService == "" {
			fmt.Printf("Warning: service name is empty in response\n")
		}

		if firstService == "" {
			firstService = currentService
			fmt.Printf("Initial service: %s\n", firstService)
		} else if currentService != firstService {
			fmt.Printf("\nDETECTED CHANGE!\n")
			fmt.Printf("Previous: %s\n", firstService)
			fmt.Printf("Current:  %s\n", currentService)
			fmt.Printf("Requests sent before change: %d\n", count)
			break
		}

		count++
		if count%100 == 0 {
			fmt.Printf(".")
		}

		// Stability check
		if count >= 10000 {
			fmt.Printf("\nFinished: Reached %d requests with no change in service (%s). Consistent hashing is stable.\n", count, firstService)
			break
		}
	}
}
