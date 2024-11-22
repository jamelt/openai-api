// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/jamelt/openai-api/models/components"
)

type SubmitToolOuputsToRunRequest struct {
	// The ID of the [thread](/docs/api-reference/threads) to which this run belongs.
	ThreadID string `pathParam:"style=simple,explode=false,name=thread_id"`
	// The ID of the run that requires the tool output submission.
	RunID                       string                                 `pathParam:"style=simple,explode=false,name=run_id"`
	SubmitToolOutputsRunRequest components.SubmitToolOutputsRunRequest `request:"mediaType=application/json"`
}

func (o *SubmitToolOuputsToRunRequest) GetThreadID() string {
	if o == nil {
		return ""
	}
	return o.ThreadID
}

func (o *SubmitToolOuputsToRunRequest) GetRunID() string {
	if o == nil {
		return ""
	}
	return o.RunID
}

func (o *SubmitToolOuputsToRunRequest) GetSubmitToolOutputsRunRequest() components.SubmitToolOutputsRunRequest {
	if o == nil {
		return components.SubmitToolOutputsRunRequest{}
	}
	return o.SubmitToolOutputsRunRequest
}

type SubmitToolOuputsToRunResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	RunObject *components.RunObject
}

func (o *SubmitToolOuputsToRunResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SubmitToolOuputsToRunResponse) GetRunObject() *components.RunObject {
	if o == nil {
		return nil
	}
	return o.RunObject
}