// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceGoogleAnalyticsDataAPICreateRequest struct {
	Configuration SourceGoogleAnalyticsDataAPI `json:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceGoogleAnalyticsDataAPICreateRequest) GetConfiguration() SourceGoogleAnalyticsDataAPI {
	if o == nil {
		return SourceGoogleAnalyticsDataAPI{}
	}
	return o.Configuration
}

func (o *SourceGoogleAnalyticsDataAPICreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *SourceGoogleAnalyticsDataAPICreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceGoogleAnalyticsDataAPICreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceGoogleAnalyticsDataAPICreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
