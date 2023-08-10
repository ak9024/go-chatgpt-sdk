package gochatgptsdk

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
