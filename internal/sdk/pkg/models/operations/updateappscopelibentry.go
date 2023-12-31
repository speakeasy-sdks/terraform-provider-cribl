// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateAppscopeLibEntryRequest struct {
	// AppscopeLibEntry object to be updated
	AppscopeLibEntry *shared.AppscopeLibEntry `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateAppscopeLibEntryResponse struct {
	// a list of AppscopeLibEntry objects
	AppscopeLibEntry *shared.AppscopeLibEntry
	ContentType      string
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
