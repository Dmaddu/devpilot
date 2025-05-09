// NOTE: Run "go run ." to compile all package files in this directory.
package main

import (
	"fmt"
	"os"

	features "github.com/Dmaddu/devpilot/features"
)

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
		summary, err := features.GetArchitectureSummary(repoPath)
		if err != nil {
			fmt.Println("Error analyzing repository:", err.Error())
		}
		fmt.Println(summary)
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
	case "loganalysis":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go loganalysis <log_file_path>")
			return
		}
		logFilePath := os.Args[2]
		result, err := features.AnalyzeLogs(logFilePath)
		if err != nil {
			fmt.Println("Error analyzing logs:", err.Error())
			return
		}
		fmt.Println(result)
	case "security-scanner":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go security-scanner <repo_path>")
			return
		}
		repoPath := os.Args[2]
		result, err := features.AnalyzeSecurityIssues(repoPath)
		if err != nil {
			fmt.Println("Error analyzing repository security:", err.Error())
			return
		}
		fmt.Println(result)
	case "dependency-analysis":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go dependency-analysis <repo_path>")
			return
		}
		repoPath := os.Args[2]
		fmt.Println("Starting Intelligent Dependency Analysis with enhanced visualization...")
		result, err := features.AnalyzeDependencies(repoPath)
		if err != nil {
			fmt.Println("Error analyzing repository security:", err.Error())
		}
		fmt.Println(result)
	case "recommend-tests":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go recommend-tests <pr_diff_file_path>")
			return
		}
		prPath := os.Args[2]
		// If the input is a valid file, add "file://" prefix.
		if _, err := os.Stat(prPath); err == nil {
			prPath = "file://" + prPath
		}
		recommendation, err := features.RecommendTestsForPR(prPath)
		if err != nil {
			fmt.Println("Error analyzing PR for tests:", err.Error())
			return
		}
		fmt.Println(recommendation)
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: analyze, testgen, review, docsummary, docgen, refactor, loganalysis, security-scanner, dependency-analysis,recommend-tests")
	}
}
