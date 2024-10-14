// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PermissionScope - Scope of a single permission, e.g. workspace, organization
type PermissionScope string

const (
	PermissionScopeWorkspace    PermissionScope = "workspace"
	PermissionScopeOrganization PermissionScope = "organization"
	PermissionScopeNone         PermissionScope = "none"
)

func (e PermissionScope) ToPointer() *PermissionScope {
	return &e
}
func (e *PermissionScope) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workspace":
		fallthrough
	case "organization":
		fallthrough
	case "none":
		*e = PermissionScope(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PermissionScope: %v", v)
	}
}
