// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceNorthpassLmsUpdate struct {
	APIKey string `json:"api_key"`
}

func (o *SourceNorthpassLmsUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}
