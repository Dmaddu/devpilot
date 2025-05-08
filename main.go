// NOTE: Run "go run ." to compile all package files in this directory.
package main

import "fmt"

func fetchCurrentTime() {
	azureEndpoint := "https://in-ep-one.openai.azure.com"
	apiKey := "FEHb9gIIByqzOKz0mczL0rU4heu5veqLu6OMSKuHHVRWMeFR80NlJQQJ99BCAC77bzfXJ3w3AAABACOGQdFM"
	deploymentName := "gpt-4o"
	apiVersion := "2024-05-01-preview"

	client := NewAzureOpenAIClient(azureEndpoint, apiKey, deploymentName, apiVersion)
	// Send prompt to fetch current time via GPT-4
	response, err := client.SendPrompt("What is the current time?")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current time:", response)
}

func main() {
	fmt.Println("Welcome to DevPilot!")
	fetchCurrentTime()
}
