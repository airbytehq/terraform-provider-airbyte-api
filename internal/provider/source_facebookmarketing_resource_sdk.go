// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceFacebookMarketingResourceModel) ToSharedSourceFacebookMarketingCreateRequest() *shared.SourceFacebookMarketingCreateRequest {
	accessToken := new(string)
	if !r.Configuration.AccessToken.IsUnknown() && !r.Configuration.AccessToken.IsNull() {
		*accessToken = r.Configuration.AccessToken.ValueString()
	} else {
		accessToken = nil
	}
	var accountIds []string = []string{}
	for _, accountIdsItem := range r.Configuration.AccountIds {
		accountIds = append(accountIds, accountIdsItem.ValueString())
	}
	actionBreakdownsAllowEmpty := new(bool)
	if !r.Configuration.ActionBreakdownsAllowEmpty.IsUnknown() && !r.Configuration.ActionBreakdownsAllowEmpty.IsNull() {
		*actionBreakdownsAllowEmpty = r.Configuration.ActionBreakdownsAllowEmpty.ValueBool()
	} else {
		actionBreakdownsAllowEmpty = nil
	}
	var adStatuses []shared.SourceFacebookMarketingValidAdStatuses = []shared.SourceFacebookMarketingValidAdStatuses{}
	for _, adStatusesItem := range r.Configuration.AdStatuses {
		adStatuses = append(adStatuses, shared.SourceFacebookMarketingValidAdStatuses(adStatusesItem.ValueString()))
	}
	var adsetStatuses []shared.SourceFacebookMarketingValidAdSetStatuses = []shared.SourceFacebookMarketingValidAdSetStatuses{}
	for _, adsetStatusesItem := range r.Configuration.AdsetStatuses {
		adsetStatuses = append(adsetStatuses, shared.SourceFacebookMarketingValidAdSetStatuses(adsetStatusesItem.ValueString()))
	}
	var campaignStatuses []shared.SourceFacebookMarketingValidCampaignStatuses = []shared.SourceFacebookMarketingValidCampaignStatuses{}
	for _, campaignStatusesItem := range r.Configuration.CampaignStatuses {
		campaignStatuses = append(campaignStatuses, shared.SourceFacebookMarketingValidCampaignStatuses(campaignStatusesItem.ValueString()))
	}
	clientID := new(string)
	if !r.Configuration.ClientID.IsUnknown() && !r.Configuration.ClientID.IsNull() {
		*clientID = r.Configuration.ClientID.ValueString()
	} else {
		clientID = nil
	}
	clientSecret := new(string)
	if !r.Configuration.ClientSecret.IsUnknown() && !r.Configuration.ClientSecret.IsNull() {
		*clientSecret = r.Configuration.ClientSecret.ValueString()
	} else {
		clientSecret = nil
	}
	var credentials *shared.SourceFacebookMarketingAuthentication
	if r.Configuration.Credentials != nil {
		var sourceFacebookMarketingAuthenticateViaFacebookMarketingOauth *shared.SourceFacebookMarketingAuthenticateViaFacebookMarketingOauth
		if r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth != nil {
			accessToken1 := new(string)
			if !r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth.AccessToken.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth.AccessToken.IsNull() {
				*accessToken1 = r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth.AccessToken.ValueString()
			} else {
				accessToken1 = nil
			}
			clientId1 := r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth.ClientID.ValueString()
			clientSecret1 := r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth.ClientSecret.ValueString()
			sourceFacebookMarketingAuthenticateViaFacebookMarketingOauth = &shared.SourceFacebookMarketingAuthenticateViaFacebookMarketingOauth{
				AccessToken:  accessToken1,
				ClientID:     clientId1,
				ClientSecret: clientSecret1,
			}
		}
		if sourceFacebookMarketingAuthenticateViaFacebookMarketingOauth != nil {
			credentials = &shared.SourceFacebookMarketingAuthentication{
				SourceFacebookMarketingAuthenticateViaFacebookMarketingOauth: sourceFacebookMarketingAuthenticateViaFacebookMarketingOauth,
			}
		}
		var sourceFacebookMarketingServiceAccountKeyAuthentication *shared.SourceFacebookMarketingServiceAccountKeyAuthentication
		if r.Configuration.Credentials.ServiceAccountKeyAuthentication != nil {
			accessToken2 := r.Configuration.Credentials.ServiceAccountKeyAuthentication.AccessToken.ValueString()
			sourceFacebookMarketingServiceAccountKeyAuthentication = &shared.SourceFacebookMarketingServiceAccountKeyAuthentication{
				AccessToken: accessToken2,
			}
		}
		if sourceFacebookMarketingServiceAccountKeyAuthentication != nil {
			credentials = &shared.SourceFacebookMarketingAuthentication{
				SourceFacebookMarketingServiceAccountKeyAuthentication: sourceFacebookMarketingServiceAccountKeyAuthentication,
			}
		}
	}
	var customInsights []shared.SourceFacebookMarketingInsightConfig = []shared.SourceFacebookMarketingInsightConfig{}
	for _, customInsightsItem := range r.Configuration.CustomInsights {
		var actionBreakdowns []shared.SourceFacebookMarketingValidActionBreakdowns = []shared.SourceFacebookMarketingValidActionBreakdowns{}
		for _, actionBreakdownsItem := range customInsightsItem.ActionBreakdowns {
			actionBreakdowns = append(actionBreakdowns, shared.SourceFacebookMarketingValidActionBreakdowns(actionBreakdownsItem.ValueString()))
		}
		actionReportTime := new(shared.SourceFacebookMarketingActionReportTime)
		if !customInsightsItem.ActionReportTime.IsUnknown() && !customInsightsItem.ActionReportTime.IsNull() {
			*actionReportTime = shared.SourceFacebookMarketingActionReportTime(customInsightsItem.ActionReportTime.ValueString())
		} else {
			actionReportTime = nil
		}
		var breakdowns []shared.SourceFacebookMarketingValidBreakdowns = []shared.SourceFacebookMarketingValidBreakdowns{}
		for _, breakdownsItem := range customInsightsItem.Breakdowns {
			breakdowns = append(breakdowns, shared.SourceFacebookMarketingValidBreakdowns(breakdownsItem.ValueString()))
		}
		endDate := new(time.Time)
		if !customInsightsItem.EndDate.IsUnknown() && !customInsightsItem.EndDate.IsNull() {
			*endDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.EndDate.ValueString())
		} else {
			endDate = nil
		}
		var fields []shared.SourceFacebookMarketingValidEnums = []shared.SourceFacebookMarketingValidEnums{}
		for _, fieldsItem := range customInsightsItem.Fields {
			fields = append(fields, shared.SourceFacebookMarketingValidEnums(fieldsItem.ValueString()))
		}
		insightsJobTimeout := new(int64)
		if !customInsightsItem.InsightsJobTimeout.IsUnknown() && !customInsightsItem.InsightsJobTimeout.IsNull() {
			*insightsJobTimeout = customInsightsItem.InsightsJobTimeout.ValueInt64()
		} else {
			insightsJobTimeout = nil
		}
		insightsLookbackWindow := new(int64)
		if !customInsightsItem.InsightsLookbackWindow.IsUnknown() && !customInsightsItem.InsightsLookbackWindow.IsNull() {
			*insightsLookbackWindow = customInsightsItem.InsightsLookbackWindow.ValueInt64()
		} else {
			insightsLookbackWindow = nil
		}
		level := new(shared.SourceFacebookMarketingLevel)
		if !customInsightsItem.Level.IsUnknown() && !customInsightsItem.Level.IsNull() {
			*level = shared.SourceFacebookMarketingLevel(customInsightsItem.Level.ValueString())
		} else {
			level = nil
		}
		name := customInsightsItem.Name.ValueString()
		startDate := new(time.Time)
		if !customInsightsItem.StartDate.IsUnknown() && !customInsightsItem.StartDate.IsNull() {
			*startDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.StartDate.ValueString())
		} else {
			startDate = nil
		}
		timeIncrement := new(int64)
		if !customInsightsItem.TimeIncrement.IsUnknown() && !customInsightsItem.TimeIncrement.IsNull() {
			*timeIncrement = customInsightsItem.TimeIncrement.ValueInt64()
		} else {
			timeIncrement = nil
		}
		customInsights = append(customInsights, shared.SourceFacebookMarketingInsightConfig{
			ActionBreakdowns:       actionBreakdowns,
			ActionReportTime:       actionReportTime,
			Breakdowns:             breakdowns,
			EndDate:                endDate,
			Fields:                 fields,
			InsightsJobTimeout:     insightsJobTimeout,
			InsightsLookbackWindow: insightsLookbackWindow,
			Level:                  level,
			Name:                   name,
			StartDate:              startDate,
			TimeIncrement:          timeIncrement,
		})
	}
	endDate1 := new(time.Time)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	} else {
		endDate1 = nil
	}
	fetchThumbnailImages := new(bool)
	if !r.Configuration.FetchThumbnailImages.IsUnknown() && !r.Configuration.FetchThumbnailImages.IsNull() {
		*fetchThumbnailImages = r.Configuration.FetchThumbnailImages.ValueBool()
	} else {
		fetchThumbnailImages = nil
	}
	insightsJobTimeout1 := new(int64)
	if !r.Configuration.InsightsJobTimeout.IsUnknown() && !r.Configuration.InsightsJobTimeout.IsNull() {
		*insightsJobTimeout1 = r.Configuration.InsightsJobTimeout.ValueInt64()
	} else {
		insightsJobTimeout1 = nil
	}
	insightsLookbackWindow1 := new(int64)
	if !r.Configuration.InsightsLookbackWindow.IsUnknown() && !r.Configuration.InsightsLookbackWindow.IsNull() {
		*insightsLookbackWindow1 = r.Configuration.InsightsLookbackWindow.ValueInt64()
	} else {
		insightsLookbackWindow1 = nil
	}
	pageSize := new(int64)
	if !r.Configuration.PageSize.IsUnknown() && !r.Configuration.PageSize.IsNull() {
		*pageSize = r.Configuration.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	startDate1 := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate1 = nil
	}
	configuration := shared.SourceFacebookMarketing{
		AccessToken:                accessToken,
		AccountIds:                 accountIds,
		ActionBreakdownsAllowEmpty: actionBreakdownsAllowEmpty,
		AdStatuses:                 adStatuses,
		AdsetStatuses:              adsetStatuses,
		CampaignStatuses:           campaignStatuses,
		ClientID:                   clientID,
		ClientSecret:               clientSecret,
		Credentials:                credentials,
		CustomInsights:             customInsights,
		EndDate:                    endDate1,
		FetchThumbnailImages:       fetchThumbnailImages,
		InsightsJobTimeout:         insightsJobTimeout1,
		InsightsLookbackWindow:     insightsLookbackWindow1,
		PageSize:                   pageSize,
		StartDate:                  startDate1,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name1 := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFacebookMarketingCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name1,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFacebookMarketingResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceFacebookMarketingResourceModel) ToSharedSourceFacebookMarketingPutRequest() *shared.SourceFacebookMarketingPutRequest {
	accessToken := new(string)
	if !r.Configuration.AccessToken.IsUnknown() && !r.Configuration.AccessToken.IsNull() {
		*accessToken = r.Configuration.AccessToken.ValueString()
	} else {
		accessToken = nil
	}
	var accountIds []string = []string{}
	for _, accountIdsItem := range r.Configuration.AccountIds {
		accountIds = append(accountIds, accountIdsItem.ValueString())
	}
	actionBreakdownsAllowEmpty := new(bool)
	if !r.Configuration.ActionBreakdownsAllowEmpty.IsUnknown() && !r.Configuration.ActionBreakdownsAllowEmpty.IsNull() {
		*actionBreakdownsAllowEmpty = r.Configuration.ActionBreakdownsAllowEmpty.ValueBool()
	} else {
		actionBreakdownsAllowEmpty = nil
	}
	var adStatuses []shared.ValidAdStatuses = []shared.ValidAdStatuses{}
	for _, adStatusesItem := range r.Configuration.AdStatuses {
		adStatuses = append(adStatuses, shared.ValidAdStatuses(adStatusesItem.ValueString()))
	}
	var adsetStatuses []shared.ValidAdSetStatuses = []shared.ValidAdSetStatuses{}
	for _, adsetStatusesItem := range r.Configuration.AdsetStatuses {
		adsetStatuses = append(adsetStatuses, shared.ValidAdSetStatuses(adsetStatusesItem.ValueString()))
	}
	var campaignStatuses []shared.ValidCampaignStatuses = []shared.ValidCampaignStatuses{}
	for _, campaignStatusesItem := range r.Configuration.CampaignStatuses {
		campaignStatuses = append(campaignStatuses, shared.ValidCampaignStatuses(campaignStatusesItem.ValueString()))
	}
	clientID := new(string)
	if !r.Configuration.ClientID.IsUnknown() && !r.Configuration.ClientID.IsNull() {
		*clientID = r.Configuration.ClientID.ValueString()
	} else {
		clientID = nil
	}
	clientSecret := new(string)
	if !r.Configuration.ClientSecret.IsUnknown() && !r.Configuration.ClientSecret.IsNull() {
		*clientSecret = r.Configuration.ClientSecret.ValueString()
	} else {
		clientSecret = nil
	}
	var credentials *shared.SourceFacebookMarketingUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var authenticateViaFacebookMarketingOauth *shared.AuthenticateViaFacebookMarketingOauth
		if r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth != nil {
			accessToken1 := new(string)
			if !r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth.AccessToken.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth.AccessToken.IsNull() {
				*accessToken1 = r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth.AccessToken.ValueString()
			} else {
				accessToken1 = nil
			}
			clientId1 := r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth.ClientID.ValueString()
			clientSecret1 := r.Configuration.Credentials.AuthenticateViaFacebookMarketingOauth.ClientSecret.ValueString()
			authenticateViaFacebookMarketingOauth = &shared.AuthenticateViaFacebookMarketingOauth{
				AccessToken:  accessToken1,
				ClientID:     clientId1,
				ClientSecret: clientSecret1,
			}
		}
		if authenticateViaFacebookMarketingOauth != nil {
			credentials = &shared.SourceFacebookMarketingUpdateAuthentication{
				AuthenticateViaFacebookMarketingOauth: authenticateViaFacebookMarketingOauth,
			}
		}
		var serviceAccountKeyAuthentication *shared.ServiceAccountKeyAuthentication
		if r.Configuration.Credentials.ServiceAccountKeyAuthentication != nil {
			accessToken2 := r.Configuration.Credentials.ServiceAccountKeyAuthentication.AccessToken.ValueString()
			serviceAccountKeyAuthentication = &shared.ServiceAccountKeyAuthentication{
				AccessToken: accessToken2,
			}
		}
		if serviceAccountKeyAuthentication != nil {
			credentials = &shared.SourceFacebookMarketingUpdateAuthentication{
				ServiceAccountKeyAuthentication: serviceAccountKeyAuthentication,
			}
		}
	}
	var customInsights []shared.InsightConfig = []shared.InsightConfig{}
	for _, customInsightsItem := range r.Configuration.CustomInsights {
		var actionBreakdowns []shared.ValidActionBreakdowns = []shared.ValidActionBreakdowns{}
		for _, actionBreakdownsItem := range customInsightsItem.ActionBreakdowns {
			actionBreakdowns = append(actionBreakdowns, shared.ValidActionBreakdowns(actionBreakdownsItem.ValueString()))
		}
		actionReportTime := new(shared.SourceFacebookMarketingUpdateActionReportTime)
		if !customInsightsItem.ActionReportTime.IsUnknown() && !customInsightsItem.ActionReportTime.IsNull() {
			*actionReportTime = shared.SourceFacebookMarketingUpdateActionReportTime(customInsightsItem.ActionReportTime.ValueString())
		} else {
			actionReportTime = nil
		}
		var breakdowns []shared.ValidBreakdowns = []shared.ValidBreakdowns{}
		for _, breakdownsItem := range customInsightsItem.Breakdowns {
			breakdowns = append(breakdowns, shared.ValidBreakdowns(breakdownsItem.ValueString()))
		}
		endDate := new(time.Time)
		if !customInsightsItem.EndDate.IsUnknown() && !customInsightsItem.EndDate.IsNull() {
			*endDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.EndDate.ValueString())
		} else {
			endDate = nil
		}
		var fields []shared.SourceFacebookMarketingUpdateValidEnums = []shared.SourceFacebookMarketingUpdateValidEnums{}
		for _, fieldsItem := range customInsightsItem.Fields {
			fields = append(fields, shared.SourceFacebookMarketingUpdateValidEnums(fieldsItem.ValueString()))
		}
		insightsJobTimeout := new(int64)
		if !customInsightsItem.InsightsJobTimeout.IsUnknown() && !customInsightsItem.InsightsJobTimeout.IsNull() {
			*insightsJobTimeout = customInsightsItem.InsightsJobTimeout.ValueInt64()
		} else {
			insightsJobTimeout = nil
		}
		insightsLookbackWindow := new(int64)
		if !customInsightsItem.InsightsLookbackWindow.IsUnknown() && !customInsightsItem.InsightsLookbackWindow.IsNull() {
			*insightsLookbackWindow = customInsightsItem.InsightsLookbackWindow.ValueInt64()
		} else {
			insightsLookbackWindow = nil
		}
		level := new(shared.Level)
		if !customInsightsItem.Level.IsUnknown() && !customInsightsItem.Level.IsNull() {
			*level = shared.Level(customInsightsItem.Level.ValueString())
		} else {
			level = nil
		}
		name := customInsightsItem.Name.ValueString()
		startDate := new(time.Time)
		if !customInsightsItem.StartDate.IsUnknown() && !customInsightsItem.StartDate.IsNull() {
			*startDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.StartDate.ValueString())
		} else {
			startDate = nil
		}
		timeIncrement := new(int64)
		if !customInsightsItem.TimeIncrement.IsUnknown() && !customInsightsItem.TimeIncrement.IsNull() {
			*timeIncrement = customInsightsItem.TimeIncrement.ValueInt64()
		} else {
			timeIncrement = nil
		}
		customInsights = append(customInsights, shared.InsightConfig{
			ActionBreakdowns:       actionBreakdowns,
			ActionReportTime:       actionReportTime,
			Breakdowns:             breakdowns,
			EndDate:                endDate,
			Fields:                 fields,
			InsightsJobTimeout:     insightsJobTimeout,
			InsightsLookbackWindow: insightsLookbackWindow,
			Level:                  level,
			Name:                   name,
			StartDate:              startDate,
			TimeIncrement:          timeIncrement,
		})
	}
	endDate1 := new(time.Time)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	} else {
		endDate1 = nil
	}
	fetchThumbnailImages := new(bool)
	if !r.Configuration.FetchThumbnailImages.IsUnknown() && !r.Configuration.FetchThumbnailImages.IsNull() {
		*fetchThumbnailImages = r.Configuration.FetchThumbnailImages.ValueBool()
	} else {
		fetchThumbnailImages = nil
	}
	insightsJobTimeout1 := new(int64)
	if !r.Configuration.InsightsJobTimeout.IsUnknown() && !r.Configuration.InsightsJobTimeout.IsNull() {
		*insightsJobTimeout1 = r.Configuration.InsightsJobTimeout.ValueInt64()
	} else {
		insightsJobTimeout1 = nil
	}
	insightsLookbackWindow1 := new(int64)
	if !r.Configuration.InsightsLookbackWindow.IsUnknown() && !r.Configuration.InsightsLookbackWindow.IsNull() {
		*insightsLookbackWindow1 = r.Configuration.InsightsLookbackWindow.ValueInt64()
	} else {
		insightsLookbackWindow1 = nil
	}
	pageSize := new(int64)
	if !r.Configuration.PageSize.IsUnknown() && !r.Configuration.PageSize.IsNull() {
		*pageSize = r.Configuration.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	startDate1 := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate1 = nil
	}
	configuration := shared.SourceFacebookMarketingUpdate{
		AccessToken:                accessToken,
		AccountIds:                 accountIds,
		ActionBreakdownsAllowEmpty: actionBreakdownsAllowEmpty,
		AdStatuses:                 adStatuses,
		AdsetStatuses:              adsetStatuses,
		CampaignStatuses:           campaignStatuses,
		ClientID:                   clientID,
		ClientSecret:               clientSecret,
		Credentials:                credentials,
		CustomInsights:             customInsights,
		EndDate:                    endDate1,
		FetchThumbnailImages:       fetchThumbnailImages,
		InsightsJobTimeout:         insightsJobTimeout1,
		InsightsLookbackWindow:     insightsLookbackWindow1,
		PageSize:                   pageSize,
		StartDate:                  startDate1,
	}
	name1 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFacebookMarketingPutRequest{
		Configuration: configuration,
		Name:          name1,
		WorkspaceID:   workspaceID,
	}
	return &out
}
