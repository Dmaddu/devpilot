package analyzer

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Dmaddu/devpilot/client"
)

// GetDocumentationSummary fetches and summarizes documentation from a given URL.
func GetDocumentationSummary(docURL string) (string, error) {
	resp, err := http.Get(docURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch documentation: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch documentation: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read documentation content: %w", err)
	}

	/*prompt := fmt.Sprintf(`
	You are an expert technical writer. Summarize the following documentation content into key points:
	- Main topics covered.
	- Key details and explanations.
	- Detailed Summary.
	- Any notable examples or use cases.

	Documentation Content:
	%s
	`, string(body))*/

	prompt := fmt.Sprintf(`You are an intelligent documentation assistant. Please read all available documentation files starting from the given base link. Recursively explore all subdirectories and summarize the contents of the documentation. Your summary should include:

An overview of what the documentation covers.

Key components, APIs, or features mentioned.

Any setup or configuration steps included.

Notable usage examples, if available.

Any limitations, warnings, or best practices discussed.

Return the summary in clear, structured bullet points or sections.
Documentation Content:
%s
`, string(body))

	client := client.NewAzureOpenAIClient()
	return client.SendPrompt(prompt)
}
