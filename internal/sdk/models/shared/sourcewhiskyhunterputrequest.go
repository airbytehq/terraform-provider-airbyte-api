// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceWhiskyHunterPutRequest struct {
	Configuration SourceWhiskyHunterUpdate `json:"configuration"`
	Name          string                   `json:"name"`
	WorkspaceID   string                   `json:"workspaceId"`
}

func (o *SourceWhiskyHunterPutRequest) GetConfiguration() SourceWhiskyHunterUpdate {
	if o == nil {
		return SourceWhiskyHunterUpdate{}
	}
	return o.Configuration
}

func (o *SourceWhiskyHunterPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceWhiskyHunterPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
