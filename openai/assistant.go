package openai

import (
	"log"
	"net/http"
)

// Assistant object

type Tools struct {
	Type string
}

type Assistant struct {
	ID           string      `json:"id"`
	Object       string      `json:"object"`
	Created_at   int         `json:"created_at"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Model        string      `json:"model"`
	Instructions string      `json:"instructions"`
	Tools        []Tools     `json:"tools"`
	File_ids     []string    `json:"file_ids"`
	Metadata     interface{} `json:"metadata"`
}

type AssistantPayload struct {
	Instructions string `json:"instructions"`
	Name         string `json:"name"`
	Tools        []M    `json:"tools"`
	Model        string `json:"model"`
}

func (c Client) CreateAssistant() *Assistant {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	var tools []M

	t := M{"type": "code_interpreter"}

	tools = append(tools, t)

	p := AssistantPayload{
		Instructions: "You are a helpful assistant",
		Name:         "George",
		Tools:        tools,
		Model:        "gpt-4",
	}

	r := Request{
		Func:    "CreateAssistant",
		Type:    http.MethodPost,
		Url:     AssistantUrl,
		Headers: h,
		Payload: p,
	}

	resp, err := c.DispatchRequest(r, &Assistant{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*Assistant)
}
