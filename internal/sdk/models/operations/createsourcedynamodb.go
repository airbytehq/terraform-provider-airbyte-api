// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type CreateSourceDynamodbResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	SourceResponse *shared.SourceResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateSourceDynamodbResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSourceDynamodbResponse) GetSourceResponse() *shared.SourceResponse {
	if o == nil {
		return nil
	}
	return o.SourceResponse
}

func (o *CreateSourceDynamodbResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSourceDynamodbResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
