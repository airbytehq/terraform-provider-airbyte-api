// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePocketResourceModel) ToSharedSourcePocketCreateRequest() *shared.SourcePocketCreateRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	consumerKey := r.Configuration.ConsumerKey.ValueString()
	contentType := new(shared.SourcePocketContentType)
	if !r.Configuration.ContentType.IsUnknown() && !r.Configuration.ContentType.IsNull() {
		*contentType = shared.SourcePocketContentType(r.Configuration.ContentType.ValueString())
	} else {
		contentType = nil
	}
	detailType := new(shared.SourcePocketDetailType)
	if !r.Configuration.DetailType.IsUnknown() && !r.Configuration.DetailType.IsNull() {
		*detailType = shared.SourcePocketDetailType(r.Configuration.DetailType.ValueString())
	} else {
		detailType = nil
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	favorite := new(bool)
	if !r.Configuration.Favorite.IsUnknown() && !r.Configuration.Favorite.IsNull() {
		*favorite = r.Configuration.Favorite.ValueBool()
	} else {
		favorite = nil
	}
	search := new(string)
	if !r.Configuration.Search.IsUnknown() && !r.Configuration.Search.IsNull() {
		*search = r.Configuration.Search.ValueString()
	} else {
		search = nil
	}
	since := new(string)
	if !r.Configuration.Since.IsUnknown() && !r.Configuration.Since.IsNull() {
		*since = r.Configuration.Since.ValueString()
	} else {
		since = nil
	}
	sort := new(shared.SourcePocketSortBy)
	if !r.Configuration.Sort.IsUnknown() && !r.Configuration.Sort.IsNull() {
		*sort = shared.SourcePocketSortBy(r.Configuration.Sort.ValueString())
	} else {
		sort = nil
	}
	state := new(shared.SourcePocketState)
	if !r.Configuration.State.IsUnknown() && !r.Configuration.State.IsNull() {
		*state = shared.SourcePocketState(r.Configuration.State.ValueString())
	} else {
		state = nil
	}
	tag := new(string)
	if !r.Configuration.Tag.IsUnknown() && !r.Configuration.Tag.IsNull() {
		*tag = r.Configuration.Tag.ValueString()
	} else {
		tag = nil
	}
	configuration := shared.SourcePocket{
		AccessToken: accessToken,
		ConsumerKey: consumerKey,
		ContentType: contentType,
		DetailType:  detailType,
		Domain:      domain,
		Favorite:    favorite,
		Search:      search,
		Since:       since,
		Sort:        sort,
		State:       state,
		Tag:         tag,
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
	out := shared.SourcePocketCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePocketResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourcePocketResourceModel) ToSharedSourcePocketPutRequest() *shared.SourcePocketPutRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	consumerKey := r.Configuration.ConsumerKey.ValueString()
	contentType := new(shared.ContentType)
	if !r.Configuration.ContentType.IsUnknown() && !r.Configuration.ContentType.IsNull() {
		*contentType = shared.ContentType(r.Configuration.ContentType.ValueString())
	} else {
		contentType = nil
	}
	detailType := new(shared.DetailType)
	if !r.Configuration.DetailType.IsUnknown() && !r.Configuration.DetailType.IsNull() {
		*detailType = shared.DetailType(r.Configuration.DetailType.ValueString())
	} else {
		detailType = nil
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	favorite := new(bool)
	if !r.Configuration.Favorite.IsUnknown() && !r.Configuration.Favorite.IsNull() {
		*favorite = r.Configuration.Favorite.ValueBool()
	} else {
		favorite = nil
	}
	search := new(string)
	if !r.Configuration.Search.IsUnknown() && !r.Configuration.Search.IsNull() {
		*search = r.Configuration.Search.ValueString()
	} else {
		search = nil
	}
	since := new(string)
	if !r.Configuration.Since.IsUnknown() && !r.Configuration.Since.IsNull() {
		*since = r.Configuration.Since.ValueString()
	} else {
		since = nil
	}
	sort := new(shared.SourcePocketUpdateSortBy)
	if !r.Configuration.Sort.IsUnknown() && !r.Configuration.Sort.IsNull() {
		*sort = shared.SourcePocketUpdateSortBy(r.Configuration.Sort.ValueString())
	} else {
		sort = nil
	}
	state := new(shared.State)
	if !r.Configuration.State.IsUnknown() && !r.Configuration.State.IsNull() {
		*state = shared.State(r.Configuration.State.ValueString())
	} else {
		state = nil
	}
	tag := new(string)
	if !r.Configuration.Tag.IsUnknown() && !r.Configuration.Tag.IsNull() {
		*tag = r.Configuration.Tag.ValueString()
	} else {
		tag = nil
	}
	configuration := shared.SourcePocketUpdate{
		AccessToken: accessToken,
		ConsumerKey: consumerKey,
		ContentType: contentType,
		DetailType:  detailType,
		Domain:      domain,
		Favorite:    favorite,
		Search:      search,
		Since:       since,
		Sort:        sort,
		State:       state,
		Tag:         tag,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePocketPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
