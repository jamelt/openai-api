// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// OpenAIFileObject - The object type, which is always `file`.
type OpenAIFileObject string

const (
	OpenAIFileObjectFile OpenAIFileObject = "file"
)

func (e OpenAIFileObject) ToPointer() *OpenAIFileObject {
	return &e
}
func (e *OpenAIFileObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "file":
		*e = OpenAIFileObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OpenAIFileObject: %v", v)
	}
}

// Purpose - The intended purpose of the file. Supported values are `assistants`, `assistants_output`, `batch`, `batch_output`, `fine-tune`, `fine-tune-results` and `vision`.
type Purpose string

const (
	PurposeAssistants       Purpose = "assistants"
	PurposeAssistantsOutput Purpose = "assistants_output"
	PurposeBatch            Purpose = "batch"
	PurposeBatchOutput      Purpose = "batch_output"
	PurposeFineTune         Purpose = "fine-tune"
	PurposeFineTuneResults  Purpose = "fine-tune-results"
	PurposeVision           Purpose = "vision"
)

func (e Purpose) ToPointer() *Purpose {
	return &e
}
func (e *Purpose) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "assistants":
		fallthrough
	case "assistants_output":
		fallthrough
	case "batch":
		fallthrough
	case "batch_output":
		fallthrough
	case "fine-tune":
		fallthrough
	case "fine-tune-results":
		fallthrough
	case "vision":
		*e = Purpose(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Purpose: %v", v)
	}
}

// OpenAIFileStatus - Deprecated. The current status of the file, which can be either `uploaded`, `processed`, or `error`.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type OpenAIFileStatus string

const (
	OpenAIFileStatusUploaded  OpenAIFileStatus = "uploaded"
	OpenAIFileStatusProcessed OpenAIFileStatus = "processed"
	OpenAIFileStatusError     OpenAIFileStatus = "error"
)

func (e OpenAIFileStatus) ToPointer() *OpenAIFileStatus {
	return &e
}
func (e *OpenAIFileStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "uploaded":
		fallthrough
	case "processed":
		fallthrough
	case "error":
		*e = OpenAIFileStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OpenAIFileStatus: %v", v)
	}
}

// OpenAIFile - The `File` object represents a document that has been uploaded to OpenAI.
type OpenAIFile struct {
	// The file identifier, which can be referenced in the API endpoints.
	ID string `json:"id"`
	// The size of the file, in bytes.
	Bytes int64 `json:"bytes"`
	// The Unix timestamp (in seconds) for when the file was created.
	CreatedAt int64 `json:"created_at"`
	// The name of the file.
	Filename string `json:"filename"`
	// The object type, which is always `file`.
	Object OpenAIFileObject `json:"object"`
	// The intended purpose of the file. Supported values are `assistants`, `assistants_output`, `batch`, `batch_output`, `fine-tune`, `fine-tune-results` and `vision`.
	Purpose Purpose `json:"purpose"`
	// Deprecated. The current status of the file, which can be either `uploaded`, `processed`, or `error`.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Status OpenAIFileStatus `json:"status"`
	// Deprecated. For details on why a fine-tuning training file failed validation, see the `error` field on `fine_tuning.job`.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	StatusDetails *string `json:"status_details,omitempty"`
}

func (o *OpenAIFile) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OpenAIFile) GetBytes() int64 {
	if o == nil {
		return 0
	}
	return o.Bytes
}

func (o *OpenAIFile) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *OpenAIFile) GetFilename() string {
	if o == nil {
		return ""
	}
	return o.Filename
}

func (o *OpenAIFile) GetObject() OpenAIFileObject {
	if o == nil {
		return OpenAIFileObject("")
	}
	return o.Object
}

func (o *OpenAIFile) GetPurpose() Purpose {
	if o == nil {
		return Purpose("")
	}
	return o.Purpose
}

func (o *OpenAIFile) GetStatus() OpenAIFileStatus {
	if o == nil {
		return OpenAIFileStatus("")
	}
	return o.Status
}

func (o *OpenAIFile) GetStatusDetails() *string {
	if o == nil {
		return nil
	}
	return o.StatusDetails
}
