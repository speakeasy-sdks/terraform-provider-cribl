// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type DispatchSearchRequest struct {
	// search job id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DispatchSearchResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of any objects
	SearchID    *shared.SearchID
	StatusCode  int
	RawResponse *http.Response
}
