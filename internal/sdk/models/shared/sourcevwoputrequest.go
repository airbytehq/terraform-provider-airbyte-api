// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceVwoPutRequest struct {
	Configuration SourceVwoUpdate `json:"configuration"`
	Name          string          `json:"name"`
	WorkspaceID   string          `json:"workspaceId"`
}

func (o *SourceVwoPutRequest) GetConfiguration() SourceVwoUpdate {
	if o == nil {
		return SourceVwoUpdate{}
	}
	return o.Configuration
}

func (o *SourceVwoPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceVwoPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
