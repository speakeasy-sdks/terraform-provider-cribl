// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListSearchLogsRequest struct {
	// search job id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type ListSearchLogsResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of string objects
	SearchLogs  *shared.SearchLogs
	StatusCode  int
	RawResponse *http.Response
}
