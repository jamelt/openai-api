// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type DeleteVectorStoreFileResponseObject string

const (
	DeleteVectorStoreFileResponseObjectVectorStoreFileDeleted DeleteVectorStoreFileResponseObject = "vector_store.file.deleted"
)

func (e DeleteVectorStoreFileResponseObject) ToPointer() *DeleteVectorStoreFileResponseObject {
	return &e
}
func (e *DeleteVectorStoreFileResponseObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vector_store.file.deleted":
		*e = DeleteVectorStoreFileResponseObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteVectorStoreFileResponseObject: %v", v)
	}
}

type DeleteVectorStoreFileResponse struct {
	ID      string                              `json:"id"`
	Deleted bool                                `json:"deleted"`
	Object  DeleteVectorStoreFileResponseObject `json:"object"`
}

func (o *DeleteVectorStoreFileResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteVectorStoreFileResponse) GetDeleted() bool {
	if o == nil {
		return false
	}
	return o.Deleted
}

func (o *DeleteVectorStoreFileResponse) GetObject() DeleteVectorStoreFileResponseObject {
	if o == nil {
		return DeleteVectorStoreFileResponseObject("")
	}
	return o.Object
}
