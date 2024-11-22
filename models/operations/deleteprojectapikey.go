// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/jamelt/openai-api/models/components"
)

type DeleteProjectAPIKeyRequest struct {
	// The ID of the project.
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
	// The ID of the API key.
	KeyID string `pathParam:"style=simple,explode=false,name=key_id"`
}

func (o *DeleteProjectAPIKeyRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *DeleteProjectAPIKeyRequest) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

type DeleteProjectAPIKeyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Project API key deleted successfully.
	ProjectAPIKeyDeleteResponse *components.ProjectAPIKeyDeleteResponse
}

func (o *DeleteProjectAPIKeyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteProjectAPIKeyResponse) GetProjectAPIKeyDeleteResponse() *components.ProjectAPIKeyDeleteResponse {
	if o == nil {
		return nil
	}
	return o.ProjectAPIKeyDeleteResponse
}