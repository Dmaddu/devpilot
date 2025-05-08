package features

import (
	"fmt"

	"github.com/Dmaddu/devpilot/client"
)

// GenerateTestsForRepo generates a GPT prompt to create Golang tests for the repository.
func GenerateTestsForRepo(repoRoot string) (string, error) {
	// Analyze repository and format summary.
	summaries, err := AnalyzeRepo(repoRoot)
	if err != nil {
		return "", err
	}
	formatted := FormatForGPT(summaries)

	// Construct prompt for GPT.
	prompt := fmt.Sprintf(`
You are an expert Golang developer.
Based on the following repository summary, generate comprehensive tests covering core functionalities and error handling.
Provide the answer as valid Go code without extra explanations.

%s
`, formatted)

	// Use the GPT client to send the prompt.
	client := client.NewAzureOpenAIClient()
	return client.SendPrompt(prompt)
}
