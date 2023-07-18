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
	client := resty.New()

	endpointChatCompletions := fmt.Sprintf("%s/chat/completions", ChatGPTAPIV1)

	result := ModelChatResponse{}

	_, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(c.OpenAIKey).
		SetBody(b).
		SetResult(&result).
		Post(endpointChatCompletions)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *chatgpt) Completions(b ModelText) (*ModelTextResponse, error) {
	client := resty.New()

	endpointCompletions := fmt.Sprintf("%s/completions", ChatGPTAPIV1)

	result := ModelTextResponse{}

	_, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(c.OpenAIKey).
		SetBody(b).
		SetResult(&result).
		Post(endpointCompletions)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
