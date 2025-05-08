package features

import (
	"fmt"
	"os"

	"github.com/Dmaddu/devpilot/client"
)

// AnalyzeLogs reads the log file from the given path, analyzes its content to find errors and suggests actionable items.
func AnalyzeLogs(logFilePath string) (string, error) {
	data, err := os.ReadFile(logFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read log file: %w", err)
	}
	logs := string(data)

	// Construct the prompt for GPT to analyze the logs.
	prompt := fmt.Sprintf(`
You are an expert systems engineer. Analyze the following log file content, identify any error patterns, and provide actionable recommendations to resolve them. Your analysis should include:
- A summary of detected errors or issues.
- Specific corrective actions to fix the problems.
- Suggestions for improving overall system stability.

Log file content:
%s
`, logs)
	cli := client.NewAzureOpenAIClient()
	return cli.SendPrompt(prompt)
}
