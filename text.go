package gochatgptsdk

import (
	"fmt"
	"net/http"
)

func (c *chatgpt) Completions(b ModelText) (*ModelTextResponse, *ErrorResponse) {
	// POST https://api.openai.com/v1/completions
	endpointCompletions := fmt.Sprintf("%s/completions", ChatGPTAPIV1)

	result := ModelTextResponse{}
	errResponse := ErrorResponse{}

	resp, err := DoRequest(c.OpenAIKey).
		SetBody(b).
		SetResult(&result).
		SetError(&errResponse).
		Post(endpointCompletions)
	if err != nil {
		return nil, &errResponse
	}

	// if status code not 200 ok, please send error message
	if resp.StatusCode() != http.StatusOK {
		return nil, &errResponse
	}

	return &result, nil
}
