// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListFieldSummariesRequest struct {
	// search job id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type ListFieldSummariesResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// FieldSummaries object
	FieldSummaries *shared.FieldSummaries
	StatusCode     int
	RawResponse    *http.Response
}
