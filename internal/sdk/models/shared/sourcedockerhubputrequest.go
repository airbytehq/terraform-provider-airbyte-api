// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceDockerhubPutRequest struct {
	Configuration SourceDockerhubUpdate `json:"configuration"`
	Name          string                `json:"name"`
	WorkspaceID   string                `json:"workspaceId"`
}

func (o *SourceDockerhubPutRequest) GetConfiguration() SourceDockerhubUpdate {
	if o == nil {
		return SourceDockerhubUpdate{}
	}
	return o.Configuration
}

func (o *SourceDockerhubPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceDockerhubPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}