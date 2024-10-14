// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourcePostgresPutRequest struct {
	Configuration SourcePostgresUpdate `json:"configuration"`
	Name          string               `json:"name"`
	WorkspaceID   string               `json:"workspaceId"`
}

func (o *SourcePostgresPutRequest) GetConfiguration() SourcePostgresUpdate {
	if o == nil {
		return SourcePostgresUpdate{}
	}
	return o.Configuration
}

func (o *SourcePostgresPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourcePostgresPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
