// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourcePaypalTransactionPutRequest struct {
	Configuration SourcePaypalTransactionUpdate `json:"configuration"`
	Name          string                        `json:"name"`
	WorkspaceID   string                        `json:"workspaceId"`
}

func (o *SourcePaypalTransactionPutRequest) GetConfiguration() SourcePaypalTransactionUpdate {
	if o == nil {
		return SourcePaypalTransactionUpdate{}
	}
	return o.Configuration
}

func (o *SourcePaypalTransactionPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourcePaypalTransactionPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}