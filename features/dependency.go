package features

import (
	"fmt"

	"github.com/Dmaddu/devpilot/client"
)

// AnalyzeDependencies analyzes and visualizes the repository dependencies intelligently.
func AnalyzeDependencies(repoRoot string) (string, error) {
	prompt := fmt.Sprintf(`
You are a senior software engineer specializing in dependency analysis and visualization. 
Analyze the repository located at: %s
Identify outdated dependencies, potential conflicts, any known CVEs associated with existing dependencies, and provide a detailed visualization plan along with suggestions for improvement.
`, repoRoot)

	cli := client.NewAzureOpenAIClient()
	return cli.SendPrompt(prompt)
}
