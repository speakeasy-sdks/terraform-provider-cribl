// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListOutputStatusResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of OutputStatus objects
	OutputStatuses *shared.OutputStatuses
	StatusCode     int
	RawResponse    *http.Response
}