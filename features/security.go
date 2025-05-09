package features

import (
	"fmt"

	"github.com/Dmaddu/devpilot/client"
)

// AnalyzeSecurityIssues analyzes the entire repository for potential code smells and security issues.
func AnalyzeSecurityIssues(repoRoot string) (string, error) {
	summary, err := GetArchitectureSummary(repoRoot)
	if err != nil {
		return "", err
	}

	// Format the summary for security analysis.
	formattedSummary := fmt.Sprintf("Repository Analysis:\n%s", summary)

	// Construct a GPT prompt that asks for code smells and security issues, include file name and line number.
	prompt := fmt.Sprintf(`
You are a senior security engineer. Based on the following repository summary, identify syntax errors,code bugs, potential code smells and security issues.
For each issue, provide a brief explanation, specify the file name and line number where the issue might occur, and suggest corrective actions.
%s
`, formattedSummary)

	cli := client.NewAzureOpenAIClient()
	return cli.SendPrompt(prompt)
}
