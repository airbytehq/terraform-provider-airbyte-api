// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceBreezyHrUpdate struct {
	APIKey    string `json:"api_key"`
	CompanyID string `json:"company_id"`
}

func (o *SourceBreezyHrUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceBreezyHrUpdate) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}
