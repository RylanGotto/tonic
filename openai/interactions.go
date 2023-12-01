package openai

type AssistantInteraction struct {
	OpenAiCall OpenAiCall
	Thread     *Thread
}

func NewAssistant(o OpenAI, thread *Thread) AssistantInteraction {
	return AssistantInteraction{
		OpenAiCall: o.OpenAiCall,
		Thread:     thread,
	}
}

func (ai *AssistantInteraction) MessageAssistant(content string) interface{} {
	ai.OpenAiCall.CreateMessage(ai.Thread.ID, content)
	r := ai.OpenAiCall.CreateRun(AssistantID, ai.Thread.ID)
	rr := ai.OpenAiCall.RetrieveRun(ai.Thread.ID, r.ID)
	status := rr.Status
	for status != "completed" {
		steps := ai.OpenAiCall.ListRunSteps(ai.Thread.ID, r.ID)
		if len(steps.Data) == 0 {
			continue
		}
		status = steps.Data[0].Status
		rt := steps.Data[0].Type
		if status == "completed" && rt == "message_creation" {
			return ai.OpenAiCall.ListMessages(ai.Thread.ID).Data[0].Content[0].Text.Value
		}
		if status == "in_progress" && rt == "tool_calls" {
			ai.HandleToolCalls(steps)
		}
	}

	return nil
}

func (ai *AssistantInteraction) HandleToolCalls(steps *ListRunStep) {
	_ = steps
}
