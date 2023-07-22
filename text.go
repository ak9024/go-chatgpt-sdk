package gochatgptsdk

import "fmt"

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
