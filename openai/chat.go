package openai

import (
	"log"
	"net/http"
)

// Chat Object
type message struct {
	Role    string
	Content string
}

type Usage struct {
	Prompt_tokens     int
	Completion_tokens int
	Total_tokens      int
}

type Choices struct {
	Index        int     `json:"index"`
	Message      message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type ChatCompletionRes struct {
	Id                 string    `json:"id"`
	Object             string    `json:"object"`
	Created            int       `json:"created"`
	Model              string    `json:"model"`
	System_fingerprint string    `json:"system_fingerprint"`
	Choices            []Choices `json:"choices"`
	Usage              Usage     `json:"usage"`
}

type ChatCompletionPayload struct {
	Model    string `json:"model"`
	Messages []M    `json:"messages"`
}

func (c Client) ChatCompletion(um string) *ChatCompletionRes {
	h := Headers()

	var messages []M
	sysm := M{"role": "system", "content": SystemMessage}
	userm := M{"role": "user", "content": um}

	messages = append(messages, sysm, userm)

	p := ChatCompletionPayload{
		Model:    Model,
		Messages: messages,
	}

	r := Request{
		Func:    "ChatCompletion",
		Type:    http.MethodPost,
		Url:     ChatCompletionUrl,
		Headers: h,
		Payload: p,
	}

	resp, err := c.DispatchRequest(r, &ChatCompletionRes{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*ChatCompletionRes)
}
