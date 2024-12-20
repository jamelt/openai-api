// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AuditLogActorAPIKeyType - The type of API key. Can be either `user` or `service_account`.
type AuditLogActorAPIKeyType string

const (
	AuditLogActorAPIKeyTypeUser           AuditLogActorAPIKeyType = "user"
	AuditLogActorAPIKeyTypeServiceAccount AuditLogActorAPIKeyType = "service_account"
)

func (e AuditLogActorAPIKeyType) ToPointer() *AuditLogActorAPIKeyType {
	return &e
}
func (e *AuditLogActorAPIKeyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "user":
		fallthrough
	case "service_account":
		*e = AuditLogActorAPIKeyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuditLogActorAPIKeyType: %v", v)
	}
}

// AuditLogActorAPIKey - The API Key used to perform the audit logged action.
type AuditLogActorAPIKey struct {
	// The tracking id of the API key.
	ID *string `json:"id,omitempty"`
	// The type of API key. Can be either `user` or `service_account`.
	Type *AuditLogActorAPIKeyType `json:"type,omitempty"`
	// The user who performed the audit logged action.
	User *AuditLogActorUser `json:"user,omitempty"`
	// The service account that performed the audit logged action.
	ServiceAccount *AuditLogActorServiceAccount `json:"service_account,omitempty"`
}

func (o *AuditLogActorAPIKey) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AuditLogActorAPIKey) GetType() *AuditLogActorAPIKeyType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AuditLogActorAPIKey) GetUser() *AuditLogActorUser {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *AuditLogActorAPIKey) GetServiceAccount() *AuditLogActorServiceAccount {
	if o == nil {
		return nil
	}
	return o.ServiceAccount
}
