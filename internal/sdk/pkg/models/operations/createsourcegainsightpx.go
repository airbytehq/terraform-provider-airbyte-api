// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceGainsightPxResponse struct {
	ContentType string
	// Successful operation
	SourceGainsightPxGetResponse *shared.SourceGainsightPxGetResponse
	StatusCode                   int
	RawResponse                  *http.Response
}