package gochatgptsdk

import (
	"fmt"
	"net/http"
)

func (c *chatgpt) ChatCompletions(b ModelChat) (*ModelChatResponse, *ErrorResponse) {
	// POST https://api.openai.com/v1/chat/completions
	endpointChatCompletions := fmt.Sprintf("%s/chat/completions", ChatGPTAPIV1)

	result := ModelChatResponse{}
	errMessage := ErrorResponse{}

	resp, err := DoRequest(c.OpenAIKey).
		SetBody(b).
		SetResult(&result).
		SetError(&errMessage).
		Post(endpointChatCompletions)
	if err != nil {
		return nil, &errMessage
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, &errMessage
	}

	return &result, nil
}
