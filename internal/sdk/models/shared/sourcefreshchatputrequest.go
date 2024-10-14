// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceFreshchatPutRequest struct {
	Configuration SourceFreshchatUpdate `json:"configuration"`
	Name          string                `json:"name"`
	WorkspaceID   string                `json:"workspaceId"`
}

func (o *SourceFreshchatPutRequest) GetConfiguration() SourceFreshchatUpdate {
	if o == nil {
		return SourceFreshchatUpdate{}
	}
	return o.Configuration
}

func (o *SourceFreshchatPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceFreshchatPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
