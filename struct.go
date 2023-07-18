package gochatgptsdk

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ModelGPT35Turbo struct {
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

type ChoiceTextDavinci003 struct {
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

type ModelGPT35TurboResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type ModelTextDavinci003Response struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChoiceTextDavinci003 `json:"choices"`
	Usage   Usage                  `json:"usage"`
}

type ModelTextDavinci003 struct {
	Model            string `json:"model"`
	Prompt           string `json:"prompt"`
	MaxTokens        int    `json:"max_tokens,omitempty"`
	Temperature      int    `json:"temperature,omitempty"`
	TopP             int    `json:"top_p,omitempty"`
	FrequencyPenalty int    `json:"frequency_penalty,omitempty"`
	PresencePenalty  int    `json:"presence_penalty,omitempty"`
}
