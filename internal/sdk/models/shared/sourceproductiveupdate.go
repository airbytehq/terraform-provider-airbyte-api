// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceProductiveUpdate struct {
	APIKey string `json:"api_key"`
	// The organization ID which could be seen from `https://app.productive.io/xxxx-xxxx/settings/api-integrations` page
	OrganizationID string `json:"organization_id"`
}

func (o *SourceProductiveUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceProductiveUpdate) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}
