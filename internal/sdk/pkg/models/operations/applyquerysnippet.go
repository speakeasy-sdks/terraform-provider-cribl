// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ApplyQuerySnippetResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of PreviewResponseBody objects
	PreviewResponses *shared.PreviewResponses
	StatusCode       int
	RawResponse      *http.Response
}
