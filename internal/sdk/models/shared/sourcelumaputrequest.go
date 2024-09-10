// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceLumaPutRequest struct {
	Configuration SourceLumaUpdate `json:"configuration"`
	Name          string           `json:"name"`
	WorkspaceID   string           `json:"workspaceId"`
}

func (o *SourceLumaPutRequest) GetConfiguration() SourceLumaUpdate {
	if o == nil {
		return SourceLumaUpdate{}
	}
	return o.Configuration
}

func (o *SourceLumaPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceLumaPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
