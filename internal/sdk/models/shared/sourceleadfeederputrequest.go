// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceLeadfeederPutRequest struct {
	Configuration SourceLeadfeederUpdate `json:"configuration"`
	Name          string                 `json:"name"`
	WorkspaceID   string                 `json:"workspaceId"`
}

func (o *SourceLeadfeederPutRequest) GetConfiguration() SourceLeadfeederUpdate {
	if o == nil {
		return SourceLeadfeederUpdate{}
	}
	return o.Configuration
}

func (o *SourceLeadfeederPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceLeadfeederPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
