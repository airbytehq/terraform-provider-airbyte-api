// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DestinationGoogleSheetsCreateRequest struct {
	Configuration DestinationGoogleSheets `json:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
}

func (o *DestinationGoogleSheetsCreateRequest) GetConfiguration() DestinationGoogleSheets {
	if o == nil {
		return DestinationGoogleSheets{}
	}
	return o.Configuration
}

func (o *DestinationGoogleSheetsCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *DestinationGoogleSheetsCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationGoogleSheetsCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
