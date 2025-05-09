// NOTE: Run "go run ." to compile all package files in this directory.
package main

import (
	"fmt"
	"os"
	"time"

	features "github.com/Dmaddu/devpilot/features"
)

func showLoader(message string, done chan bool) {
	go func() {
		frames := []string{"|", "/", "-", "\\"}
		i := 0
		for {
			select {
			case <-done:
				fmt.Print("\r\033[K") // Clear the line
				return
			default:
				fmt.Printf("\r%s %s", message, frames[i%len(frames)])
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

func showFeaturesAndPrompt() {
	// Print "Welcome to DevPilot CLI!" with a more professional color scheme
	fmt.Println("\033[1;32m=========================================\033[0m")
	fmt.Println("\033[1;32m   🚀 Welcome to DevPilot CLI! 🚀   \033[0m")
	fmt.Println("\033[1;32m=========================================\033[0m")
	featureDescriptions := []struct {
		Key         string
		Description string
	}{
		{"analyze", "🔍 Analyze a repository's architecture"},
		{"refactor", "🔧 Refactor a Go file"},
		{"review", "📝 Review a pull request"},
		{"security-scanner", "🔒 Scan the repository for security issues"},
		{"dependency-analysis", "📦 Analyze and visualize repository dependencies"},
		{"loganalysis", "📊 Analyze log files for errors and suggestions"},
		{"testgen", "🧪 Generate tests for a repository"},
		{"docgen", "📚 Generate documentation for a repository"},
		{"docsummary", "📄 Summarize documentation from a URL"},
		{"test-review", "🧪 Recommend tests for a pull request"},
	}

	for {
		fmt.Println()
		// Print "Available features:" with a subtle separator
		fmt.Println("\033[1;37m-----------------------------------------\033[0m")
		fmt.Println("\033[1;37mAvailable features:\033[0m")
		fmt.Println("\033[1;37m-----------------------------------------\033[0m")
		for i, feature := range featureDescriptions {
			fmt.Printf("%d. \033[1;36m%s\033[0m: %s\n", i+1, feature.Key, feature.Description)
		}
		fmt.Println("0. 🚪 Quit")

		fmt.Print("\n\033[1;32m✨ Enter the number corresponding to the feature you want to use:\033[0m ")
		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			fmt.Println("\033[1;31m👋 Exiting... Thank you for using DevPilot!\033[0m")
			break
		}

		if choice < 1 || choice > len(featureDescriptions) {
			fmt.Println("\033[1;31m❌ Invalid choice. Please select a valid number.\033[0m")
			continue
		}

		feature := featureDescriptions[choice-1].Key

		switch feature {
		case "analyze", "testgen", "docgen":
			fmt.Print("\033[1;32m📂 Enter the repository path:\033[0m ")
			var repoPath string
			fmt.Scanln(&repoPath)
			switch feature {
			case "analyze":
				fmt.Printf("\033[1;34m🔍 Analyzing repository at path: %s\033[0m\n", repoPath)
				done := make(chan bool)
				showLoader("Fetching analysis results", done)
				summary, err := features.GetArchitectureSummary(repoPath)
				done <- true
				if err != nil {
					fmt.Println("Error analyzing repository:", err.Error())
				}
				fmt.Println(summary) // Display formatted response
			case "testgen":
				fmt.Printf("\033[1;34m🧪 Generating tests for repository at path: %s\033[0m\n", repoPath)
				done := make(chan bool)
				showLoader("Generating tests", done)
				result, err := features.GenerateTestsForRepo(repoPath)
				done <- true
				if err != nil {
					fmt.Println("Error generating tests:", err.Error())
					return
				}
				fmt.Println(result) // Display formatted response
			case "docgen":
				fmt.Printf("\033[1;34m📚 Generating documentation for repository at path: %s\033[0m\n", repoPath)
				done := make(chan bool)
				showLoader("Generating documentation", done)
				documentation, err := features.GenerateDocumentation(repoPath)
				done <- true
				if err != nil {
					fmt.Println("Error generating documentation:", err.Error())
					return
				}
				fmt.Println(documentation) // Display formatted response
			}
		case "review":
			fmt.Print("\033[1;32m📄 Enter the PR diff file path:\033[0m ")
			var prPath string
			fmt.Scanln(&prPath)
			done := make(chan bool)
			showLoader("Reviewing PR", done)
			review, err := features.GetReviewSummary(prPath)
			done <- true
			if err != nil {
				fmt.Println("Error reviewing PR:", err.Error())
				return
			}
			fmt.Println(review) // Display formatted response
		case "docsummary":
			fmt.Print("\033[1;32m🌐 Enter the documentation URL:\033[0m ")
			var docURL string
			fmt.Scanln(&docURL)
			done := make(chan bool)
			showLoader("Summarizing documentation", done)
			summary, err := features.GetDocumentationSummary(docURL)
			done <- true
			if err != nil {
				fmt.Println("Error summarizing documentation:", err.Error())
				return
			}
			fmt.Println(summary) // Display formatted response
		case "refactor":
			fmt.Print("\033[1;32m🔧 Enter the file path to refactor:\033[0m ")
			var filePath string
			fmt.Scanln(&filePath)
			done := make(chan bool)
			showLoader("Refactoring file", done)
			refactoredCode, err := features.RefactorFile(filePath)
			done <- true
			if err != nil {
				fmt.Println("Error refactoring file:", err.Error())
				return
			}
			fmt.Println(refactoredCode) // Display formatted response
		case "loganalysis":
			fmt.Print("\033[1;32m📊 Enter the log file path:\033[0m ")
			var logFilePath string
			fmt.Scanln(&logFilePath)
			done := make(chan bool)
			showLoader("Analyzing logs", done)
			result, err := features.AnalyzeLogs(logFilePath)
			done <- true
			if err != nil {
				fmt.Println("Error analyzing logs:", err.Error())
				return
			}
			fmt.Println(result) // Display formatted response
		case "dependency-analysis":
			fmt.Print("\033[1;32m📦 Enter the repository path:\033[0m ")
			var repoPath string
			fmt.Scanln(&repoPath)
			done := make(chan bool)
			showLoader("Analyzing dependencies", done)
			result, err := features.AnalyzeDependencies(repoPath)
			done <- true
			if err != nil {
				fmt.Println("Error analyzing dependencies:", err.Error())
				return
			}
			fmt.Println(result) // Display formatted response
		case "security-scanner":
			fmt.Print("\033[1;32m🔒 Enter the repository path:\033[0m ")
			var repoPath string
			fmt.Scanln(&repoPath)
			done := make(chan bool)
			showLoader("Scanning for security issues", done)
			result, err := features.AnalyzeSecurityIssues(repoPath)
			done <- true
			if err != nil {
				fmt.Println("Error scanning for security issues:", err.Error())
				return
			}
			fmt.Println(result) // Display formatted response
		case "test-review":
			fmt.Print("\033[1;32m🧪 Enter the PR diff file path:\033[0m ")
			var prPath string
			fmt.Scanln(&prPath)
			done := make(chan bool)
			showLoader("Recommending tests for PR", done)
			recommendation, err := features.RecommendTestsForPR(prPath)
			done <- true
			if err != nil {
				fmt.Println("Error recommending tests for PR:", err.Error())
				return
			}
			fmt.Println(recommendation) // Display formatted response
		default:
			fmt.Println("\033[1;31m❌ Unknown feature. Please choose a valid feature from the list.\033[0m")
		}

		// Prompt user to continue or quit
		fmt.Print("\n\033[1;32m✨ Press any key to continue or 'q' to quit:\033[0m ")
		var nextAction string
		fmt.Scanln(&nextAction)
		if nextAction == "q" {
			fmt.Println("\033[1;31m👋 Exiting... Thank you for using DevPilot!\033[0m")
			break
		}

		// Clear the screen before showing available features again
		fmt.Print("\033[H\033[2J")
	}
}

func main() {
	if len(os.Args) < 2 {
		showFeaturesAndPrompt()
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
		fmt.Printf("Analyzing repository at path: %s\n", repoPath)
		done := make(chan bool)
		showLoader("Fetching analysis results", done)
		summary, err := features.GetArchitectureSummary(repoPath)
		done <- true
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
		done := make(chan bool)
		showLoader("Generating tests", done)
		result, err := features.GenerateTestsForRepo(repoPath)
		done <- true
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
		done := make(chan bool)
		showLoader("Reviewing PR", done)
		review, err := features.GetReviewSummary(prLink)
		done <- true
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
		done := make(chan bool)
		showLoader("Summarizing documentation", done)
		summary, err := features.GetDocumentationSummary(docURL)
		done <- true
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
		done := make(chan bool)
		showLoader("Generating documentation", done)
		documentation, err := features.GenerateDocumentation(repoPath)
		done <- true
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
		done := make(chan bool)
		showLoader("Refactoring file", done)
		refactoredCode, err := features.RefactorFile(filePath)
		done <- true
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
		done := make(chan bool)
		showLoader("Analyzing logs", done)
		result, err := features.AnalyzeLogs(logFilePath)
		done <- true
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
	case "test-review":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go test-review <pr_diff_file_path>")
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
		fmt.Println("Available commands: analyze, testgen, review, docsummary, docgen, refactor, loganalysis, security-scanner, dependency-analysis,test-review")
	}
}
