# go-chatgpt-sdk

## Install

```bash
go get -u github.com/ak9024/go-chatgpt-sdk
```

## Usage

```go
import (
    "fmt"

	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

func main() {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: "",
	})

	resp, _ := c.ChatCompletionsWithModelGPT35Turbo(gochatgptsdk.ModelGPT35Turbo{
		Model: gochatgptsdk.GPT35Turbo,
		Messages: []gochatgptsdk.Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant",
			},
			{
				Role:    "user",
				Content: "How the weather today?",
			},
		},
	})

	fmt.Println(resp.Choices[len(resp.Choices)-1].Message.Content)
}
```
