// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceCustomRequest struct {
	SourceCustomPutRequest *shared.SourceCustomPutRequest `request:"mediaType=application/json"`
	SourceID               string                         `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceCustomRequest) GetSourceCustomPutRequest() *shared.SourceCustomPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceCustomPutRequest
}

func (o *PutSourceCustomRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceCustomResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceCustomResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceCustomResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceCustomResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
