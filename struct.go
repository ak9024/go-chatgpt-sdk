package gochatgptsdk

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ModelChat struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	MaxTokens        int       `json:"max_tokens,omitempty"`
	Temperature      int       `json:"temperature,omitempty"`
	TopP             int       `json:"top_p,omitempty"`
	FrequencyPenalty int       `json:"frequency_penalty,omitempty"`
	PresencePenalty  int       `json:"presence_penalty,omitempty"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type ChoiceText struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ModelChatResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type ModelTextResponse struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Model   string       `json:"model"`
	Choices []ChoiceText `json:"choices"`
	Usage   Usage        `json:"usage"`
}

type ModelText struct {
	Model            string `json:"model"`
	Prompt           string `json:"prompt"`
	MaxTokens        int    `json:"max_tokens,omitempty"`
	Temperature      int    `json:"temperature,omitempty"`
	TopP             int    `json:"top_p,omitempty"`
	FrequencyPenalty int    `json:"frequency_penalty,omitempty"`
	PresencePenalty  int    `json:"presence_penalty,omitempty"`
}

type ModelImages struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n,omitempty"`               // default to 1
	Size           string `json:"size,omitempty"`            // default 1024x1024
	ResponseFormat string `json:"response_format,omitempty"` // url or b64_json
	User           string `json:"user,omitempty"`
}

type ModelImagesVariations struct {
	Image          string `json:"image"`                     // must be valid PNG file, less than 4MB, and square
	N              string `json:"n,omitempty"`               // default to 1
	Size           string `json:"size,omitempty"`            // default 1024x1024
	ResponseFormat string `json:"response_format,omitempty"` // url or b64_json
	User           string `json:"user,omitempty"`
}

type DataURL struct {
	URL string `json:"url"`
}

type DataB64JSON struct {
	B64JSON string `json:"b64_json"`
}

type ModelImagesResponse[T DataURL | DataB64JSON] struct {
	Created int64 `json:"created"`
	Data    []T   `json:"data"`
}
