// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceXkcdPutRequest struct {
	Configuration SourceXkcdUpdate `json:"configuration"`
	Name          string           `json:"name"`
	WorkspaceID   string           `json:"workspaceId"`
}

func (o *SourceXkcdPutRequest) GetConfiguration() SourceXkcdUpdate {
	if o == nil {
		return SourceXkcdUpdate{}
	}
	return o.Configuration
}

func (o *SourceXkcdPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceXkcdPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
