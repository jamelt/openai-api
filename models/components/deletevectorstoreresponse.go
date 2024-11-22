// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type DeleteVectorStoreResponseObject string

const (
	DeleteVectorStoreResponseObjectVectorStoreDeleted DeleteVectorStoreResponseObject = "vector_store.deleted"
)

func (e DeleteVectorStoreResponseObject) ToPointer() *DeleteVectorStoreResponseObject {
	return &e
}
func (e *DeleteVectorStoreResponseObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vector_store.deleted":
		*e = DeleteVectorStoreResponseObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteVectorStoreResponseObject: %v", v)
	}
}

type DeleteVectorStoreResponse struct {
	ID      string                          `json:"id"`
	Deleted bool                            `json:"deleted"`
	Object  DeleteVectorStoreResponseObject `json:"object"`
}

func (o *DeleteVectorStoreResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteVectorStoreResponse) GetDeleted() bool {
	if o == nil {
		return false
	}
	return o.Deleted
}

func (o *DeleteVectorStoreResponse) GetObject() DeleteVectorStoreResponseObject {
	if o == nil {
		return DeleteVectorStoreResponseObject("")
	}
	return o.Object
}