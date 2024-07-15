// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	customTypes "github.com/airbytehq/terraform-provider-airbyte/internal/sdk/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceShopifyResourceModel) ToSharedSourceShopifyCreateRequest() *shared.SourceShopifyCreateRequest {
	bulkWindowInDays := new(int64)
	if !r.Configuration.BulkWindowInDays.IsUnknown() && !r.Configuration.BulkWindowInDays.IsNull() {
		*bulkWindowInDays = r.Configuration.BulkWindowInDays.ValueInt64()
	} else {
		bulkWindowInDays = nil
	}
	var credentials *shared.SourceShopifyShopifyAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceShopifyOAuth20 *shared.SourceShopifyOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := new(string)
			if !r.Configuration.Credentials.OAuth20.AccessToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.AccessToken.IsNull() {
				*accessToken = r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			} else {
				accessToken = nil
			}
			clientID := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			sourceShopifyOAuth20 = &shared.SourceShopifyOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
		}
		if sourceShopifyOAuth20 != nil {
			credentials = &shared.SourceShopifyShopifyAuthorizationMethod{
				SourceShopifyOAuth20: sourceShopifyOAuth20,
			}
		}
		var sourceShopifyAPIPassword *shared.SourceShopifyAPIPassword
		if r.Configuration.Credentials.APIPassword != nil {
			apiPassword := r.Configuration.Credentials.APIPassword.APIPassword.ValueString()
			sourceShopifyAPIPassword = &shared.SourceShopifyAPIPassword{
				APIPassword: apiPassword,
			}
		}
		if sourceShopifyAPIPassword != nil {
			credentials = &shared.SourceShopifyShopifyAuthorizationMethod{
				SourceShopifyAPIPassword: sourceShopifyAPIPassword,
			}
		}
	}
	fetchTransactionsUserID := new(bool)
	if !r.Configuration.FetchTransactionsUserID.IsUnknown() && !r.Configuration.FetchTransactionsUserID.IsNull() {
		*fetchTransactionsUserID = r.Configuration.FetchTransactionsUserID.ValueBool()
	} else {
		fetchTransactionsUserID = nil
	}
	jobTerminationThreshold := new(int64)
	if !r.Configuration.JobTerminationThreshold.IsUnknown() && !r.Configuration.JobTerminationThreshold.IsNull() {
		*jobTerminationThreshold = r.Configuration.JobTerminationThreshold.ValueInt64()
	} else {
		jobTerminationThreshold = nil
	}
	shop := r.Configuration.Shop.ValueString()
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceShopify{
		BulkWindowInDays:        bulkWindowInDays,
		Credentials:             credentials,
		FetchTransactionsUserID: fetchTransactionsUserID,
		JobTerminationThreshold: jobTerminationThreshold,
		Shop:                    shop,
		StartDate:               startDate,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceShopifyCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceShopifyResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceShopifyResourceModel) ToSharedSourceShopifyPutRequest() *shared.SourceShopifyPutRequest {
	bulkWindowInDays := new(int64)
	if !r.Configuration.BulkWindowInDays.IsUnknown() && !r.Configuration.BulkWindowInDays.IsNull() {
		*bulkWindowInDays = r.Configuration.BulkWindowInDays.ValueInt64()
	} else {
		bulkWindowInDays = nil
	}
	var credentials *shared.ShopifyAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceShopifyUpdateOAuth20 *shared.SourceShopifyUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := new(string)
			if !r.Configuration.Credentials.OAuth20.AccessToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.AccessToken.IsNull() {
				*accessToken = r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			} else {
				accessToken = nil
			}
			clientID := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			sourceShopifyUpdateOAuth20 = &shared.SourceShopifyUpdateOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
		}
		if sourceShopifyUpdateOAuth20 != nil {
			credentials = &shared.ShopifyAuthorizationMethod{
				SourceShopifyUpdateOAuth20: sourceShopifyUpdateOAuth20,
			}
		}
		var apiPassword *shared.APIPassword
		if r.Configuration.Credentials.APIPassword != nil {
			apiPassword1 := r.Configuration.Credentials.APIPassword.APIPassword.ValueString()
			apiPassword = &shared.APIPassword{
				APIPassword: apiPassword1,
			}
		}
		if apiPassword != nil {
			credentials = &shared.ShopifyAuthorizationMethod{
				APIPassword: apiPassword,
			}
		}
	}
	fetchTransactionsUserID := new(bool)
	if !r.Configuration.FetchTransactionsUserID.IsUnknown() && !r.Configuration.FetchTransactionsUserID.IsNull() {
		*fetchTransactionsUserID = r.Configuration.FetchTransactionsUserID.ValueBool()
	} else {
		fetchTransactionsUserID = nil
	}
	jobTerminationThreshold := new(int64)
	if !r.Configuration.JobTerminationThreshold.IsUnknown() && !r.Configuration.JobTerminationThreshold.IsNull() {
		*jobTerminationThreshold = r.Configuration.JobTerminationThreshold.ValueInt64()
	} else {
		jobTerminationThreshold = nil
	}
	shop := r.Configuration.Shop.ValueString()
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceShopifyUpdate{
		BulkWindowInDays:        bulkWindowInDays,
		Credentials:             credentials,
		FetchTransactionsUserID: fetchTransactionsUserID,
		JobTerminationThreshold: jobTerminationThreshold,
		Shop:                    shop,
		StartDate:               startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceShopifyPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
