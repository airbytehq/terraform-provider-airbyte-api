// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutDestinationPgvectorRequest struct {
	DestinationPgvectorPutRequest *shared.DestinationPgvectorPutRequest `request:"mediaType=application/json"`
	DestinationID                 string                                `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationPgvectorRequest) GetDestinationPgvectorPutRequest() *shared.DestinationPgvectorPutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationPgvectorPutRequest
}

func (o *PutDestinationPgvectorRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationPgvectorResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationPgvectorResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationPgvectorResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationPgvectorResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
