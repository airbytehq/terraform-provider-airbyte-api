// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceDbtPutRequest struct {
	Configuration SourceDbtUpdate `json:"configuration"`
	Name          string          `json:"name"`
	WorkspaceID   string          `json:"workspaceId"`
}

func (o *SourceDbtPutRequest) GetConfiguration() SourceDbtUpdate {
	if o == nil {
		return SourceDbtUpdate{}
	}
	return o.Configuration
}

func (o *SourceDbtPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceDbtPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
