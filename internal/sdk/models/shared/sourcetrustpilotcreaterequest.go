// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceTrustpilotCreateRequest struct {
	Configuration SourceTrustpilot `json:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceTrustpilotCreateRequest) GetConfiguration() SourceTrustpilot {
	if o == nil {
		return SourceTrustpilot{}
	}
	return o.Configuration
}

func (o *SourceTrustpilotCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *SourceTrustpilotCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceTrustpilotCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceTrustpilotCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
