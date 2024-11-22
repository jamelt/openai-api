// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// CreateCompletionResponseFinishReason - The reason the model stopped generating tokens. This will be `stop` if the model hit a natural stop point or a provided stop sequence,
// `length` if the maximum number of tokens specified in the request was reached,
// or `content_filter` if content was omitted due to a flag from our content filters.
type CreateCompletionResponseFinishReason string

const (
	CreateCompletionResponseFinishReasonStop          CreateCompletionResponseFinishReason = "stop"
	CreateCompletionResponseFinishReasonLength        CreateCompletionResponseFinishReason = "length"
	CreateCompletionResponseFinishReasonContentFilter CreateCompletionResponseFinishReason = "content_filter"
)

func (e CreateCompletionResponseFinishReason) ToPointer() *CreateCompletionResponseFinishReason {
	return &e
}
func (e *CreateCompletionResponseFinishReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stop":
		fallthrough
	case "length":
		fallthrough
	case "content_filter":
		*e = CreateCompletionResponseFinishReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCompletionResponseFinishReason: %v", v)
	}
}

type CreateCompletionResponseLogprobs struct {
	TextOffset    []int64              `json:"text_offset,omitempty"`
	TokenLogprobs []float64            `json:"token_logprobs,omitempty"`
	Tokens        []string             `json:"tokens,omitempty"`
	TopLogprobs   []map[string]float64 `json:"top_logprobs,omitempty"`
}

func (o *CreateCompletionResponseLogprobs) GetTextOffset() []int64 {
	if o == nil {
		return nil
	}
	return o.TextOffset
}

func (o *CreateCompletionResponseLogprobs) GetTokenLogprobs() []float64 {
	if o == nil {
		return nil
	}
	return o.TokenLogprobs
}

func (o *CreateCompletionResponseLogprobs) GetTokens() []string {
	if o == nil {
		return nil
	}
	return o.Tokens
}

func (o *CreateCompletionResponseLogprobs) GetTopLogprobs() []map[string]float64 {
	if o == nil {
		return nil
	}
	return o.TopLogprobs
}

type CreateCompletionResponseChoices struct {
	// The reason the model stopped generating tokens. This will be `stop` if the model hit a natural stop point or a provided stop sequence,
	// `length` if the maximum number of tokens specified in the request was reached,
	// or `content_filter` if content was omitted due to a flag from our content filters.
	//
	FinishReason CreateCompletionResponseFinishReason `json:"finish_reason"`
	Index        int64                                `json:"index"`
	Logprobs     *CreateCompletionResponseLogprobs    `json:"logprobs"`
	Text         string                               `json:"text"`
}

func (o *CreateCompletionResponseChoices) GetFinishReason() CreateCompletionResponseFinishReason {
	if o == nil {
		return CreateCompletionResponseFinishReason("")
	}
	return o.FinishReason
}

func (o *CreateCompletionResponseChoices) GetIndex() int64 {
	if o == nil {
		return 0
	}
	return o.Index
}

func (o *CreateCompletionResponseChoices) GetLogprobs() *CreateCompletionResponseLogprobs {
	if o == nil {
		return nil
	}
	return o.Logprobs
}

func (o *CreateCompletionResponseChoices) GetText() string {
	if o == nil {
		return ""
	}
	return o.Text
}

// CreateCompletionResponseObject - The object type, which is always "text_completion"
type CreateCompletionResponseObject string

const (
	CreateCompletionResponseObjectTextCompletion CreateCompletionResponseObject = "text_completion"
)

func (e CreateCompletionResponseObject) ToPointer() *CreateCompletionResponseObject {
	return &e
}
func (e *CreateCompletionResponseObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text_completion":
		*e = CreateCompletionResponseObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCompletionResponseObject: %v", v)
	}
}

// CreateCompletionResponse - Represents a completion response from the API. Note: both the streamed and non-streamed response objects share the same shape (unlike the chat endpoint).
type CreateCompletionResponse struct {
	// A unique identifier for the completion.
	ID string `json:"id"`
	// The list of completion choices the model generated for the input prompt.
	Choices []CreateCompletionResponseChoices `json:"choices"`
	// The Unix timestamp (in seconds) of when the completion was created.
	Created int64 `json:"created"`
	// The model used for completion.
	Model string `json:"model"`
	// This fingerprint represents the backend configuration that the model runs with.
	//
	// Can be used in conjunction with the `seed` request parameter to understand when backend changes have been made that might impact determinism.
	//
	SystemFingerprint *string `json:"system_fingerprint,omitempty"`
	// The object type, which is always "text_completion"
	Object CreateCompletionResponseObject `json:"object"`
	// Usage statistics for the completion request.
	Usage *CompletionUsage `json:"usage,omitempty"`
}

func (o *CreateCompletionResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateCompletionResponse) GetChoices() []CreateCompletionResponseChoices {
	if o == nil {
		return []CreateCompletionResponseChoices{}
	}
	return o.Choices
}

func (o *CreateCompletionResponse) GetCreated() int64 {
	if o == nil {
		return 0
	}
	return o.Created
}

func (o *CreateCompletionResponse) GetModel() string {
	if o == nil {
		return ""
	}
	return o.Model
}

func (o *CreateCompletionResponse) GetSystemFingerprint() *string {
	if o == nil {
		return nil
	}
	return o.SystemFingerprint
}

func (o *CreateCompletionResponse) GetObject() CreateCompletionResponseObject {
	if o == nil {
		return CreateCompletionResponseObject("")
	}
	return o.Object
}

func (o *CreateCompletionResponse) GetUsage() *CompletionUsage {
	if o == nil {
		return nil
	}
	return o.Usage
}