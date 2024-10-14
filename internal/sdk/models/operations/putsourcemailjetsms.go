// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceMailjetSmsRequest struct {
	SourceMailjetSmsPutRequest *shared.SourceMailjetSmsPutRequest `request:"mediaType=application/json"`
	SourceID                   string                             `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceMailjetSmsRequest) GetSourceMailjetSmsPutRequest() *shared.SourceMailjetSmsPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceMailjetSmsPutRequest
}

func (o *PutSourceMailjetSmsRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceMailjetSmsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceMailjetSmsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceMailjetSmsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceMailjetSmsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
