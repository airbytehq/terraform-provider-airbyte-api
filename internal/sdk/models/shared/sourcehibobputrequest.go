// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceHibobPutRequest struct {
	Configuration SourceHibobUpdate `json:"configuration"`
	Name          string            `json:"name"`
	WorkspaceID   string            `json:"workspaceId"`
}

func (o *SourceHibobPutRequest) GetConfiguration() SourceHibobUpdate {
	if o == nil {
		return SourceHibobUpdate{}
	}
	return o.Configuration
}

func (o *SourceHibobPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceHibobPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
