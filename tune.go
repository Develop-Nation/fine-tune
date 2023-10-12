package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sashabaranov/go-openai"
)

func fineTune() {
	client := openai.NewClient("secret")

	file, err := client.CreateFile(context.Background(), openai.FileRequest{
		FileName: "dialogue1",
		FilePath: "./dialogue.jsonl",
		Purpose:  "fine-tune",
	})
	if err != nil {
		panic(err)
	}

	job, err := client.CreateFineTuningJob(context.Background(), openai.FineTuningJobRequest{
		TrainingFile: file.ID,
		Model:        openai.GPT3Dot5Turbo,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("created fine tuning job")

	for job.Status != "succeeded" {
		job, err = client.RetrieveFineTuningJob(context.Background(), job.ID)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Minute)
		log.Println(job.Status, job.ID)
	}

	log.Println(job.FineTunedModel)
	log.Printf("%#v", job)
}
