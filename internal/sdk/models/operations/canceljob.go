// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type CancelJobRequest struct {
	JobID int64 `pathParam:"style=simple,explode=false,name=jobId"`
}

func (o *CancelJobRequest) GetJobID() int64 {
	if o == nil {
		return 0
	}
	return o.JobID
}

type CancelJobResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Cancel a Job.
	JobResponse *shared.JobResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CancelJobResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CancelJobResponse) GetJobResponse() *shared.JobResponse {
	if o == nil {
		return nil
	}
	return o.JobResponse
}

func (o *CancelJobResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CancelJobResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}