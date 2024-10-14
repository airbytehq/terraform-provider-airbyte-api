// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceShopifyPutRequest struct {
	Configuration SourceShopifyUpdate `json:"configuration"`
	Name          string              `json:"name"`
	WorkspaceID   string              `json:"workspaceId"`
}

func (o *SourceShopifyPutRequest) GetConfiguration() SourceShopifyUpdate {
	if o == nil {
		return SourceShopifyUpdate{}
	}
	return o.Configuration
}

func (o *SourceShopifyPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceShopifyPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
