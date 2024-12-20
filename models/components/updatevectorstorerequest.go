// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// UpdateVectorStoreRequestMetadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
type UpdateVectorStoreRequestMetadata struct {
}

type UpdateVectorStoreRequest struct {
	// The name of the vector store.
	Name *string `json:"name,omitempty"`
	// The expiration policy for a vector store.
	ExpiresAfter *VectorStoreExpirationAfter `json:"expires_after,omitempty"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
	//
	Metadata *UpdateVectorStoreRequestMetadata `json:"metadata,omitempty"`
}

func (o *UpdateVectorStoreRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateVectorStoreRequest) GetExpiresAfter() *VectorStoreExpirationAfter {
	if o == nil {
		return nil
	}
	return o.ExpiresAfter
}

func (o *UpdateVectorStoreRequest) GetMetadata() *UpdateVectorStoreRequestMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}
