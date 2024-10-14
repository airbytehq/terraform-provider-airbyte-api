// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceLookerPutRequest struct {
	Configuration SourceLookerUpdate `json:"configuration"`
	Name          string             `json:"name"`
	WorkspaceID   string             `json:"workspaceId"`
}

func (o *SourceLookerPutRequest) GetConfiguration() SourceLookerUpdate {
	if o == nil {
		return SourceLookerUpdate{}
	}
	return o.Configuration
}

func (o *SourceLookerPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceLookerPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
