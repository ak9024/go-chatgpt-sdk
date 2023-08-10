package gochatgptsdk

type Config struct {
	OpenAIKey string
}

type chatgpt struct {
	OpenAIKey string
}

type Chatgpt interface {
	ChatCompletions(b ModelChat) (*ModelChatResponse, *ErrorResponse)
	Completions(b ModelText) (*ModelTextResponse, *ErrorResponse)
	ImagesGenerations(b ModelImages) (*ModelImagesResponse[DataURL], *ErrorResponse)
	ImagesGenerationsB64JSON(b ModelImages) (*ModelImagesResponse[DataB64JSON], *ErrorResponse)
	ImagesVariations(b ModelImagesVariations) (*ModelImagesResponse[DataURL], *ErrorResponse)
	ImagesVariationsB64JSON(b ModelImagesVariations) (*ModelImagesResponse[DataB64JSON], *ErrorResponse)
}

func NewConfig(c Config) Chatgpt {
	return &chatgpt{
		OpenAIKey: c.OpenAIKey,
	}
}
