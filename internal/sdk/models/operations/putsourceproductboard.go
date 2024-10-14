// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceProductboardRequest struct {
	SourceProductboardPutRequest *shared.SourceProductboardPutRequest `request:"mediaType=application/json"`
	SourceID                     string                               `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceProductboardRequest) GetSourceProductboardPutRequest() *shared.SourceProductboardPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceProductboardPutRequest
}

func (o *PutSourceProductboardRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceProductboardResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceProductboardResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceProductboardResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceProductboardResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
