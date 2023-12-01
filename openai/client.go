package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Client struct {
	C *http.Client
}

type OpenAiCall interface {
	ChatCompletion(um string) *ChatCompletion
	CreateAssistant() *Assistant
	CreateThread() *Thread
	CreateMessage(thread_id string, content string) *Message
	RetrieveMessage(thread_id string, message_id string) *Message
	ListMessages(thread_id string) *MessageList
	CreateRun(assistant_id string, thread_id string) *Runs
	CreateThreadAndRun(assistant_id string, content string) *Runs
	RetrieveRun(thread_id string, run_id string) *Runs
	ListRuns(thread_id string) *Runs
	CancelRun(thread_id string, run_id string) *Runs
	SubmitToolOutput(thread_id string, run_id string, tool_outputs []interface{}) *Runs
	RetrieveRunStep(thread_id string, run_id string, step_id string) *RunStep
	ListRunSteps(thread_id string, run_id string) *ListRunStep
}

type OpenAI struct {
	OpenAiCall OpenAiCall
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

func New() OpenAI {
	c := Client{C: &http.Client{}}
	return OpenAI{
		OpenAiCall: c,
	}
}

func createRequest(r Request) (*http.Request, error) {
	if r.Type == "GET" {
		req, err := http.NewRequest(r.Type, r.Url, nil)
		for i, k := range r.Headers {
			req.Header.Set(i, k.(string))
		}
		if err != nil {
			log.Fatalln(err)
		}
		return req, nil
	}
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
	return M{"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")), "Content-Type": ContentType}
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

	derr := json.NewDecoder(resp.Body).Decode(t)
	defer resp.Body.Close()

	if derr != nil {
		log.Fatalln(derr)
	}

	return t, nil
}
