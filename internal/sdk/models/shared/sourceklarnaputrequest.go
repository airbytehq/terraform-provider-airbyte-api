// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceKlarnaPutRequest struct {
	Configuration SourceKlarnaUpdate `json:"configuration"`
	Name          string             `json:"name"`
	WorkspaceID   string             `json:"workspaceId"`
}

func (o *SourceKlarnaPutRequest) GetConfiguration() SourceKlarnaUpdate {
	if o == nil {
		return SourceKlarnaUpdate{}
	}
	return o.Configuration
}

func (o *SourceKlarnaPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceKlarnaPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
