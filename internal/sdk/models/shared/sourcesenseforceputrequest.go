// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceSenseforcePutRequest struct {
	Configuration SourceSenseforceUpdate `json:"configuration"`
	Name          string                 `json:"name"`
	WorkspaceID   string                 `json:"workspaceId"`
}

func (o *SourceSenseforcePutRequest) GetConfiguration() SourceSenseforceUpdate {
	if o == nil {
		return SourceSenseforceUpdate{}
	}
	return o.Configuration
}

func (o *SourceSenseforcePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceSenseforcePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
