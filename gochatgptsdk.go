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

func (c *chatgpt) ChatCompletionsWithModelGPT35Turbo(b ModelGPT35Turbo) (*ModelGPT35TurboResponse, error) {
	client := resty.New()

	endpointChatCompletions := fmt.Sprintf("%s/chat/completions", ChatGPTAPIV1)

	result := ModelGPT35TurboResponse{}

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

func (c *chatgpt) CompletionWithModelTextDavinci003(b ModelTextDavinci003) (*ModelTextDavinci003Response, error) {
	client := resty.New()

	endpointCompletions := fmt.Sprintf("%s/completions", ChatGPTAPIV1)

	result := ModelTextDavinci003Response{}

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
