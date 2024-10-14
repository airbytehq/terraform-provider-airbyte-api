// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	customTypes "github.com/airbytehq/terraform-provider-airbyte/internal/sdk/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePolygonStockAPIResourceModel) ToSharedSourcePolygonStockAPICreateRequest() *shared.SourcePolygonStockAPICreateRequest {
	adjusted := new(string)
	if !r.Configuration.Adjusted.IsUnknown() && !r.Configuration.Adjusted.IsNull() {
		*adjusted = r.Configuration.Adjusted.ValueString()
	} else {
		adjusted = nil
	}
	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	endDate := customTypes.MustDateFromString(r.Configuration.EndDate.ValueString())
	limit := new(int64)
	if !r.Configuration.Limit.IsUnknown() && !r.Configuration.Limit.IsNull() {
		*limit = r.Configuration.Limit.ValueInt64()
	} else {
		limit = nil
	}
	var multiplier int64
	multiplier = r.Configuration.Multiplier.ValueInt64()

	sort := new(string)
	if !r.Configuration.Sort.IsUnknown() && !r.Configuration.Sort.IsNull() {
		*sort = r.Configuration.Sort.ValueString()
	} else {
		sort = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	var stocksTicker string
	stocksTicker = r.Configuration.StocksTicker.ValueString()

	var timespan string
	timespan = r.Configuration.Timespan.ValueString()

	configuration := shared.SourcePolygonStockAPI{
		Adjusted:     adjusted,
		APIKey:       apiKey,
		EndDate:      endDate,
		Limit:        limit,
		Multiplier:   multiplier,
		Sort:         sort,
		StartDate:    startDate,
		StocksTicker: stocksTicker,
		Timespan:     timespan,
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

	out := shared.SourcePolygonStockAPICreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePolygonStockAPIResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourcePolygonStockAPIResourceModel) ToSharedSourcePolygonStockAPIPutRequest() *shared.SourcePolygonStockAPIPutRequest {
	adjusted := new(string)
	if !r.Configuration.Adjusted.IsUnknown() && !r.Configuration.Adjusted.IsNull() {
		*adjusted = r.Configuration.Adjusted.ValueString()
	} else {
		adjusted = nil
	}
	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	endDate := customTypes.MustDateFromString(r.Configuration.EndDate.ValueString())
	limit := new(int64)
	if !r.Configuration.Limit.IsUnknown() && !r.Configuration.Limit.IsNull() {
		*limit = r.Configuration.Limit.ValueInt64()
	} else {
		limit = nil
	}
	var multiplier int64
	multiplier = r.Configuration.Multiplier.ValueInt64()

	sort := new(string)
	if !r.Configuration.Sort.IsUnknown() && !r.Configuration.Sort.IsNull() {
		*sort = r.Configuration.Sort.ValueString()
	} else {
		sort = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	var stocksTicker string
	stocksTicker = r.Configuration.StocksTicker.ValueString()

	var timespan string
	timespan = r.Configuration.Timespan.ValueString()

	configuration := shared.SourcePolygonStockAPIUpdate{
		Adjusted:     adjusted,
		APIKey:       apiKey,
		EndDate:      endDate,
		Limit:        limit,
		Multiplier:   multiplier,
		Sort:         sort,
		StartDate:    startDate,
		StocksTicker: stocksTicker,
		Timespan:     timespan,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourcePolygonStockAPIPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
