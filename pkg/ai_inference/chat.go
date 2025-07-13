package aiinference

import (
	"context"
	"fmt"
	"log"

	"github.com/openai/openai-go"
)

func ChatWithAI() {
	res, err := Client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		// Messages: []openai.Message{},
		Model: "qwen3-0.6b",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
