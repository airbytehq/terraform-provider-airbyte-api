// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceOutbrainAmplifyResponse struct {
	ContentType string
	// Successful operation
	SourceOutbrainAmplifyGetResponse *shared.SourceOutbrainAmplifyGetResponse
	StatusCode                       int
	RawResponse                      *http.Response
}