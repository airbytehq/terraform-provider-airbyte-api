// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutSourceTvmazeScheduleRequest struct {
	SourceTvmazeSchedulePutRequest *shared.SourceTvmazeSchedulePutRequest `request:"mediaType=application/json"`
	SourceID                       string                                 `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceTvmazeScheduleRequest) GetSourceTvmazeSchedulePutRequest() *shared.SourceTvmazeSchedulePutRequest {
	if o == nil {
		return nil
	}
	return o.SourceTvmazeSchedulePutRequest
}

func (o *PutSourceTvmazeScheduleRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceTvmazeScheduleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceTvmazeScheduleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceTvmazeScheduleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceTvmazeScheduleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
