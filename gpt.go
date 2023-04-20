package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
)

func main() {
	client := openai.NewClient("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")

	resp, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "what is your name!",
				},
			},
			Stream: true,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	recv, err := resp.Recv()
	for i := range recv.Choices {

		log.Println(recv.Choices[i])
	}
}
