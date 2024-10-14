// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceGreenhousePutRequest struct {
	Configuration SourceGreenhouseUpdate `json:"configuration"`
	Name          string                 `json:"name"`
	WorkspaceID   string                 `json:"workspaceId"`
}

func (o *SourceGreenhousePutRequest) GetConfiguration() SourceGreenhouseUpdate {
	if o == nil {
		return SourceGreenhouseUpdate{}
	}
	return o.Configuration
}

func (o *SourceGreenhousePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceGreenhousePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
