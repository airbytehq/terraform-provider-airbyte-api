// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutDestinationAstraRequest struct {
	DestinationAstraPutRequest *shared.DestinationAstraPutRequest `request:"mediaType=application/json"`
	DestinationID              string                             `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationAstraRequest) GetDestinationAstraPutRequest() *shared.DestinationAstraPutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationAstraPutRequest
}

func (o *PutDestinationAstraRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationAstraResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationAstraResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationAstraResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationAstraResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
