// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceChargebeeResourceModel) ToSharedSourceChargebeeCreateRequest() *shared.SourceChargebeeCreateRequest {
	productCatalog := new(shared.SourceChargebeeProductCatalog)
	if !r.Configuration.ProductCatalog.IsUnknown() && !r.Configuration.ProductCatalog.IsNull() {
		*productCatalog = shared.SourceChargebeeProductCatalog(r.Configuration.ProductCatalog.ValueString())
	} else {
		productCatalog = nil
	}
	site := r.Configuration.Site.ValueString()
	siteAPIKey := r.Configuration.SiteAPIKey.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceChargebee{
		ProductCatalog: productCatalog,
		Site:           site,
		SiteAPIKey:     siteAPIKey,
		StartDate:      startDate,
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
	out := shared.SourceChargebeeCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceChargebeeResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceChargebeeResourceModel) ToSharedSourceChargebeePutRequest() *shared.SourceChargebeePutRequest {
	productCatalog := new(shared.ProductCatalog)
	if !r.Configuration.ProductCatalog.IsUnknown() && !r.Configuration.ProductCatalog.IsNull() {
		*productCatalog = shared.ProductCatalog(r.Configuration.ProductCatalog.ValueString())
	} else {
		productCatalog = nil
	}
	site := r.Configuration.Site.ValueString()
	siteAPIKey := r.Configuration.SiteAPIKey.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceChargebeeUpdate{
		ProductCatalog: productCatalog,
		Site:           site,
		SiteAPIKey:     siteAPIKey,
		StartDate:      startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceChargebeePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
