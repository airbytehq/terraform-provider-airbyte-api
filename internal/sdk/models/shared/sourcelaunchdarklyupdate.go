// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceLaunchdarklyUpdate struct {
	// Your Access token. See <a href="https://apidocs.launchdarkly.com/#section/Overview/Authentication">here</a>.
	AccessToken string `json:"access_token"`
}

func (o *SourceLaunchdarklyUpdate) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}
