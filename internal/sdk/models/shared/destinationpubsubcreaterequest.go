// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DestinationPubsubCreateRequest struct {
	Configuration DestinationPubsub `json:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
}

func (o *DestinationPubsubCreateRequest) GetConfiguration() DestinationPubsub {
	if o == nil {
		return DestinationPubsub{}
	}
	return o.Configuration
}

func (o *DestinationPubsubCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *DestinationPubsubCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationPubsubCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
