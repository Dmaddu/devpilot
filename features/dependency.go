package features

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/Dmaddu/devpilot/client"
)

// AnalyzeDependencies analyzes and visualizes the repository dependencies intelligently.
func AnalyzeDependencies(repoRoot string) (string, error) {
	// Use the local Go tool to extract dependency data.
	cmd := exec.Command("go", "list", "-m", "-json", "all")
	cmd.Dir = repoRoot
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command and capture the output.
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to execute dependency extraction: %v", err)
	}
	depData := out.String()

	// Build a prompt that includes the extracted dependency data and explicitly requires CVE numbers.
	prompt := fmt.Sprintf(`
You are a senior software engineer specializing in dependency analysis and visualization.
The dependency data below has been extracted locally using "go list -m -json all".
Analyze this data to identify outdated dependencies, potential conflicts, and any known vulnerabilities.
For each dependency with a vulnerability, include its CVE identifier(s) (CVE numbers) in your response.
Provide a detailed visualization plan along with practical suggestions for improvement along with safer versions for each dependency.
 
Dependency Data:
%s
`, depData)

	cli := client.NewAzureOpenAIClient()
	return cli.SendPrompt(prompt)
}
