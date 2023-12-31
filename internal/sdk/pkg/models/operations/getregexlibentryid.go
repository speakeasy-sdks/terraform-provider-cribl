// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetRegexLibEntryIDRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetRegexLibEntryIDResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of RegexLibEntry objects
	RegexLibEntries *shared.RegexLibEntries
	StatusCode      int
	RawResponse     *http.Response
}
