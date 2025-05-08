// NOTE: Run "go run ." to compile all package files in this directory.
package main

import (
	"fmt"
	"os"

	features "github.com/Dmaddu/devpilot/features"
)

func analyzeRepo(repoPath string) {
	fmt.Printf("Analyzing repository at path: %s\n", repoPath)
	summary, err := features.GetArchitectureSummary(repoPath)
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
		result, err := features.GenerateTestsForRepo(repoPath)
		if err != nil {
			fmt.Println("Error generating tests:", err.Error())
			return
		}
		fmt.Println(result)
	case "review":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go review <pr_link>")
			return
		}
		prLink := os.Args[2]
		review, err := features.GetReviewSummary(prLink)
		if err != nil {
			fmt.Println("Error reviewing PR:", err.Error())
			return
		}
		fmt.Println(review)
	case "docsummary":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go docsummary <doc_url>")
			return
		}
		docURL := os.Args[2]
		summary, err := features.GetDocumentationSummary(docURL)
		if err != nil {
			fmt.Println("Error summarizing documentation:", err.Error())
			return
		}
		fmt.Println(summary)
	case "docgen":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go docgen <repo_path>")
			return
		}
		repoPath := os.Args[2]
		documentation, err := features.GenerateDocumentation(repoPath)
		if err != nil {
			fmt.Println("Error generating documentation:", err.Error())
			return
		}
		fmt.Println(documentation)
	case "refactor":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go refactor <file_path>")
			return
		}
		filePath := os.Args[2]
		refactoredCode, err := features.RefactorFile(filePath)
		if err != nil {
			fmt.Println("Error refactoring file:", err.Error())
			return
		}
		fmt.Println(refactoredCode)
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: analyze, testgen, review, docsummary, docgen, refactor")
	}
}
