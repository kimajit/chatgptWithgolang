package main

import (
	"context"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("missing Api key ")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Choices[0].Text)
}
