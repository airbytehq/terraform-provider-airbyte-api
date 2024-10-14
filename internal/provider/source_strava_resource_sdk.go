// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceStravaResourceModel) ToSharedSourceStravaCreateRequest() *shared.SourceStravaCreateRequest {
	var athleteID int64
	athleteID = r.Configuration.AthleteID.ValueInt64()

	var clientID string
	clientID = r.Configuration.ClientID.ValueString()

	var clientSecret string
	clientSecret = r.Configuration.ClientSecret.ValueString()

	var refreshToken string
	refreshToken = r.Configuration.RefreshToken.ValueString()

	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceStrava{
		AthleteID:    athleteID,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
		StartDate:    startDate,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	var name string
	name = r.Name.ValueString()

	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceStravaCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceStravaResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceStravaResourceModel) ToSharedSourceStravaPutRequest() *shared.SourceStravaPutRequest {
	var athleteID int64
	athleteID = r.Configuration.AthleteID.ValueInt64()

	var clientID string
	clientID = r.Configuration.ClientID.ValueString()

	var clientSecret string
	clientSecret = r.Configuration.ClientSecret.ValueString()

	var refreshToken string
	refreshToken = r.Configuration.RefreshToken.ValueString()

	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceStravaUpdate{
		AthleteID:    athleteID,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
		StartDate:    startDate,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceStravaPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
