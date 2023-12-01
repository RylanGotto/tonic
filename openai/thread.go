package openai

import "log"

type Thread struct {
	ID         string      `json:"id"`
	Object     string      `json:"object"`
	Created_at int         `json:"create_at"`
	Metadata   interface{} `json:"metadata"`
}

func (c Client) CreateThread() *Thread {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	r := Request{
		Func:    "CreateThread",
		Type:    "POST",
		Url:     ThreadsUrl,
		Headers: h,
	}

	resp, err := c.DispatchRequest(r, &Thread{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*Thread)
}
