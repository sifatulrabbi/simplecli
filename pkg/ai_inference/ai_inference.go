package aiinference

import (
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

/*
curl http://localhost:8089/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "qwen3-0.6b",
    "messages": [
      { "role": "system", "content": "Always answer in rhymes. Today is Thursday" },
      { "role": "user", "content": "What day is it today?" }
    ],
    "temperature": 0.7,
    "max_tokens": -1,
    "stream": false
}'
*/

var (
	BASE_URL = "http://locahost:8089"
	Client   openai.Client
)

func init() {
	Client = openai.NewClient(option.WithBaseURL(BASE_URL))
	fmt.Println(Client.Chat.Options)
}
