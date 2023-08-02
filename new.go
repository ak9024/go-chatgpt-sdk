package gochatgptsdk

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
