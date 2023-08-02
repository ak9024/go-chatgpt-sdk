package gochatgptsdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChatCompletions(t *testing.T) {
	c := NewConfig(Config{
		OpenAIKey: "",
	})

	resp, err := c.ChatCompletions(ModelChat{
		Model: "",
		Messages: []Message{
			{
				Role:    "system",
				Content: "",
			},
		},
	})

	assert.NotNil(t, err)
	expectedErrType := &ErrorResponse{}
	assert.IsType(t, expectedErrType, err)
	assert.Nil(t, resp)
	expectedRespType := &ModelChatResponse{}
	assert.IsType(t, expectedRespType, resp)
}
