// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceSnowflakePutRequest struct {
	Configuration SourceSnowflakeUpdate `json:"configuration"`
	Name          string                `json:"name"`
	WorkspaceID   string                `json:"workspaceId"`
}

func (o *SourceSnowflakePutRequest) GetConfiguration() SourceSnowflakeUpdate {
	if o == nil {
		return SourceSnowflakeUpdate{}
	}
	return o.Configuration
}

func (o *SourceSnowflakePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceSnowflakePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}