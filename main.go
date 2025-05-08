// NOTE: Run "go run ." to compile all package files in this directory.
package main

import (
	"fmt"
	"os"

	analyzer "github.com/Dmaddu/devpilot/features"
)

func analyzeRepo(repoPath string) {
	fmt.Printf("Analyzing repository at path: %s\n", repoPath)
	summary, err := analyzer.GetArchitectureSummary(repoPath)
	if err != nil {
		fmt.Println("Error analyzing repository:", err.Error())
	}
	fmt.Println(summary)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "analyze":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go analyze <repo_path>")
			return
		}
		repoPath := os.Args[2]
		analyzeRepo(repoPath)
	case "testgen":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go testgen <repo_path>")
			return
		}
		repoPath := os.Args[2]
		result, err := analyzer.GenerateTestsForRepo(repoPath)
		if err != nil {
			fmt.Println("Error generating tests:", err.Error())
			return
		}
		fmt.Println(result)
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: analyze, testgen")
	case "review":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go review <pr_link>")
			return
		}
		prLink := os.Args[2]
		review, err := analyzer.GetReviewSummary(prLink)
		if err != nil {
			fmt.Println("Error reviewing PR:", err.Error())
			return
		}
		fmt.Println(review)
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: analyze, review")
	}
}
