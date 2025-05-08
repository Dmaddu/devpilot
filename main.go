// NOTE: Run "go run ." to compile all package files in this directory.
package main

import (
	"fmt"

	"github.com/Dmaddu/devpilot/client"
)

func fetchCurrentTime() {

	client := client.NewAzureOpenAIClient()
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
