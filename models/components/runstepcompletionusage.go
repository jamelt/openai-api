// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// RunStepCompletionUsage - Usage statistics related to the run step. This value will be `null` while the run step's status is `in_progress`.
type RunStepCompletionUsage struct {
	// Number of completion tokens used over the course of the run step.
	CompletionTokens int64 `json:"completion_tokens"`
	// Number of prompt tokens used over the course of the run step.
	PromptTokens int64 `json:"prompt_tokens"`
	// Total number of tokens used (prompt + completion).
	TotalTokens int64 `json:"total_tokens"`
}

func (o *RunStepCompletionUsage) GetCompletionTokens() int64 {
	if o == nil {
		return 0
	}
	return o.CompletionTokens
}

func (o *RunStepCompletionUsage) GetPromptTokens() int64 {
	if o == nil {
		return 0
	}
	return o.PromptTokens
}

func (o *RunStepCompletionUsage) GetTotalTokens() int64 {
	if o == nil {
		return 0
	}
	return o.TotalTokens
}
