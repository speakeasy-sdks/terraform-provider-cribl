// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateScriptRequest struct {
	// Script object to be updated
	ScriptLibEntry *shared.ScriptLibEntry `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateScriptResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of Script objects
	ScriptLibEntry *shared.ScriptLibEntry
	StatusCode     int
	RawResponse    *http.Response
}
