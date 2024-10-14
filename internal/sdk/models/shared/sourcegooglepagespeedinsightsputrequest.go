// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceGooglePagespeedInsightsPutRequest struct {
	Configuration SourceGooglePagespeedInsightsUpdate `json:"configuration"`
	Name          string                              `json:"name"`
	WorkspaceID   string                              `json:"workspaceId"`
}

func (o *SourceGooglePagespeedInsightsPutRequest) GetConfiguration() SourceGooglePagespeedInsightsUpdate {
	if o == nil {
		return SourceGooglePagespeedInsightsUpdate{}
	}
	return o.Configuration
}

func (o *SourceGooglePagespeedInsightsPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceGooglePagespeedInsightsPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
