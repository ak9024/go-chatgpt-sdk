package gochatgptsdk

import "fmt"

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

func (c *chatgpt) ImagesVariations(b ModelImagesVariations) (*ModelImagesResponse[DataURL], error) {
	// POST https/api.openai.com/v1/images/variations
	endpointImagesVariations := fmt.Sprintf("%s/images/variations", ChatGPTAPIV1)

	result := ModelImagesResponse[DataURL]{}

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
		b.ResponseFormat = ResponseFormatURL
	}

	// to prevent nil pointer, we need to check if user empty, fill with default value
	if b.User == "" {
		b.User = "user"
	}

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

func (c *chatgpt) ImagesVariationsB64JSON(b ModelImagesVariations) (*ModelImagesResponse[DataB64JSON], error) {
	// POST https/api.openai.com/v1/images/variations
	endpointImagesVariations := fmt.Sprintf("%s/images/variations", ChatGPTAPIV1)

	result := ModelImagesResponse[DataB64JSON]{}

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
		b.ResponseFormat = ResponseFormatB64JSON
	}

	// to prevent nil pointer, we need to check if user empty, fill with default value
	if b.User == "" {
		b.User = "user"
	}

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
