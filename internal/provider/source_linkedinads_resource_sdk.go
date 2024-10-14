// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	customTypes "github.com/airbytehq/terraform-provider-airbyte/internal/sdk/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceLinkedinAdsResourceModel) ToSharedSourceLinkedinAdsCreateRequest() *shared.SourceLinkedinAdsCreateRequest {
	var accountIds []int64 = []int64{}
	for _, accountIdsItem := range r.Configuration.AccountIds {
		accountIds = append(accountIds, accountIdsItem.ValueInt64())
	}
	var adAnalyticsReports []shared.SourceLinkedinAdsAdAnalyticsReportConfiguration = []shared.SourceLinkedinAdsAdAnalyticsReportConfiguration{}
	for _, adAnalyticsReportsItem := range r.Configuration.AdAnalyticsReports {
		var name string
		name = adAnalyticsReportsItem.Name.ValueString()

		pivotBy := shared.SourceLinkedinAdsPivotCategory(adAnalyticsReportsItem.PivotBy.ValueString())
		timeGranularity := shared.SourceLinkedinAdsTimeGranularity(adAnalyticsReportsItem.TimeGranularity.ValueString())
		adAnalyticsReports = append(adAnalyticsReports, shared.SourceLinkedinAdsAdAnalyticsReportConfiguration{
			Name:            name,
			PivotBy:         pivotBy,
			TimeGranularity: timeGranularity,
		})
	}
	var credentials *shared.SourceLinkedinAdsAuthentication
	if r.Configuration.Credentials != nil {
		var sourceLinkedinAdsOAuth20 *shared.SourceLinkedinAdsOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			var clientID string
			clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()

			var clientSecret string
			clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()

			var refreshToken string
			refreshToken = r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()

			sourceLinkedinAdsOAuth20 = &shared.SourceLinkedinAdsOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceLinkedinAdsOAuth20 != nil {
			credentials = &shared.SourceLinkedinAdsAuthentication{
				SourceLinkedinAdsOAuth20: sourceLinkedinAdsOAuth20,
			}
		}
		var sourceLinkedinAdsAccessToken *shared.SourceLinkedinAdsAccessToken
		if r.Configuration.Credentials.AccessToken != nil {
			var accessToken string
			accessToken = r.Configuration.Credentials.AccessToken.AccessToken.ValueString()

			sourceLinkedinAdsAccessToken = &shared.SourceLinkedinAdsAccessToken{
				AccessToken: accessToken,
			}
		}
		if sourceLinkedinAdsAccessToken != nil {
			credentials = &shared.SourceLinkedinAdsAuthentication{
				SourceLinkedinAdsAccessToken: sourceLinkedinAdsAccessToken,
			}
		}
	}
	lookbackWindow := new(int64)
	if !r.Configuration.LookbackWindow.IsUnknown() && !r.Configuration.LookbackWindow.IsNull() {
		*lookbackWindow = r.Configuration.LookbackWindow.ValueInt64()
	} else {
		lookbackWindow = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceLinkedinAds{
		AccountIds:         accountIds,
		AdAnalyticsReports: adAnalyticsReports,
		Credentials:        credentials,
		LookbackWindow:     lookbackWindow,
		StartDate:          startDate,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	var name1 string
	name1 = r.Name.ValueString()

	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceLinkedinAdsCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name1,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLinkedinAdsResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceLinkedinAdsResourceModel) ToSharedSourceLinkedinAdsPutRequest() *shared.SourceLinkedinAdsPutRequest {
	var accountIds []int64 = []int64{}
	for _, accountIdsItem := range r.Configuration.AccountIds {
		accountIds = append(accountIds, accountIdsItem.ValueInt64())
	}
	var adAnalyticsReports []shared.AdAnalyticsReportConfiguration = []shared.AdAnalyticsReportConfiguration{}
	for _, adAnalyticsReportsItem := range r.Configuration.AdAnalyticsReports {
		var name string
		name = adAnalyticsReportsItem.Name.ValueString()

		pivotBy := shared.PivotCategory(adAnalyticsReportsItem.PivotBy.ValueString())
		timeGranularity := shared.TimeGranularity(adAnalyticsReportsItem.TimeGranularity.ValueString())
		adAnalyticsReports = append(adAnalyticsReports, shared.AdAnalyticsReportConfiguration{
			Name:            name,
			PivotBy:         pivotBy,
			TimeGranularity: timeGranularity,
		})
	}
	var credentials *shared.SourceLinkedinAdsUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var sourceLinkedinAdsUpdateOAuth20 *shared.SourceLinkedinAdsUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			var clientID string
			clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()

			var clientSecret string
			clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()

			var refreshToken string
			refreshToken = r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()

			sourceLinkedinAdsUpdateOAuth20 = &shared.SourceLinkedinAdsUpdateOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceLinkedinAdsUpdateOAuth20 != nil {
			credentials = &shared.SourceLinkedinAdsUpdateAuthentication{
				SourceLinkedinAdsUpdateOAuth20: sourceLinkedinAdsUpdateOAuth20,
			}
		}
		var accessToken *shared.AccessToken
		if r.Configuration.Credentials.AccessToken != nil {
			var accessToken1 string
			accessToken1 = r.Configuration.Credentials.AccessToken.AccessToken.ValueString()

			accessToken = &shared.AccessToken{
				AccessToken: accessToken1,
			}
		}
		if accessToken != nil {
			credentials = &shared.SourceLinkedinAdsUpdateAuthentication{
				AccessToken: accessToken,
			}
		}
	}
	lookbackWindow := new(int64)
	if !r.Configuration.LookbackWindow.IsUnknown() && !r.Configuration.LookbackWindow.IsNull() {
		*lookbackWindow = r.Configuration.LookbackWindow.ValueInt64()
	} else {
		lookbackWindow = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceLinkedinAdsUpdate{
		AccountIds:         accountIds,
		AdAnalyticsReports: adAnalyticsReports,
		Credentials:        credentials,
		LookbackWindow:     lookbackWindow,
		StartDate:          startDate,
	}
	var name1 string
	name1 = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceLinkedinAdsPutRequest{
		Configuration: configuration,
		Name:          name1,
		WorkspaceID:   workspaceID,
	}
	return &out
}
