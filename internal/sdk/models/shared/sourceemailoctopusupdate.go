// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceEmailoctopusUpdate struct {
	// EmailOctopus API Key. See the <a href="https://help.emailoctopus.com/article/165-how-to-create-and-delete-api-keys">docs</a> for information on how to generate this key.
	APIKey string `json:"api_key"`
}

func (o *SourceEmailoctopusUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}
