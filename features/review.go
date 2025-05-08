package features

import (
	"fmt"
	"os" // added import

	"github.com/Dmaddu/devpilot/client"
)

// GetReviewSummary generates a review for a given PR diff file path.
func GetReviewSummary(prDiffPath string) (string, error) {
	data, err := os.ReadFile(prDiffPath) // read file content
	if err != nil {
		return "", err
	}
	prDiff := string(data)
	prompt := fmt.Sprintf(`
You are a senior code reviewer. Review the following Pull Request diff and provide comprehensive feedback that includes:
- Identification of strengths and weaknesses.
- Assessment for potential security vulnerabilities.
- Identification of code smells.
- Evaluation of adherence to best practices.

PR Diff: %s
`, prDiff)
	cli := client.NewAzureOpenAIClient()
	return cli.SendPrompt(prompt)
}
