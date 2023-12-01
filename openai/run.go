package openai

import (
	"fmt"
	"log"
)

type Runs struct {
	ID           string        `json:"id"`
	Object       string        `json:"object"`
	Created_at   int           `json:"created_at"`
	Assistant_id string        `json:"assistant_id"`
	Thread_id    string        `json:"thread_id"`
	Status       string        `json:"status"`
	Started_at   int           `json:"started_at"`
	Expires_at   int           `json:"expires_at"`
	Cancelled_at int           `json:"cancelled_at"`
	Failed_at    int           `json:"failed_at"`
	Completed_at int           `json:"completed_at"`
	Last_error   string        `json:"last_error"`
	Model        string        `json:"model"`
	Instructions string        `json:"instructions"`
	Tools        []interface{} `json:"tools"`
	File_ids     []string      `json:"file_ids"`
	Metadata     interface{}   `json:"metadata"`
}

type MessageCreation struct {
	Message_id string `json:"message_id"`
}

type StepDetails struct {
	Type             string          `json:"type"`
	Message_creation MessageCreation `json:"message_creation"`
	Tool_calls       []struct {
		ID       string `json:"id"`
		Type     string `json:"type"`
		Function struct {
			Name      string      `json:"name"`
			Arguments interface{} `json:"arguments"`
		}
	} `json:"tool_calls"`
}

type RunStep struct {
	ID           string      `json:"id"`
	Object       string      `json:"object"`
	Created_at   int         `json:"created_at"`
	Assistant_id string      `json:"assistant_id"`
	Thread_id    string      `json:"thread_id"`
	Status       string      `json:"status"`
	Model        string      `json:"model"`
	Type         string      `json:"type"`
	Cancelled_at int         `json:"cancelled_at"`
	Completed_at int         `json:"completed_at"`
	Expired_at   int         `json:"expired_at"`
	Failed_at    int         `json:"failed_at"`
	Last_error   string      `json:"last_error"`
	Step_details StepDetails `json:"step_details"`
}

type ListRunStep struct {
	Object   string    `json:"object"`
	Data     []RunStep `json:"data"`
	First_id string    `json:"first_id"`
	Last_id  string    `json:"last_id"`
	Has_more bool      `json:"has_more"`
}

func (c Client) CreateRun(assistant_id string, thread_id string) *Runs {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	p := M{
		"assistant_id": assistant_id,
	}

	r := Request{
		Func:    "CreateRun",
		Type:    "POST",
		Url:     fmt.Sprintf(CreateRunUrl, thread_id),
		Headers: h,
		Payload: p,
	}

	resp, err := c.DispatchRequest(r, &Runs{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*Runs)
}

func (c Client) RetrieveRun(thread_id string, run_id string) *Runs {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	r := Request{
		Func:    "RetrieveRun",
		Type:    "GET",
		Url:     fmt.Sprintf(RetrieveRunUrl, thread_id, run_id),
		Headers: h,
	}

	resp, err := c.DispatchRequest(r, &Runs{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*Runs)
}

func (c Client) ListRuns(thread_id string) *Runs {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	r := Request{
		Func:    "ListRuns",
		Type:    "GET",
		Url:     fmt.Sprintf(ListRunsUrl, thread_id),
		Headers: h,
	}

	resp, err := c.DispatchRequest(r, &Runs{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*Runs)
}

func (c Client) CancelRun(thread_id string, run_id string) *Runs {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	r := Request{
		Func:    "CancelRun",
		Type:    "POST",
		Url:     fmt.Sprintf(CancelRunUrl, thread_id, run_id),
		Headers: h,
	}

	resp, err := c.DispatchRequest(r, &Runs{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*Runs)
}

func (c Client) SubmitToolOutput(thread_id string, run_id string, tool_outputs []interface{}) *Runs {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	p := M{
		"tool_outputs": tool_outputs,
	}

	r := Request{
		Func:    "SubmitToolOutput",
		Type:    "POST",
		Url:     fmt.Sprintf(SubmitToolOutputUrl, thread_id, run_id),
		Headers: h,
		Payload: p,
	}

	resp, err := c.DispatchRequest(r, &Runs{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*Runs)
}

func (c Client) CreateThreadAndRun(assistant_id string, content string) *Runs {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	var messages []M
	userm := M{"role": "user", "content": content}

	messages = append(messages, userm)

	p := M{
		"assistant_id": assistant_id,
		"thread":       M{"messages": messages},
	}

	r := Request{
		Func:    "CreateThreadAndRun",
		Type:    "POST",
		Url:     CreateThreadAndRunUrl,
		Headers: h,
		Payload: p,
	}

	resp, err := c.DispatchRequest(r, &Runs{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*Runs)
}

func (c Client) ListRunSteps(thread_id string, run_id string) *ListRunStep {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	r := Request{
		Func:    "ListRunSteps",
		Type:    "GET",
		Url:     fmt.Sprintf(ListRunStepsUrl, thread_id, run_id),
		Headers: h,
	}

	resp, err := c.DispatchRequest(r, &ListRunStep{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*ListRunStep)
}

func (c Client) RetrieveRunStep(thread_id string, run_id string, step_id string) *RunStep {
	h := Headers()
	h["OpenAI-Beta"] = AssistantHeader

	r := Request{
		Func:    "RetrieveRunStep",
		Type:    "GET",
		Url:     fmt.Sprintf(RetrieveRunStepUrl, thread_id, run_id, step_id),
		Headers: h,
	}

	resp, err := c.DispatchRequest(r, &RunStep{})

	if err != nil {
		log.Fatalln(err)
	}

	return resp.(*RunStep)
}
