package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Client struct {
	C *http.Client
}

type OpenAI struct {
	OpenAiCall interface {
		ChatCompletion(um string) *ChatCompletionRes
		CreateAssistant() *AssistantRes
	}
}

type Response struct {
	Response string
}

type M map[string]interface{}

type Request struct {
	Func    string
	Type    string
	Url     string
	Headers M
	Payload interface{}
}

type RequestError struct {
	Err error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func Error() error {
	return &RequestError{Err: errors.New("error creating request")}
}

func InitClient() OpenAI {
	c := Client{C: &http.Client{}}
	return OpenAI{
		OpenAiCall: c,
	}
}

func createRequest(r Request) (*http.Request, error) {
	if r.Type == "POST" {

		j, err := json.Marshal(r.Payload)

		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest(r.Type, r.Url, bytes.NewBuffer(j))
		for i, k := range r.Headers {
			req.Header.Set(i, k.(string))
		}
		if err != nil {
			log.Fatalln(err)
		}
		return req, nil
	}

	return nil, Error()
}

func Headers() M {
	return M{"Authorization": OpenAiToken, "Content-Type": ContentType}
}

func (c Client) DispatchRequest(r Request, t interface{}) (interface{}, error) {

	req, err := createRequest(r)

	if err != nil {
		return nil, err
	}

	resp, err := c.C.Do(req)

	if err != nil {
		return nil, err
	}

	if err != nil {
		log.Fatalln(err)
	}

	derr := json.NewDecoder(resp.Body).Decode(t)
	defer resp.Body.Close()

	if derr != nil {
		log.Fatalln(derr)
	}

	return t, nil
}
