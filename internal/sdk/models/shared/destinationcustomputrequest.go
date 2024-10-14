// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DestinationCustomPutRequest struct {
	Configuration any    `json:"configuration"`
	Name          string `json:"name"`
	WorkspaceID   string `json:"workspaceId"`
}

func (o *DestinationCustomPutRequest) GetConfiguration() any {
	if o == nil {
		return nil
	}
	return o.Configuration
}

func (o *DestinationCustomPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationCustomPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
