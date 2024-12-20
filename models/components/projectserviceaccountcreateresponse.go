// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type ProjectServiceAccountCreateResponseObject string

const (
	ProjectServiceAccountCreateResponseObjectOrganizationProjectServiceAccount ProjectServiceAccountCreateResponseObject = "organization.project.service_account"
)

func (e ProjectServiceAccountCreateResponseObject) ToPointer() *ProjectServiceAccountCreateResponseObject {
	return &e
}
func (e *ProjectServiceAccountCreateResponseObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "organization.project.service_account":
		*e = ProjectServiceAccountCreateResponseObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProjectServiceAccountCreateResponseObject: %v", v)
	}
}

// ProjectServiceAccountCreateResponseRole - Service accounts can only have one role of type `member`
type ProjectServiceAccountCreateResponseRole string

const (
	ProjectServiceAccountCreateResponseRoleMember ProjectServiceAccountCreateResponseRole = "member"
)

func (e ProjectServiceAccountCreateResponseRole) ToPointer() *ProjectServiceAccountCreateResponseRole {
	return &e
}
func (e *ProjectServiceAccountCreateResponseRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "member":
		*e = ProjectServiceAccountCreateResponseRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProjectServiceAccountCreateResponseRole: %v", v)
	}
}

type ProjectServiceAccountCreateResponse struct {
	Object ProjectServiceAccountCreateResponseObject `json:"object"`
	ID     string                                    `json:"id"`
	Name   string                                    `json:"name"`
	// Service accounts can only have one role of type `member`
	Role      ProjectServiceAccountCreateResponseRole `json:"role"`
	CreatedAt int64                                   `json:"created_at"`
	APIKey    ProjectServiceAccountAPIKey             `json:"api_key"`
}

func (o *ProjectServiceAccountCreateResponse) GetObject() ProjectServiceAccountCreateResponseObject {
	if o == nil {
		return ProjectServiceAccountCreateResponseObject("")
	}
	return o.Object
}

func (o *ProjectServiceAccountCreateResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ProjectServiceAccountCreateResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ProjectServiceAccountCreateResponse) GetRole() ProjectServiceAccountCreateResponseRole {
	if o == nil {
		return ProjectServiceAccountCreateResponseRole("")
	}
	return o.Role
}

func (o *ProjectServiceAccountCreateResponse) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *ProjectServiceAccountCreateResponse) GetAPIKey() ProjectServiceAccountAPIKey {
	if o == nil {
		return ProjectServiceAccountAPIKey{}
	}
	return o.APIKey
}
