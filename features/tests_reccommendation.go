package features

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Dmaddu/devpilot/client"
)

// RecommendTestsForPR analyzes the changes provided via a PR URL or local diff file.
// If prInput starts with "file://", the function treats it as a path to a local diff file.
func RecommendTestsForPR(prInput string) (string, error) {
	var content string
	if strings.HasPrefix(prInput, "file://") {
		filePath := strings.TrimPrefix(prInput, "file://")
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return "", err
		}
		content = string(data)
	} else {
		content = prInput
	}

	// Updated prompt instructs GPT to provide test recommendations based on the PR changes
	prompt := fmt.Sprintf(`
You are a senior software engineer specializing in testing strategies.
Analyze the following changes from the pull request:
%s
Based on these changes, identify the modified sections of the code and recommend tests to run, including unit tests, integration tests, and UI tests if applicable.
Provide a detailed testing plan that covers all affected areas.
`, content)

	cli := client.NewAzureOpenAIClient()
	return cli.SendPrompt(prompt)
}
