package gochatgptsdk

import (
	"fmt"
	"net/http"
)

func (c *chatgpt) ChatCompletions(b ModelChat) (*ModelChatResponse, *ErrorResponse) {
	// POST https://api.openai.com/v1/chat/completions
	endpointChatCompletions := fmt.Sprintf("%s/chat/completions", ChatGPTAPIV1)

	result := ModelChatResponse{}
	errResponse := ErrorResponse{}

	resp, err := DoRequest(c.OpenAIKey).
		SetBody(b).
		SetResult(&result).
		SetError(&errResponse).
		Post(endpointChatCompletions)
	if err != nil {
		return nil, &errResponse
	}

	// if status code not 200 ok, please send error message
	if resp.StatusCode() != http.StatusOK {
		return nil, &errResponse
	}

	return &result, nil
}
