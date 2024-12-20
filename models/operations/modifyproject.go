// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/jamelt/openai-api/models/components"
)

type ModifyProjectRequest struct {
	// The ID of the project.
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
	// The project update request payload.
	ProjectUpdateRequest components.ProjectUpdateRequest `request:"mediaType=application/json"`
}

func (o *ModifyProjectRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *ModifyProjectRequest) GetProjectUpdateRequest() components.ProjectUpdateRequest {
	if o == nil {
		return components.ProjectUpdateRequest{}
	}
	return o.ProjectUpdateRequest
}

type ModifyProjectResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Project updated successfully.
	Project *components.Project
}

func (o *ModifyProjectResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ModifyProjectResponse) GetProject() *components.Project {
	if o == nil {
		return nil
	}
	return o.Project
}
