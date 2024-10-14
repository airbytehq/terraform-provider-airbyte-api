// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type GetDestinationRedshiftRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *GetDestinationRedshiftRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type GetDestinationRedshiftResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Get a Destination by the id in the path.
	DestinationResponse *shared.DestinationResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDestinationRedshiftResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDestinationRedshiftResponse) GetDestinationResponse() *shared.DestinationResponse {
	if o == nil {
		return nil
	}
	return o.DestinationResponse
}

func (o *GetDestinationRedshiftResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDestinationRedshiftResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
