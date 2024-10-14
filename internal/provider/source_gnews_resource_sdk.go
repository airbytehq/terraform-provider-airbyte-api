// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGnewsResourceModel) ToSharedSourceGnewsCreateRequest() *shared.SourceGnewsCreateRequest {
	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	country := new(shared.SourceGnewsCountry)
	if !r.Configuration.Country.IsUnknown() && !r.Configuration.Country.IsNull() {
		*country = shared.SourceGnewsCountry(r.Configuration.Country.ValueString())
	} else {
		country = nil
	}
	endDate := new(string)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueString()
	} else {
		endDate = nil
	}
	var in []shared.SourceGnewsIn = []shared.SourceGnewsIn{}
	for _, inItem := range r.Configuration.In {
		in = append(in, shared.SourceGnewsIn(inItem.ValueString()))
	}
	language := new(shared.SourceGnewsLanguage)
	if !r.Configuration.Language.IsUnknown() && !r.Configuration.Language.IsNull() {
		*language = shared.SourceGnewsLanguage(r.Configuration.Language.ValueString())
	} else {
		language = nil
	}
	var nullable []shared.SourceGnewsNullable = []shared.SourceGnewsNullable{}
	for _, nullableItem := range r.Configuration.Nullable {
		nullable = append(nullable, shared.SourceGnewsNullable(nullableItem.ValueString()))
	}
	var query string
	query = r.Configuration.Query.ValueString()

	sortby := new(shared.SourceGnewsSortBy)
	if !r.Configuration.Sortby.IsUnknown() && !r.Configuration.Sortby.IsNull() {
		*sortby = shared.SourceGnewsSortBy(r.Configuration.Sortby.ValueString())
	} else {
		sortby = nil
	}
	startDate := new(string)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueString()
	} else {
		startDate = nil
	}
	topHeadlinesQuery := new(string)
	if !r.Configuration.TopHeadlinesQuery.IsUnknown() && !r.Configuration.TopHeadlinesQuery.IsNull() {
		*topHeadlinesQuery = r.Configuration.TopHeadlinesQuery.ValueString()
	} else {
		topHeadlinesQuery = nil
	}
	topHeadlinesTopic := new(shared.SourceGnewsTopHeadlinesTopic)
	if !r.Configuration.TopHeadlinesTopic.IsUnknown() && !r.Configuration.TopHeadlinesTopic.IsNull() {
		*topHeadlinesTopic = shared.SourceGnewsTopHeadlinesTopic(r.Configuration.TopHeadlinesTopic.ValueString())
	} else {
		topHeadlinesTopic = nil
	}
	configuration := shared.SourceGnews{
		APIKey:            apiKey,
		Country:           country,
		EndDate:           endDate,
		In:                in,
		Language:          language,
		Nullable:          nullable,
		Query:             query,
		Sortby:            sortby,
		StartDate:         startDate,
		TopHeadlinesQuery: topHeadlinesQuery,
		TopHeadlinesTopic: topHeadlinesTopic,
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

	out := shared.SourceGnewsCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGnewsResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceGnewsResourceModel) ToSharedSourceGnewsPutRequest() *shared.SourceGnewsPutRequest {
	var apiKey string
	apiKey = r.Configuration.APIKey.ValueString()

	country := new(shared.Country)
	if !r.Configuration.Country.IsUnknown() && !r.Configuration.Country.IsNull() {
		*country = shared.Country(r.Configuration.Country.ValueString())
	} else {
		country = nil
	}
	endDate := new(string)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueString()
	} else {
		endDate = nil
	}
	var in []shared.In = []shared.In{}
	for _, inItem := range r.Configuration.In {
		in = append(in, shared.In(inItem.ValueString()))
	}
	language := new(shared.Language)
	if !r.Configuration.Language.IsUnknown() && !r.Configuration.Language.IsNull() {
		*language = shared.Language(r.Configuration.Language.ValueString())
	} else {
		language = nil
	}
	var nullable []shared.Nullable = []shared.Nullable{}
	for _, nullableItem := range r.Configuration.Nullable {
		nullable = append(nullable, shared.Nullable(nullableItem.ValueString()))
	}
	var query string
	query = r.Configuration.Query.ValueString()

	sortby := new(shared.SortBy)
	if !r.Configuration.Sortby.IsUnknown() && !r.Configuration.Sortby.IsNull() {
		*sortby = shared.SortBy(r.Configuration.Sortby.ValueString())
	} else {
		sortby = nil
	}
	startDate := new(string)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueString()
	} else {
		startDate = nil
	}
	topHeadlinesQuery := new(string)
	if !r.Configuration.TopHeadlinesQuery.IsUnknown() && !r.Configuration.TopHeadlinesQuery.IsNull() {
		*topHeadlinesQuery = r.Configuration.TopHeadlinesQuery.ValueString()
	} else {
		topHeadlinesQuery = nil
	}
	topHeadlinesTopic := new(shared.TopHeadlinesTopic)
	if !r.Configuration.TopHeadlinesTopic.IsUnknown() && !r.Configuration.TopHeadlinesTopic.IsNull() {
		*topHeadlinesTopic = shared.TopHeadlinesTopic(r.Configuration.TopHeadlinesTopic.ValueString())
	} else {
		topHeadlinesTopic = nil
	}
	configuration := shared.SourceGnewsUpdate{
		APIKey:            apiKey,
		Country:           country,
		EndDate:           endDate,
		In:                in,
		Language:          language,
		Nullable:          nullable,
		Query:             query,
		Sortby:            sortby,
		StartDate:         startDate,
		TopHeadlinesQuery: topHeadlinesQuery,
		TopHeadlinesTopic: topHeadlinesTopic,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.SourceGnewsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
