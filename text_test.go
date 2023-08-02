package gochatgptsdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	c := NewConfig(Config{
		OpenAIKey: "",
	})

	resp, err := c.Completions(ModelText{
		Model:  "",
		Prompt: "",
	})

	assert.NotNil(t, err)
	expectedErrType := &ErrorResponse{}
	assert.IsType(t, expectedErrType, err)
	assert.Nil(t, resp)
	expectedRespType := &ModelTextResponse{}
	assert.IsType(t, expectedRespType, resp)
}
