package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Issue struct {
	Number      int         `json:"number"`
	PullRequest interface{} `json:"pull_request,omitempty"`
}

func removeLabelFromIssues(owner, repo, label, token string, dryRun bool) error {
	client := &http.Client{}
	page := 1

	for {
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues?labels=%s&state=open&per_page=100&page=%d", owner, repo, label, page)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Authorization", "token "+token)
		req.Header.Add("Accept", "application/vnd.github.v3+json")

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("error fetching issues: %s - %s", resp.Status, string(body))
		}

		var issues []Issue
		if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
			return err
		}

		if len(issues) == 0 {
			break
		}

		for _, issue := range issues {
			typeStr := "Issue"
			if issue.PullRequest != nil {
				typeStr = "PR"
			}

			fmt.Printf("Removing label '%s' from %s #%d...\n", label, typeStr, issue.Number)

			if dryRun {
				fmt.Printf("[Dry-run] Would remove label from #%d\n", issue.Number)
				continue
			}

			deleteURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d/labels/%s", owner, repo, issue.Number, label)
			delReq, err := http.NewRequest("DELETE", deleteURL, nil)
			if err != nil {
				fmt.Printf("Error creating delete request for #%d: %v\n", issue.Number, err)
				continue
			}
			delReq.Header.Add("Authorization", "token "+token)
			delReq.Header.Add("Accept", "application/vnd.github.v3+json")

			delResp, err := client.Do(delReq)
			if err != nil {
				fmt.Printf("Error deleting label from #%d: %v\n", issue.Number, err)
				continue
			}
			defer delResp.Body.Close()

			switch delResp.StatusCode {
			case http.StatusOK, http.StatusNoContent:
				fmt.Printf("Successfully removed from #%d\n", issue.Number)
			case http.StatusNotFound:
				fmt.Printf("Label not found on #%d (already removed?)\n", issue.Number)
			default:
				fmt.Printf("Failed to remove label from #%d: %s\n", issue.Number, delResp.Status)
			}
		}

		page++
	}

	return nil
}

func main() {
	owner := flag.String("owner", "", "GitHub repository owner")
	repo := flag.String("repo", "", "GitHub repository name")
	label := flag.String("label", "lfx-mentorship/docs-analysis", "Label to remove")
	token := flag.String("token", "", "GitHub Personal Access Token")
	dryRun := flag.Bool("dry-run", true, "Enable dry-run mode (do not actually delete labels)")

	flag.Parse()

	if *owner == "" || *repo == "" || *token == "" {
		fmt.Println("Error: owner, repo, and token are required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *dryRun {
		fmt.Println("Dry-run mode enabled. No labels will be actually removed.")
	}

	if err := removeLabelFromIssues(*owner, *repo, *label, *token, *dryRun); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
