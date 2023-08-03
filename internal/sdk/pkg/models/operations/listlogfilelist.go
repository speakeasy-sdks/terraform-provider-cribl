// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListLogFileListRequest struct {
	// query array[string] optional List of allowed-filename wildcard patterns
	Allow *string `queryParam:"style=form,explode=true,name=allow"`
	// Search depth for "manual" mode
	Depth *int64 `queryParam:"style=form,explode=true,name=depth"`
	// Discovery Mode (default is "auto")
	Mode *string `queryParam:"style=form,explode=true,name=mode"`
	// Search directory for "manual" mode
	Path *string `queryParam:"style=form,explode=true,name=path"`
}

type ListLogFileListResponse struct {
	ContentType string
	// a list of EdgeFile objects
	EdgeFiles *shared.EdgeFiles
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
