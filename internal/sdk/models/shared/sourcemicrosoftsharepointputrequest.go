// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceMicrosoftSharepointPutRequest struct {
	// SourceMicrosoftSharePointSpec class for Microsoft SharePoint Source Specification.
	// This class combines the authentication details with additional configuration for the SharePoint API.
	Configuration SourceMicrosoftSharepointUpdate `json:"configuration"`
	Name          string                          `json:"name"`
	WorkspaceID   string                          `json:"workspaceId"`
}

func (o *SourceMicrosoftSharepointPutRequest) GetConfiguration() SourceMicrosoftSharepointUpdate {
	if o == nil {
		return SourceMicrosoftSharepointUpdate{}
	}
	return o.Configuration
}

func (o *SourceMicrosoftSharepointPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceMicrosoftSharepointPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
