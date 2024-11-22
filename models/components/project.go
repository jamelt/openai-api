// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ProjectObject - The object type, which is always `organization.project`
type ProjectObject string

const (
	ProjectObjectOrganizationProject ProjectObject = "organization.project"
)

func (e ProjectObject) ToPointer() *ProjectObject {
	return &e
}
func (e *ProjectObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "organization.project":
		*e = ProjectObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProjectObject: %v", v)
	}
}

// ProjectStatus - `active` or `archived`
type ProjectStatus string

const (
	ProjectStatusActive   ProjectStatus = "active"
	ProjectStatusArchived ProjectStatus = "archived"
)

func (e ProjectStatus) ToPointer() *ProjectStatus {
	return &e
}
func (e *ProjectStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "archived":
		*e = ProjectStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProjectStatus: %v", v)
	}
}

// Project - Represents an individual project.
type Project struct {
	// The identifier, which can be referenced in API endpoints
	ID string `json:"id"`
	// The object type, which is always `organization.project`
	Object ProjectObject `json:"object"`
	// The name of the project. This appears in reporting.
	Name string `json:"name"`
	// The Unix timestamp (in seconds) of when the project was created.
	CreatedAt int64 `json:"created_at"`
	// The Unix timestamp (in seconds) of when the project was archived or `null`.
	ArchivedAt *int64 `json:"archived_at,omitempty"`
	// `active` or `archived`
	Status ProjectStatus `json:"status"`
}

func (o *Project) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Project) GetObject() ProjectObject {
	if o == nil {
		return ProjectObject("")
	}
	return o.Object
}

func (o *Project) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Project) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *Project) GetArchivedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *Project) GetStatus() ProjectStatus {
	if o == nil {
		return ProjectStatus("")
	}
	return o.Status
}