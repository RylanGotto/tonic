package openai

const (
	Model                 = "gpt-3.5-turbo-16k"
	ContentType           = "application/json"
	SystemMessage         = "You are a helpful assistant."
	AssistantHeader       = "assistants=v1"
	AssistantID           = "asst_4E8As8KjgcMjb5Pe9WDfktU2"
	ChatCompletionUrl     = "https://api.openai.com/v1/chat/completions"
	AssistantUrl          = "https://api.openai.com/v1/assistants"
	ThreadsUrl            = "https://api.openai.com/v1/threads"
	CreateMessageUrl      = "https://api.openai.com/v1/threads/%s/messages"
	RetrieveMessageUrl    = "https://api.openai.com/v1/threads/%s/messages/%s"
	ListMessagesUrl       = "https://api.openai.com/v1/threads/%s/messages"
	CreateRunUrl          = "https://api.openai.com/v1/threads/%s/runs"
	RetrieveRunUrl        = "https://api.openai.com/v1/threads/%s/runs/%s"
	ListRunsUrl           = "https://api.openai.com/v1/threads/%s/runs"
	CancelRunUrl          = "https://api.openai.com/v1/threads/%s/runs/%s/cancel"
	SubmitToolOutputUrl   = "https://api.openai.com/v1/threads/%s/runs/%s/submit_tool_outputs"
	CreateThreadAndRunUrl = "https://api.openai.com/v1/threads/runs"
	RetrieveRunStepUrl    = "https://api.openai.com/v1/threads/%s/runs/%s/steps/%s"
	ListRunStepsUrl       = "https://api.openai.com/v1/threads/%s/runs/%s/steps"
)
