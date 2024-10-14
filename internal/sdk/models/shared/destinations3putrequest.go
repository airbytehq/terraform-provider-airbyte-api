// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DestinationS3PutRequest struct {
	Configuration DestinationS3Update `json:"configuration"`
	Name          string              `json:"name"`
	WorkspaceID   string              `json:"workspaceId"`
}

func (o *DestinationS3PutRequest) GetConfiguration() DestinationS3Update {
	if o == nil {
		return DestinationS3Update{}
	}
	return o.Configuration
}

func (o *DestinationS3PutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationS3PutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
