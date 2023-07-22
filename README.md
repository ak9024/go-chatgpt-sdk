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
- [ ] Moderation
- [x] Image
- [ ] Audio
- [ ] Other

### Prerequisite

- [Go](https://go.dev/doc/install)
- [OpenAI Key](https://platform.openai.com/account/api-keys)

### Usage

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

```go
func main() {
	resp, _ := c.Completions(gochatgptsdk.ModelText{
		Model:  "text-davinci-003",
		Prompt: "What the weather today?",
	})
}
```

### Model Options

|                MODEL                |    TPM     |  RPM  |
| :---------------------------------: | :--------: | :---: |
|                CHAT                 |            |       |
|            gpt-3.5-turbo            |   90,000   | 3,500 |
|         gpt-3.5-turbo-0301          |   90,000   | 3,500 |
|         gpt-3.5-turbo-0613          |   90,000   | 3,500 |
|          gpt-3.5-turbo-16k          |  180,000   | 3,500 |
|       gpt-3.5-turbo-16k-0613        |  180,000   | 3,500 |
|                TEXT                 |            |       |
|                 ada                 |  250,000   | 3,000 |
|        ada-code-search-code         |  250,000   | 3,000 |
|        ada-code-search-text         |  250,000   | 3,000 |
|         ada-search-document         |  250,000   | 3,000 |
|          ada-search-query           |  250,000   | 3,000 |
|           ada-similarity            |  250,000   | 3,000 |
|               babbage               |  250,000   | 3,000 |
|      babbage-code-search-code       |  250,000   | 3,000 |
|      babbage-code-search-text       |  250,000   | 3,000 |
|       babbage-search-document       |  250,000   | 3,000 |
|        babbage-search-query         |  250,000   | 3,000 |
|         babbage-similarity          |  250,000   | 3,000 |
|        code-davinci-edit-001        |  150,000   |  20   |
|      code-search-ada-code-001       |  250,000   | 3,000 |
|      code-search-ada-text-001       |  250,000   | 3,000 |
|    code-search-babbage-code-001     |  250,000   | 3,000 |
|    code-search-babbage-text-001     |  250,000   | 3,000 |
|                curie                |  250,000   | 3,000 |
|         curie-instruct-beta         |  250,000   | 3,000 |
|        curie-search-document        |  250,000   | 3,000 |
|         curie-search-query          |  250,000   | 3,000 |
|          curie-similarity           |  250,000   | 3,000 |
|               davinci               |  250,000   | 3,000 |
|        davinci-instruct-beta        |  250,000   | 3,000 |
|       davinci-search-document       |  250,000   | 3,000 |
|        davinci-search-query         |  250,000   | 3,000 |
|         davinci-similarity          |  250,000   | 3,000 |
|            text-ada-001             |  250,000   | 3,000 |
|          text-babbage-001           |  250,000   | 3,000 |
|           text-curie-001            |  250,000   | 3,000 |
|          text-davinci-001           |  250,000   | 3,000 |
|          text-davinci-002           |  250,000   | 3,000 |
|          text-davinci-003           |  250,000   | 3,000 |
|        text-davinci-edit-001        |  150,000   |  20   |
|       text-embedding-ada-002        | 1,000,000  | 3,000 |
|       text-search-ada-doc-001       |  250,000   | 3,000 |
|      text-search-ada-query-001      |  250,000   | 3,000 |
|     text-search-babbage-doc-001     |  250,000   | 3,000 |
|    text-search-babbage-query-001    |  250,000   | 3,000 |
|      text-search-curie-doc-001      |  250,000   | 3,000 |
|     text-search-curie-query-001     |  250,000   | 3,000 |
|     text-search-davinci-doc-001     |  250,000   | 3,000 |
|    text-search-davinci-query-001    |  250,000   | 3,000 |
|       text-similarity-ada-001       |  250,000   | 3,000 |
|     text-similarity-babbage-001     |  250,000   | 3,000 |
|      text-similarity-curie-001      |  250,000   | 3,000 |
|     text-similarity-davinci-001     |  250,000   | 3,000 |
|             MODERATION              |            |       |
|       text-moderation-latest        |   1,250    | 1,000 |
|       text-moderation-stable        |   1,250    | 1,000 |
|                IMAGE                | IMG / MIN  |       |
|              DALL·E 2               |     50     |   ∞   |
|                AUDIO                |            |       |
|              whisper-1              | 25,000,000 |  50   |
|                OTHER                |            |       |
| Default limits for all other models |  250,000   | 3,000 |
