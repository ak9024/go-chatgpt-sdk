package gochatgptsdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type chatgpt struct {
	OpenAIKey string
}

type Config struct {
	OpenAIKey string
}

func NewConfig(c Config) *chatgpt {
	return &chatgpt{
		OpenAIKey: c.OpenAIKey,
	}
}

func (c *chatgpt) ChatCompletions(b ModelChat) (*ModelChatResponse, error) {
	// POST https://api.openai.com/v1/chat/completions
	endpointChatCompletions := fmt.Sprintf("%s/chat/completions", ChatGPTAPIV1)

	result := ModelChatResponse{}

	_, err := DoRequest(c.OpenAIKey).
		SetBody(b).
		SetResult(&result).
		Post(endpointChatCompletions)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *chatgpt) Completions(b ModelText) (*ModelTextResponse, error) {
	// POST https://api.openai.com/v1/completions
	endpointCompletions := fmt.Sprintf("%s/completions", ChatGPTAPIV1)

	result := ModelTextResponse{}

	_, err := DoRequest(c.OpenAIKey).
		SetBody(b).
		SetResult(&result).
		Post(endpointCompletions)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DoRequest(t string) to compose HTTP Request
func DoRequest(t string) *resty.Request {
	client := resty.New()
	return client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(t)
}
