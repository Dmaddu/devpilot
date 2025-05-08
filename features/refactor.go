package features

import (
	"fmt"
	"io/ioutil"

	"github.com/Dmaddu/devpilot/client"
)

// RefactorFile takes a file path and refactors its content into idiomatic Go code.
func RefactorFile(filePath string) (string, error) {
	// Read the file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Construct the prompt for GPT
	prompt := fmt.Sprintf(`
You are an expert Golang developer. Refactor the following code into idiomatic Go code:
- Ensure best practices are followed.
- Optimize for readability and maintainability.
- Include comments where necessary.

Code:
%s
`, string(content))

	// Use the GPT client to send the prompt
	client := client.NewAzureOpenAIClient()
	return client.SendPrompt(prompt)
}
