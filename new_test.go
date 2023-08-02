package gochatgptsdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	c := Config{
		OpenAIKey: "api key",
	}
	expectedType := ""
	assert.IsType(t, expectedType, c.OpenAIKey)
	expectedValue := "api key"
	assert.Equal(t, expectedValue, c.OpenAIKey)
}

func TestInitNewConfig(t *testing.T) {
	c := NewConfig(Config{
		OpenAIKey: "api key",
	})
	expectedType := &chatgpt{}
	assert.IsType(t, expectedType, c)
	expectedValue := "api key"
	assert.Equal(t, expectedValue, c.OpenAIKey)
}
