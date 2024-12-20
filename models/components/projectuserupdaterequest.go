// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ProjectUserUpdateRequestRole - `owner` or `member`
type ProjectUserUpdateRequestRole string

const (
	ProjectUserUpdateRequestRoleOwner  ProjectUserUpdateRequestRole = "owner"
	ProjectUserUpdateRequestRoleMember ProjectUserUpdateRequestRole = "member"
)

func (e ProjectUserUpdateRequestRole) ToPointer() *ProjectUserUpdateRequestRole {
	return &e
}
func (e *ProjectUserUpdateRequestRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "owner":
		fallthrough
	case "member":
		*e = ProjectUserUpdateRequestRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProjectUserUpdateRequestRole: %v", v)
	}
}

type ProjectUserUpdateRequest struct {
	// `owner` or `member`
	Role ProjectUserUpdateRequestRole `json:"role"`
}

func (o *ProjectUserUpdateRequest) GetRole() ProjectUserUpdateRequestRole {
	if o == nil {
		return ProjectUserUpdateRequestRole("")
	}
	return o.Role
}
