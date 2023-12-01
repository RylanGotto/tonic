package openai

import (
	"fmt"
	"log"
)

type Text struct {
	Value       string        `json:"value"`
	Annotations []interface{} `json:"annotations"`
}

type Content struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type Message struct {
	ID            string      `json:"id"`
	Object        string      `json:"object"`
	Created_at    int         `json:"created_at"`
	Thread_id     string      `json:"thread_id"`
	Role          string      `json:"role"`
	Content       []Content   `json:"content"`
	File_ids      []string    `json:"file_ids"`
	Assistant_ids []string    `json:"assistant_ids"`
	Run_id        string      `json:"run_id"`
	Metadata      interface{} `json:"metadata"`
}

type MessageList struct {
	Object   string    `json:"object"`
	Data     []Message `json:"data"`
	First_id string    `json:"first_id"`
	Last_id  string    `json:"last_id"`
	Has_more bool      `json:"has_more"`
}

func (c Client) CreateMessage(thread_id string, content string) *Message {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	p := M{
		"role":    "user",
		"content": content,
	}

	r := Request{
		Func:    "CreateMessage",
		Type:    "POST",
		Url:     fmt.Sprintf(CreateMessageUrl, thread_id),
		Headers: h,
		Payload: p,
	}

	resp, err := c.DispatchRequest(r, &Message{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*Message)
}

func (c Client) RetrieveMessage(thread_id string, message_id string) *Message {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	r := Request{
		Func:    "RetrieveMessage",
		Type:    "GET",
		Url:     fmt.Sprintf(RetrieveMessageUrl, thread_id, message_id),
		Headers: h,
	}

	resp, err := c.DispatchRequest(r, &Message{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*Message)
}

func (c Client) ListMessages(thread_id string) *MessageList {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	r := Request{
		Func:    "ListMessages",
		Type:    "GET",
		Url:     fmt.Sprintf(ListMessagesUrl, thread_id),
		Headers: h,
	}

	resp, err := c.DispatchRequest(r, &MessageList{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*MessageList)
}
