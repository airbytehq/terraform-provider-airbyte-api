// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceExchangeRatesPutRequest struct {
	Configuration SourceExchangeRatesUpdate `json:"configuration"`
	Name          string                    `json:"name"`
	WorkspaceID   string                    `json:"workspaceId"`
}

func (o *SourceExchangeRatesPutRequest) GetConfiguration() SourceExchangeRatesUpdate {
	if o == nil {
		return SourceExchangeRatesUpdate{}
	}
	return o.Configuration
}

func (o *SourceExchangeRatesPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceExchangeRatesPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}