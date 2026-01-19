package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

// Send request to Envoy admin API to retrieve listener.\.$.downstream_cx_active stat
const statFilter = "^listener\\.(^admin|.*19001|.*19003|.*worker).*\\.downstream_cx_active"

var (
	re = regexp.MustCompile(statFilter)
)

func main() {
	if m := re.MatchString("listener.0.0.0.0_10080.downstream_cx_active"); !m {
		log.Fatalf("regex did not match expected string")
	}
	t, err := getTotalConnections()
	if err != nil {
		log.Fatalf("error getting total connections: %v", err)
	}

	log.Printf("Total Connections: %d", *t)
}

func getTotalConnections() (*int, error) {
	// urlencode the filter parameter
	filter := url.QueryEscape(statFilter)
	resp, err := http.Get(fmt.Sprintf("http://%s:%d//stats?filter=%s&format=json",
		"localhost", 19000, filter))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		body := make([]byte, 1024)
		n, _ := resp.Body.Read(body)
		log.Printf("response body: %s", string(body[:n]))
		return nil, fmt.Errorf("unexpected response status: %s body: %s", resp.Status, string(body[:n]))
	}
	// Define struct to decode JSON response into; expecting a single stat in the response in the format:
	// {"stats":[{"name":"server.total_connections","value":123}]}
	var r *struct {
		Stats []struct {
			Name  string
			Value int
		}
	}

	// Decode JSON response into struct
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	// Defensive check for empty stats
	if len(r.Stats) == 0 {
		return nil, fmt.Errorf("no stats found")
	}

	// Log and return total connections
	c := r.Stats[0].Value
	log.Printf("total connections: %d", c)
	return &c, nil
}
