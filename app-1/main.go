package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// GitHubRepo holds the basic repository information we're interested in
type GitHubRepo struct {
	FullName string `json:"full_name"`
	Stars    int    `json:"stargazers_count"`
}

// fetchStars queries GitHub's API for a repository's star count
func fetchStars(user, repo string) (int, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", user, repo)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch repository: %s", resp.Status)
	}

	var repository GitHubRepo
	err = json.NewDecoder(resp.Body).Decode(&repository)
	if err != nil {
		return 0, err
	}

	return repository.Stars, nil
}

// handler is an HTTP server handler that prints the star count of a GitHub repository
// handler is an HTTP server handler that prints the star count of a GitHub repository
func handler(w http.ResponseWriter, r *http.Request) {
	// Change "user" and "repo" to the GitHub user and repository you're interested in
	user := "golang"
	repo := "go"

	stars, err := fetchStars(user, repo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s has %d stars", repo, stars)

	// Log the request
	logRequest(r)
}

// logRequest logs the details of a request to a file in the temp directory
func logRequest(r *http.Request) {
	// Get the temp directory
	tempDir := os.TempDir()

	// Create a log file in the temp directory
	logFile, err := os.CreateTemp(tempDir, "log-*.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Log file created: ", logFile.Name())

	// Write the request details to the log file
	_, err = logFile.WriteString(time.Now().Format(time.RFC3339) + " - " + r.Method + " " + r.URL.String() + "\n")
	if err != nil {
		log.Fatal(err)
	}

	// Close the log file
	err = logFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
