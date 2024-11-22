// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AssistantToolsFunctionType - The type of tool being defined: `function`
type AssistantToolsFunctionType string

const (
	AssistantToolsFunctionTypeFunction AssistantToolsFunctionType = "function"
)

func (e AssistantToolsFunctionType) ToPointer() *AssistantToolsFunctionType {
	return &e
}
func (e *AssistantToolsFunctionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "function":
		*e = AssistantToolsFunctionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssistantToolsFunctionType: %v", v)
	}
}

type AssistantToolsFunction struct {
	// The type of tool being defined: `function`
	Type     AssistantToolsFunctionType `json:"type"`
	Function FunctionObject             `json:"function"`
}

func (o *AssistantToolsFunction) GetType() AssistantToolsFunctionType {
	if o == nil {
		return AssistantToolsFunctionType("")
	}
	return o.Type
}

func (o *AssistantToolsFunction) GetFunction() FunctionObject {
	if o == nil {
		return FunctionObject{}
	}
	return o.Function
}