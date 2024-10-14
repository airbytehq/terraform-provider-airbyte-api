// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceCoinmarketcapCreateRequest struct {
	Configuration SourceCoinmarketcap `json:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceCoinmarketcapCreateRequest) GetConfiguration() SourceCoinmarketcap {
	if o == nil {
		return SourceCoinmarketcap{}
	}
	return o.Configuration
}

func (o *SourceCoinmarketcapCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *SourceCoinmarketcapCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceCoinmarketcapCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceCoinmarketcapCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
