package gochatgptsdk

import "github.com/go-resty/resty/v2"

// DoRequest(t string) to compose HTTP Request
func DoRequest(t string) *resty.Request {
	client := resty.New()
	return client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(t)
}
