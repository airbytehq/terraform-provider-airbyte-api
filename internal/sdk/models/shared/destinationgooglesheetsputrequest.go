// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationGoogleSheetsPutRequest struct {
	Configuration DestinationGoogleSheetsUpdate `json:"configuration"`
	Name          string                        `json:"name"`
	WorkspaceID   string                        `json:"workspaceId"`
}

func (o *DestinationGoogleSheetsPutRequest) GetConfiguration() DestinationGoogleSheetsUpdate {
	if o == nil {
		return DestinationGoogleSheetsUpdate{}
	}
	return o.Configuration
}

func (o *DestinationGoogleSheetsPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationGoogleSheetsPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}