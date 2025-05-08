package features

import (
	"fmt"

	"github.com/Dmaddu/devpilot/client"
)

// GenerateDocumentation generates documentation for a given repository path.
func GenerateDocumentation(repoRoot string) (string, error) {
	summaries, err := AnalyzeRepo(repoRoot)
	if err != nil {
		return "", err
	}

	formatted := FormatForGPT(summaries)

	prompt := fmt.Sprintf(`
You are an expert technical writer. Based on the following repository summary, generate comprehensive documentation:
- Include an overview of the repository.
- Describe the purpose of each package and its key components.
- Provide usage examples for important functions or modules.
- Highlight any configuration or setup steps required.

%s
`, formatted)

	client := client.NewAzureOpenAIClient()
	return client.SendPrompt(prompt)
}
