// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSourceLeverHiringRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type DeleteSourceLeverHiringResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}