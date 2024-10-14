// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationQdrantDataSourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	if resp != nil {
		configurationResult, _ := json.Marshal(resp.Configuration)
		r.Configuration = types.StringValue(string(configurationResult))
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.DestinationType = types.StringValue(resp.DestinationType)
		r.Name = types.StringValue(resp.Name)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}
