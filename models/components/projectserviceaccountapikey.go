// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ProjectServiceAccountAPIKeyObject - The object type, which is always `organization.project.service_account.api_key`
type ProjectServiceAccountAPIKeyObject string

const (
	ProjectServiceAccountAPIKeyObjectOrganizationProjectServiceAccountAPIKey ProjectServiceAccountAPIKeyObject = "organization.project.service_account.api_key"
)

func (e ProjectServiceAccountAPIKeyObject) ToPointer() *ProjectServiceAccountAPIKeyObject {
	return &e
}
func (e *ProjectServiceAccountAPIKeyObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "organization.project.service_account.api_key":
		*e = ProjectServiceAccountAPIKeyObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProjectServiceAccountAPIKeyObject: %v", v)
	}
}

type ProjectServiceAccountAPIKey struct {
	// The object type, which is always `organization.project.service_account.api_key`
	Object    ProjectServiceAccountAPIKeyObject `json:"object"`
	Value     string                            `json:"value"`
	Name      string                            `json:"name"`
	CreatedAt int64                             `json:"created_at"`
	ID        string                            `json:"id"`
}

func (o *ProjectServiceAccountAPIKey) GetObject() ProjectServiceAccountAPIKeyObject {
	if o == nil {
		return ProjectServiceAccountAPIKeyObject("")
	}
	return o.Object
}

func (o *ProjectServiceAccountAPIKey) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *ProjectServiceAccountAPIKey) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ProjectServiceAccountAPIKey) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *ProjectServiceAccountAPIKey) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
