// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListSchemasResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of Schema objects
	SchemaLibEntries *shared.SchemaLibEntries
	StatusCode       int
	RawResponse      *http.Response
}
