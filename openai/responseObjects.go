package openai

// Thread Object

type Thread2 struct {
	ID         string      `json:"id"`
	Object     string      `json:"object"`
	Created_at int         `json:"create_at"`
	Metadata   interface{} `json:"metadata"`
}
