// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateScriptResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of Script objects
	ScriptLibEntry *shared.ScriptLibEntry
	StatusCode     int
	RawResponse    *http.Response
}