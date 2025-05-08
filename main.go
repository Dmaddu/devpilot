// NOTE: Run "go run ." to compile all package files in this directory.
package main

import (
	"fmt"
	analyzer "github.com/Dmaddu/devpilot/features"
	"os"
)

func analyzeRepo(repoPath string) {
	fmt.Printf("Analyzing repository at path: %s\n", repoPath)
	analyzer.GetArchitectureSummary(repoPath)
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
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: analyze")
	}
}
