// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceMailjetSmsCreateRequest struct {
	Configuration SourceMailjetSms `json:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceMailjetSmsCreateRequest) GetConfiguration() SourceMailjetSms {
	if o == nil {
		return SourceMailjetSms{}
	}
	return o.Configuration
}

func (o *SourceMailjetSmsCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *SourceMailjetSmsCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceMailjetSmsCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceMailjetSmsCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
