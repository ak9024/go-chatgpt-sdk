![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/ak9024/go-chatgpt-sdk/.github%2Fworkflows%2Ftest.yml)
![GitHub top language](https://img.shields.io/github/languages/top/ak9024/go-chatgpt-sdk)
![GitHub License](https://img.shields.io/github/license/ak9024/go-chatgpt-sdk)

# go-chatgpt-sdk

This Library Provides Unofficial Go Client SDK for OpenAI API

### Install

```bash
go get -u github.com/ak9024/go-chatgpt-sdk
```

### Model Support

- [x] Chat
- [x] Text
- [x] Image
- [ ] Moderations
- [ ] Audio

### Prerequisite

- [Go](https://go.dev/doc/install)
- [OpenAI Key](https://platform.openai.com/account/api-keys)

### Usage with model chat

> to usage with model chat, please usage `c.ChatCompletions`

```go
package main

import (
	"fmt"

	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

func main() {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: "",
	})

	resp, _ := c.ChatCompletions(gochatgptsdk.ModelChat{
		Model: "gpt-3.5-turbo",
		Messages: []gochatgptsdk.Message{
			{
				Role:    "system",
				Content: "You are a Software engineer",
			},
			{
				Role:    "user",
				Content: "Please create a simple function to using Go language",
			},
		},
	})

	fmt.Println(resp.Choices)
}
```

### Usage with model text

> To usage with model text, please usage `c.Completions()`

```go
package main

import (
	"fmt"

	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

func main() {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: "",
	})

	resp, _ := c.Completions(gochatgptsdk.ModelText{
		Model:     "text-davinci-003",
		Prompt:    "What the weather Kota Palu for today?",
		MaxTokens: 100, // max generates of word
	})

	fmt.Println(resp.Choices)
}
```

### Usage with model images

> Create images, please use `c.ImagesGenerations`, but if you want generate base64 image usage `c.ImageGenerationsB64JSON`

```go
package main

import (
	"fmt"

	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

func main() {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: "",
	})

	resp, _ := c.ImagesGenerations(gochatgptsdk.ModelImages{
		Prompt: "Sunset sketch art",
		N:      5, // generates 5 images
		Size:   gochatgptsdk.Size512, // with size 512x512
	})

	fmt.Println(resp.Data)
}
```

> Create images variations of a given image, please use `c.ImagesVariations()` but if you want generate base64 image usage `c.ImageVariationsB64JSON`

```go
package main

import (
	"fmt"

	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

func main() {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: "",
	})

	resp, _ := c.ImagesVariations(gochatgptsdk.ModelImagesVariations{
		Image: "./path/to/example-img.png", // please suitable with your path image
		N:     "2",                  // generate 2 images
		Size:  gochatgptsdk.Size256, // with size 256x256
	})

	fmt.Println(resp.Data)
}
```

For all of response data please read more in file [struct_response.go](./response.go)
