// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PatchSourceRequest struct {
	SourcePatchRequest *shared.SourcePatchRequest `request:"mediaType=application/json"`
	SourceID           string                     `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PatchSourceRequest) GetSourcePatchRequest() *shared.SourcePatchRequest {
	if o == nil {
		return nil
	}
	return o.SourcePatchRequest
}

func (o *PatchSourceRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PatchSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Update a Source
	SourceResponse *shared.SourceResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchSourceResponse) GetSourceResponse() *shared.SourceResponse {
	if o == nil {
		return nil
	}
	return o.SourceResponse
}

func (o *PatchSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
