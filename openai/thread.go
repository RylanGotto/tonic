package openai

type Thread struct {
	ID         string      `json:"id"`
	Object     string      `json:"object"`
	Created_at int         `json:"create_at"`
	Metadata   interface{} `json:"metadata"`
}

// func (c Client) CreateThread() Thread {

// }
