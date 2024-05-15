// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceGainsightPxPutRequest struct {
	Configuration SourceGainsightPxUpdate `json:"configuration"`
	Name          string                  `json:"name"`
	WorkspaceID   string                  `json:"workspaceId"`
}

func (o *SourceGainsightPxPutRequest) GetConfiguration() SourceGainsightPxUpdate {
	if o == nil {
		return SourceGainsightPxUpdate{}
	}
	return o.Configuration
}

func (o *SourceGainsightPxPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceGainsightPxPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}