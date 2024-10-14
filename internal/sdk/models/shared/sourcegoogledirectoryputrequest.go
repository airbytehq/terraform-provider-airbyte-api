// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceGoogleDirectoryPutRequest struct {
	Configuration SourceGoogleDirectoryUpdate `json:"configuration"`
	Name          string                      `json:"name"`
	WorkspaceID   string                      `json:"workspaceId"`
}

func (o *SourceGoogleDirectoryPutRequest) GetConfiguration() SourceGoogleDirectoryUpdate {
	if o == nil {
		return SourceGoogleDirectoryUpdate{}
	}
	return o.Configuration
}

func (o *SourceGoogleDirectoryPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceGoogleDirectoryPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
