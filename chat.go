package gochatgptsdk

import "fmt"

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
