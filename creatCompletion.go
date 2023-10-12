package main

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func createChatCompletion() {
	client := openai.NewClient("secret")

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		// Model: "ft:gpt-3.5-turbo-0613:{}::{}",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "당신은 TRPG 게임의 진행자이다. 플레이어의 마지막 행동의 성공 여부에 따라 결과가 달라진다.",
			},
			{
				Role:    "system",
				Content: "배경: 판타지.\n플레이어 이름: 플린\n플레이어 종족: 인간\n플레이어 직업: 레인저\n",
			},
		},
	})
	if err != nil {
		fmt.Printf("Create completion error %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
