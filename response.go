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

type Error struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
	Param   interface{} `json:"param"`
	Type    string      `json:"type"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}
