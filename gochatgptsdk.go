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

func (c *chatgpt) ImagesGenerations(b ModelImages) (*ModelImagesResponse[DataURL], error) {
	// POST https://api.openai.com/v1/images/generations
	endpointImagesGeneratios := fmt.Sprintf("%s/images/generations", ChatGPTAPIV1)

	result := ModelImagesResponse[DataURL]{}

	_, err := DoRequest(c.OpenAIKey).
		SetBody(b).
		SetResult(&result).
		Post(endpointImagesGeneratios)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *chatgpt) ImagesGenerationsB64JSON(b ModelImages) (*ModelImagesResponse[DataB64JSON], error) {
	// POST https/api.openai.com/v1/images/generations
	endpointImagesGeneratios := fmt.Sprintf("%s/images/generations", ChatGPTAPIV1)

	result := ModelImagesResponse[DataB64JSON]{}

	_, err := DoRequest(c.OpenAIKey).
		SetBody(b).
		SetResult(&result).
		Post(endpointImagesGeneratios)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *chatgpt) ImagesVariations(b *ModelImagesVariations) (*ModelImagesResponse[DataURL], error) {
	// POST https/api.openai.com/v1/images/variations
	endpointImagesVariations := fmt.Sprintf("%s/images/variations", ChatGPTAPIV1)

	result := ModelImagesResponse[DataURL]{}

	_, err := DoRequest(c.OpenAIKey).
		SetFile("image", *&b.Image).
		SetFormData(map[string]string{
			"n":               b.N,
			"size":            b.Size,
			"response_format": b.ResponseFormat,
			"user":            b.User,
		}).
		SetResult(&result).
		Post(endpointImagesVariations)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *chatgpt) ImagesVariationsB64JSON(b *ModelImagesVariations) (*ModelImagesResponse[DataB64JSON], error) {
	// POST https/api.openai.com/v1/images/variations
	endpointImagesVariations := fmt.Sprintf("%s/images/variations", ChatGPTAPIV1)

	result := ModelImagesResponse[DataB64JSON]{}

	_, err := DoRequest(c.OpenAIKey).
		SetFile("image", *&b.Image).
		SetFormData(map[string]string{
			"n":               b.N,
			"size":            b.Size,
			"response_format": b.ResponseFormat,
			"user":            b.User,
		}).
		SetResult(&result).
		Post(endpointImagesVariations)
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
