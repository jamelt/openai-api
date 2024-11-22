// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type ChatCompletionModalities string

const (
	ChatCompletionModalitiesText  ChatCompletionModalities = "text"
	ChatCompletionModalitiesAudio ChatCompletionModalities = "audio"
)

func (e ChatCompletionModalities) ToPointer() *ChatCompletionModalities {
	return &e
}
func (e *ChatCompletionModalities) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "audio":
		*e = ChatCompletionModalities(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChatCompletionModalities: %v", v)
	}
}