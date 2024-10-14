// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceZendeskSunshineCreateRequest struct {
	Configuration SourceZendeskSunshine `json:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceZendeskSunshineCreateRequest) GetConfiguration() SourceZendeskSunshine {
	if o == nil {
		return SourceZendeskSunshine{}
	}
	return o.Configuration
}

func (o *SourceZendeskSunshineCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *SourceZendeskSunshineCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceZendeskSunshineCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceZendeskSunshineCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
