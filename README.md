<div align="center">
    <img src="./assets/logo.png">
    <h2>go-chatgpt-sdk</h2>
    <p>This library provides unofficial Go client SDK for OpenAI API</p>
</div>

### Install

```bash
go get -u github.com/ak9024/go-chatgpt-sdk
```

### Model Support

- [x] Chat
- [x] Text
- [x] Image
- [ ] Modeerations
- [ ] Audio

### Prerequisite

- [Go](https://go.dev/doc/install)
- [OpenAI Key](https://platform.openai.com/account/api-keys)

### Usage

> to usage with model chat, please usage `c.ChatCompletions`

```go
import (
	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

func main() {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: "",
	})

	resp, _ := c.ChatCompletions(gochatgptsdk.ModelChat{
		Model: "gpt-3.5-turbo-16k",
		Messages: []gochatgptsdk.Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant",
			},
			{
				Role:    "user",
				Content: "How the weather today?",
			},
		},
	})
}
```

### Usage with model text

> To usage with model text, please usage `c.Completions()`

```go
package main

import (
	"fmt"
	"os"

	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

func main() {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: os.Getenv("OPENAI_API_KEY"),
	})

	resp, _ := c.Completions(gochatgptsdk.ModelText{
		Model:     "text-davinci-003",
		Prompt:    "Please info what the weather today in Kota Palu?",
		MaxTokens: 100, // max generates of word
	})

	fmt.Println(resp.Choices)
}
```

### Usage with model images

> Create images, please use `c.ImagesGenerations`

```go
package main

import (
	"fmt"
	"os"

	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

func main() {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: os.Getenv("OPENAI_API_KEY"),
	})

	resp, _ := c.ImagesGenerations(gochatgptsdk.ModelImages{
		Prompt: "Sunset sketch art",
		N:      5, // generates 5 images
		Size:   gochatgptsdk.Size512, // with size 512x512
	})

	fmt.Println(resp.Data)
}
```

> Create images variations of a given image, please use `c.ImagesVariations()`

```go
package main

import (
	"fmt"
	"os"

	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

func main() {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: os.Getenv("OPENAI_API_KEY"),
	})

	resp, _ := c.ImagesVariations(gochatgptsdk.ModelImagesVariations{
		Image: "./path/to/example-img.png", // please suitable with your path image
		N:     "2",                  // generate 2 images
		Size:  gochatgptsdk.Size256, // with size 256x256
	})

	fmt.Println(resp.Data)
}
```
