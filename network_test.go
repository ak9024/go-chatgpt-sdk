package gochatgptsdk

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var (
	token string
)

func TestNetwork(t *testing.T) {
	_, err := DoRequest(token).Get("https://jsonplaceholder.typicode.com/todos/1")
	assert.Nil(t, err)
}

func TestNetworkType(t *testing.T) {
	req := DoRequest(token)
	expectedType := &resty.Request{}
	assert.IsType(t, expectedType, req)
}
