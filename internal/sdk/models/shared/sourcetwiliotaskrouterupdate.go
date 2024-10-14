// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceTwilioTaskrouterUpdate struct {
	// Twilio Account ID
	AccountSid string `json:"account_sid"`
	// Twilio Auth Token
	AuthToken string `json:"auth_token"`
}

func (o *SourceTwilioTaskrouterUpdate) GetAccountSid() string {
	if o == nil {
		return ""
	}
	return o.AccountSid
}

func (o *SourceTwilioTaskrouterUpdate) GetAuthToken() string {
	if o == nil {
		return ""
	}
	return o.AuthToken
}
