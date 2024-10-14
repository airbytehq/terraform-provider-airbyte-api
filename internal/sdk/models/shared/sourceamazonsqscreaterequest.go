// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceAmazonSqsCreateRequest struct {
	Configuration SourceAmazonSqs `json:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceAmazonSqsCreateRequest) GetConfiguration() SourceAmazonSqs {
	if o == nil {
		return SourceAmazonSqs{}
	}
	return o.Configuration
}

func (o *SourceAmazonSqsCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *SourceAmazonSqsCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceAmazonSqsCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceAmazonSqsCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
