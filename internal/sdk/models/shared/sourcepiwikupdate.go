// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourcePiwikUpdate struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	// The organization id appearing at URL of your piwik website
	OrganizationID string `json:"organization_id"`
}

func (o *SourcePiwikUpdate) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourcePiwikUpdate) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourcePiwikUpdate) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}
