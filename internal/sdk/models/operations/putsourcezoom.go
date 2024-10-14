// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceZoomRequest struct {
	SourceZoomPutRequest *shared.SourceZoomPutRequest `request:"mediaType=application/json"`
	SourceID             string                       `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceZoomRequest) GetSourceZoomPutRequest() *shared.SourceZoomPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceZoomPutRequest
}

func (o *PutSourceZoomRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceZoomResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceZoomResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceZoomResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceZoomResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
