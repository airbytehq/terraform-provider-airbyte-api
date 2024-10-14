// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceSlackPutRequest struct {
	Configuration SourceSlackUpdate `json:"configuration"`
	Name          string            `json:"name"`
	WorkspaceID   string            `json:"workspaceId"`
}

func (o *SourceSlackPutRequest) GetConfiguration() SourceSlackUpdate {
	if o == nil {
		return SourceSlackUpdate{}
	}
	return o.Configuration
}

func (o *SourceSlackPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceSlackPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
