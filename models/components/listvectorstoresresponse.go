// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListVectorStoresResponse struct {
	Object  string              `json:"object"`
	Data    []VectorStoreObject `json:"data"`
	FirstID string              `json:"first_id"`
	LastID  string              `json:"last_id"`
	HasMore bool                `json:"has_more"`
}

func (o *ListVectorStoresResponse) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}

func (o *ListVectorStoresResponse) GetData() []VectorStoreObject {
	if o == nil {
		return []VectorStoreObject{}
	}
	return o.Data
}

func (o *ListVectorStoresResponse) GetFirstID() string {
	if o == nil {
		return ""
	}
	return o.FirstID
}

func (o *ListVectorStoresResponse) GetLastID() string {
	if o == nil {
		return ""
	}
	return o.LastID
}

func (o *ListVectorStoresResponse) GetHasMore() bool {
	if o == nil {
		return false
	}
	return o.HasMore
}
