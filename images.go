package gochatgptsdk

import (
	"fmt"
	"net/http"
)

func (c *chatgpt) ImagesGenerations(b ModelImages) (*ModelImagesResponse[DataURL], *ErrorResponse) {
	// POST https://api.openai.com/v1/images/generations
	endpointImagesGeneratios := fmt.Sprintf("%s/images/generations", ChatGPTAPIV1)

	result := ModelImagesResponse[DataURL]{}
	errorResponse := ErrorResponse{}

	resp, err := DoRequest(c.OpenAIKey).
		SetBody(b).
		SetResult(&result).
		SetError(&errorResponse).
		Post(endpointImagesGeneratios)
	if err != nil {
		return nil, &errorResponse
	}

	// if status code not 200 ok, please send error message
	if resp.StatusCode() != http.StatusOK {
		return nil, &errorResponse
	}

	return &result, nil
}

func (c *chatgpt) ImagesGenerationsB64JSON(b ModelImages) (*ModelImagesResponse[DataB64JSON], *ErrorResponse) {
	// POST https/api.openai.com/v1/images/generations
	endpointImagesGeneratios := fmt.Sprintf("%s/images/generations", ChatGPTAPIV1)

	result := ModelImagesResponse[DataB64JSON]{}
	errorResponse := ErrorResponse{}

	resp, err := DoRequest(c.OpenAIKey).
		SetBody(b).
		SetResult(&result).
		SetError(&errorResponse).
		Post(endpointImagesGeneratios)
	if err != nil {
		return nil, &errorResponse
	}

	// if status code not 200 ok, please send error message
	if resp.StatusCode() != http.StatusOK {
		return nil, &errorResponse
	}

	return &result, nil
}

func (c *chatgpt) ImagesVariations(b ModelImagesVariations) (*ModelImagesResponse[DataURL], *ErrorResponse) {
	// POST https/api.openai.com/v1/images/variations
	endpointImagesVariations := fmt.Sprintf("%s/images/variations", ChatGPTAPIV1)

	result := ModelImagesResponse[DataURL]{}
	errResponse := ErrorResponse{}

	// set default images variations
	d := defaultImagesVariations(b, ResponseFormatURL)

	resp, err := DoRequest(c.OpenAIKey).
		SetFile("image", *&d.Image).
		SetFormData(map[string]string{
			"n":               d.N,
			"size":            d.Size,
			"response_format": d.ResponseFormat,
			"user":            d.User,
		}).
		SetResult(&result).
		SetError(&errResponse).
		Post(endpointImagesVariations)
	if err != nil {
		return nil, &errResponse
	}

	// if status code not 200 ok, please send error message
	if resp.StatusCode() != http.StatusOK {
		return nil, &errResponse
	}

	return &result, nil
}

func (c *chatgpt) ImagesVariationsB64JSON(b ModelImagesVariations) (*ModelImagesResponse[DataB64JSON], *ErrorResponse) {
	// POST https/api.openai.com/v1/images/variations
	endpointImagesVariations := fmt.Sprintf("%s/images/variations", ChatGPTAPIV1)

	result := ModelImagesResponse[DataB64JSON]{}
	errResponse := ErrorResponse{}

	// set default images variations
	d := defaultImagesVariations(b, ResponseFormatB64JSON)

	resp, err := DoRequest(c.OpenAIKey).
		SetFile("image", *&d.Image).
		SetFormData(map[string]string{
			"n":               d.N,
			"size":            d.Size,
			"response_format": d.ResponseFormat,
			"user":            d.User,
		}).
		SetResult(&result).
		SetError(&errResponse).
		Post(endpointImagesVariations)
	if err != nil {
		return nil, &errResponse
	}

	// if status code not 200 ok, please send error message
	if resp.StatusCode() != http.StatusOK {
		return nil, &errResponse
	}

	return &result, nil
}

func defaultImagesVariations(b ModelImagesVariations, f string) *ModelImagesVariations {
	// to prevent nil pointer, we need to check if n empty, fill with default value
	if b.N == "" {
		b.N = "1"
	}

	// to prevent nil pointer, we need to check if size empty, fill with default value
	if b.Size == "" {
		b.Size = Size1024
	}

	// to prevent nil pointer, we need to check if response_format empty, fill with default value
	if b.ResponseFormat == "" {
		b.ResponseFormat = f
	}

	// to prevent nil pointer, we need to check if user empty, fill with default value
	if b.User == "" {
		b.User = "user"
	}

	*&b.Image = b.Image

	return &b
}
