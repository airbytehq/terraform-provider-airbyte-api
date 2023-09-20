// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceRecurlyGetResponse struct {
	Configuration SourceRecurly `json:"configuration"`
	Name          string        `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	SourceID    *string `json:"sourceId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}