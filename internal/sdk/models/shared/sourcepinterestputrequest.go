// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourcePinterestPutRequest struct {
	Configuration SourcePinterestUpdate `json:"configuration"`
	Name          string                `json:"name"`
	WorkspaceID   string                `json:"workspaceId"`
}

func (o *SourcePinterestPutRequest) GetConfiguration() SourcePinterestUpdate {
	if o == nil {
		return SourcePinterestUpdate{}
	}
	return o.Configuration
}

func (o *SourcePinterestPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourcePinterestPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
