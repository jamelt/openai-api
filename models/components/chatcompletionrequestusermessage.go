// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

type ChatCompletionRequestUserMessageContentType string

const (
	ChatCompletionRequestUserMessageContentTypeStr                                                ChatCompletionRequestUserMessageContentType = "str"
	ChatCompletionRequestUserMessageContentTypeArrayOfChatCompletionRequestUserMessageContentPart ChatCompletionRequestUserMessageContentType = "arrayOfChatCompletionRequestUserMessageContentPart"
)

// ChatCompletionRequestUserMessageContent - The contents of the user message.
type ChatCompletionRequestUserMessageContent struct {
	Str                                                *string
	ArrayOfChatCompletionRequestUserMessageContentPart []ChatCompletionRequestUserMessageContentPart

	Type ChatCompletionRequestUserMessageContentType
}

func CreateChatCompletionRequestUserMessageContentStr(str string) ChatCompletionRequestUserMessageContent {
	typ := ChatCompletionRequestUserMessageContentTypeStr

	return ChatCompletionRequestUserMessageContent{
		Str:  &str,
		Type: typ,
	}
}

func CreateChatCompletionRequestUserMessageContentArrayOfChatCompletionRequestUserMessageContentPart(arrayOfChatCompletionRequestUserMessageContentPart []ChatCompletionRequestUserMessageContentPart) ChatCompletionRequestUserMessageContent {
	typ := ChatCompletionRequestUserMessageContentTypeArrayOfChatCompletionRequestUserMessageContentPart

	return ChatCompletionRequestUserMessageContent{
		ArrayOfChatCompletionRequestUserMessageContentPart: arrayOfChatCompletionRequestUserMessageContentPart,
		Type: typ,
	}
}

func (u *ChatCompletionRequestUserMessageContent) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ChatCompletionRequestUserMessageContentTypeStr
		return nil
	}

	var arrayOfChatCompletionRequestUserMessageContentPart []ChatCompletionRequestUserMessageContentPart = []ChatCompletionRequestUserMessageContentPart{}
	if err := utils.UnmarshalJSON(data, &arrayOfChatCompletionRequestUserMessageContentPart, "", true, true); err == nil {
		u.ArrayOfChatCompletionRequestUserMessageContentPart = arrayOfChatCompletionRequestUserMessageContentPart
		u.Type = ChatCompletionRequestUserMessageContentTypeArrayOfChatCompletionRequestUserMessageContentPart
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ChatCompletionRequestUserMessageContent", string(data))
}

func (u ChatCompletionRequestUserMessageContent) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfChatCompletionRequestUserMessageContentPart != nil {
		return utils.MarshalJSON(u.ArrayOfChatCompletionRequestUserMessageContentPart, "", true)
	}

	return nil, errors.New("could not marshal union type ChatCompletionRequestUserMessageContent: all fields are null")
}

// ChatCompletionRequestUserMessageRole - The role of the messages author, in this case `user`.
type ChatCompletionRequestUserMessageRole string

const (
	ChatCompletionRequestUserMessageRoleUser ChatCompletionRequestUserMessageRole = "user"
)

func (e ChatCompletionRequestUserMessageRole) ToPointer() *ChatCompletionRequestUserMessageRole {
	return &e
}
func (e *ChatCompletionRequestUserMessageRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "user":
		*e = ChatCompletionRequestUserMessageRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChatCompletionRequestUserMessageRole: %v", v)
	}
}

type ChatCompletionRequestUserMessage struct {
	// The contents of the user message.
	//
	Content ChatCompletionRequestUserMessageContent `json:"content"`
	// The role of the messages author, in this case `user`.
	Role ChatCompletionRequestUserMessageRole `json:"role"`
	// An optional name for the participant. Provides the model information to differentiate between participants of the same role.
	Name *string `json:"name,omitempty"`
}

func (o *ChatCompletionRequestUserMessage) GetContent() ChatCompletionRequestUserMessageContent {
	if o == nil {
		return ChatCompletionRequestUserMessageContent{}
	}
	return o.Content
}

func (o *ChatCompletionRequestUserMessage) GetRole() ChatCompletionRequestUserMessageRole {
	if o == nil {
		return ChatCompletionRequestUserMessageRole("")
	}
	return o.Role
}

func (o *ChatCompletionRequestUserMessage) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}