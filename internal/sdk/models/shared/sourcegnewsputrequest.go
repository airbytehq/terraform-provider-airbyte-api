// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceGnewsPutRequest struct {
	Configuration SourceGnewsUpdate `json:"configuration"`
	Name          string            `json:"name"`
	WorkspaceID   string            `json:"workspaceId"`
}

func (o *SourceGnewsPutRequest) GetConfiguration() SourceGnewsUpdate {
	if o == nil {
		return SourceGnewsUpdate{}
	}
	return o.Configuration
}

func (o *SourceGnewsPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceGnewsPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
