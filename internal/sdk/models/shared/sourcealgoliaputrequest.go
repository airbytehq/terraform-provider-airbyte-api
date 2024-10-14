// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceAlgoliaPutRequest struct {
	Configuration SourceAlgoliaUpdate `json:"configuration"`
	Name          string              `json:"name"`
	WorkspaceID   string              `json:"workspaceId"`
}

func (o *SourceAlgoliaPutRequest) GetConfiguration() SourceAlgoliaUpdate {
	if o == nil {
		return SourceAlgoliaUpdate{}
	}
	return o.Configuration
}

func (o *SourceAlgoliaPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceAlgoliaPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
