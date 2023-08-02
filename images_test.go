package gochatgptsdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImagesGenerations(t *testing.T) {
	c := NewConfig(Config{
		OpenAIKey: "",
	})

	resp, err := c.ImagesGenerations(ModelImages{
		Prompt: "",
	})

	assert.NotNil(t, err)
	expectedErrType := &ErrorResponse{}
	assert.IsType(t, expectedErrType, err)
	assert.Nil(t, resp)
	expectedRespType := &ModelImagesResponse[DataURL]{}
	assert.IsType(t, expectedRespType, resp)
}

func TestImagesGenerationsBase64JSON(t *testing.T) {
	c := NewConfig(Config{
		OpenAIKey: "",
	})

	resp, err := c.ImagesGenerationsB64JSON(ModelImages{
		Prompt: "",
	})

	assert.NotNil(t, err)
	expectedErrType := &ErrorResponse{}
	assert.IsType(t, expectedErrType, err)
	assert.Nil(t, resp)
	expectedRespType := &ModelImagesResponse[DataB64JSON]{}
	assert.IsType(t, expectedRespType, resp)
}
