// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type ListDestinationsRequest struct {
	// Include deleted destinations in the returned results.
	IncludeDeleted *bool `default:"false" queryParam:"style=form,explode=true,name=includeDeleted"`
	// Set the limit on the number of destinations returned. The default is 20.
	Limit *int `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Set the offset to start at when returning destinations. The default is 0
	Offset *int `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// The UUIDs of the workspaces you wish to list destinations for. Empty list will retrieve all allowed workspaces.
	WorkspaceIds []string `queryParam:"style=form,explode=true,name=workspaceIds"`
}

func (l ListDestinationsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListDestinationsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListDestinationsRequest) GetIncludeDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeleted
}

func (o *ListDestinationsRequest) GetLimit() *int {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListDestinationsRequest) GetOffset() *int {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListDestinationsRequest) GetWorkspaceIds() []string {
	if o == nil {
		return nil
	}
	return o.WorkspaceIds
}

type ListDestinationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	DestinationsResponse *shared.DestinationsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListDestinationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDestinationsResponse) GetDestinationsResponse() *shared.DestinationsResponse {
	if o == nil {
		return nil
	}
	return o.DestinationsResponse
}

func (o *ListDestinationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDestinationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
