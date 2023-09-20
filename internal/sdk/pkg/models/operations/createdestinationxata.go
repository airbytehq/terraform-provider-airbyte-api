// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateDestinationXataResponse struct {
	ContentType string
	// Successful operation
	DestinationXataGetResponse *shared.DestinationXataGetResponse
	StatusCode                 int
	RawResponse                *http.Response
}