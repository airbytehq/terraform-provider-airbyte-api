// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DestinationMssqlCreateRequest struct {
	Configuration DestinationMssql `json:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
}

func (o *DestinationMssqlCreateRequest) GetConfiguration() DestinationMssql {
	if o == nil {
		return DestinationMssql{}
	}
	return o.Configuration
}

func (o *DestinationMssqlCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *DestinationMssqlCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationMssqlCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
