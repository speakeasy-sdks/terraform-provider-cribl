// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostRegexLibEntryResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of RegexLibEntry objects
	RegexLibEntries *shared.RegexLibEntries
	StatusCode      int
	RawResponse     *http.Response
}
