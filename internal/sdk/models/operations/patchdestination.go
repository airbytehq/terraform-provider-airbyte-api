// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PatchDestinationRequest struct {
	DestinationPatchRequest *shared.DestinationPatchRequest `request:"mediaType=application/json"`
	DestinationID           string                          `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PatchDestinationRequest) GetDestinationPatchRequest() *shared.DestinationPatchRequest {
	if o == nil {
		return nil
	}
	return o.DestinationPatchRequest
}

func (o *PatchDestinationRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PatchDestinationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Update a Destination
	DestinationResponse *shared.DestinationResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchDestinationResponse) GetDestinationResponse() *shared.DestinationResponse {
	if o == nil {
		return nil
	}
	return o.DestinationResponse
}

func (o *PatchDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
