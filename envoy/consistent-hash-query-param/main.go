package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	for _, name := range user_names {
		for _, org := range user_orgs {
			fmt.Printf("Running %s %s\n", name, org)
			url := fmt.Sprintf("%s/get?user_id=%s&user_org=%s", baseURL, name, org)
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				runConsistentHashQueryParamTest(ctx, cancel, url)
			}(url)
			fmt.Printf("\n")
		}

	}
	wg.Wait()
}

func runConsistentHashQueryParamTest(ctx context.Context, cancel context.CancelFunc, url string) {
	var firstService string
	count := 0

	fmt.Printf("Starting requests to %s\n", url)

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Printf("Error sending request: %v\n", err)
				continue
			}
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
			fmt.Printf("time: %v\n", time.Now())
			fmt.Printf("URL %s!\n", url)
			fmt.Printf("Previous: %s\n", firstService)
			fmt.Printf("Current:  %s\n", currentService)
			fmt.Printf("Requests sent before change: %d\n", count)
			cancel()
			return
		}

		count++
		if count%100 == 0 {
			fmt.Println(".")
		}

		// Stability check
		if count >= 10000 {
			fmt.Printf("\nFinished: Reached %d requests with no change in service (%s). Consistent hashing is stable.\n", count, firstService)
			return
		}
	}
}
