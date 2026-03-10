package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	envoyProxy  = "http://localhost:10000"
	envoyAdmin  = "http://localhost:9901"
	numRequests = 100
)

var trackedStats = []string{
	"cluster.service_echo.upstream_rq_retry",
	"cluster.service_echo.upstream_rq_retry_overflow",
	"cluster.service_echo.upstream_rq_5xx",
	"cluster.service_echo.upstream_rq_pending_overflow",
	"cluster.service_echo.retry_or_shadow_abandoned",
}

func main() {
	fmt.Println("=== Retry Budget Test ===")
	fmt.Println("Config: budgetPercent=0%, minRetryConcurrency=0 => all retries should be blocked")

	// 1. Initial Statistics
	fmt.Println("\n--- Initial Stats ---")
	before := fetchStats()
	printStats(before)

	// 2. Send requests sequentially so each one reaches the upstream and gets a real 5xx
	// response (not an Envoy-generated pending-overflow 503). Only real upstream 5xx
	// responses trigger the retry logic and thus the retry budget check.
	// echo-basic returns 503 when the path is /status/503.
	fmt.Printf("\n--- Sending %d Sequential Requests to /status/503 ---\n", numRequests)
	for i := 0; i < numRequests; i++ {
		req, _ := http.NewRequest("GET", envoyProxy+"/status/503", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("Req %d: error: %v\n", i, err)
			continue
		}
		defer resp.Body.Close()
		io.ReadAll(resp.Body)
	}

	// 3. Check Statistics after load
	fmt.Println("\n--- Stats After Load ---")
	after := fetchStats()
	printStats(after)

	// 4. Compute deltas and verify retry budget behavior
	fmt.Println("\n--- Delta & Verification ---")
	retries := delta(before, after, "cluster.service_echo.upstream_rq_retry")
	retryOverflow := delta(before, after, "cluster.service_echo.upstream_rq_retry_overflow")
	rq5xx := delta(before, after, "cluster.service_echo.upstream_rq_5xx")
	abandoned := delta(before, after, "cluster.service_echo.retry_or_shadow_abandoned")

	fmt.Printf("  upstream_rq_5xx              (upstream 5xx responses):   %d\n", rq5xx)
	fmt.Printf("  upstream_rq_retry            (retries actually sent):     %d\n", retries)
	fmt.Printf("  upstream_rq_retry_overflow   (retries blocked by budget): %d\n", retryOverflow)
	fmt.Printf("  retry_or_shadow_abandoned:                                 %d\n", abandoned)

	fmt.Println()
	pass := true

	// With retryBudget budgetPercent=0 and minRetryConcurrency=0, every attempted
	// retry should be rejected and counted in upstream_rq_retry_overflow.
	if retryOverflow == 0 {
		fmt.Println("FAIL: upstream_rq_retry_overflow is 0 — retry budget does not appear to be blocking retries")
		pass = false
	} else {
		fmt.Printf("PASS: %d retries were blocked by the retry budget (upstream_rq_retry_overflow)\n", retryOverflow)
	}

	if retries > 0 {
		fmt.Printf("FAIL: %d retries went through despite budget=0%%\n", retries)
		pass = false
	} else {
		fmt.Println("PASS: No retries went through (upstream_rq_retry == 0)")
	}

	if pass {
		fmt.Println("\n=== OVERALL: PASS - Retry budget is working as expected ===")
	} else {
		fmt.Println("\n=== OVERALL: FAIL - Retry budget behavior is unexpected ===")
	}
}

func fetchStats() map[string]int64 {
	resp, err := http.Get(envoyAdmin + "/stats")
	if err != nil {
		fmt.Printf("Error fetching stats: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	lines := strings.Split(string(body), "\n")

	result := make(map[string]int64)
	for _, stat := range trackedStats {
		result[stat] = 0
		for _, line := range lines {
			if strings.HasPrefix(line, stat+":") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					val, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
					if err == nil {
						result[stat] = val
					}
				}
				break
			}
		}
	}
	return result
}

func printStats(stats map[string]int64) {
	for _, stat := range trackedStats {
		fmt.Printf("  %s: %d\n", stat, stats[stat])
	}
}

func delta(before, after map[string]int64, key string) int64 {
	return after[key] - before[key]
}
