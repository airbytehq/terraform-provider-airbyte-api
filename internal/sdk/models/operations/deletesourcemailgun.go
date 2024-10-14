// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSourceMailgunRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *DeleteSourceMailgunRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type DeleteSourceMailgunResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteSourceMailgunResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSourceMailgunResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSourceMailgunResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
