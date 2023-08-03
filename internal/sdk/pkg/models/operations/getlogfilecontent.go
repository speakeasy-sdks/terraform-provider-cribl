// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetLogFileContentRequest struct {
	// in the current log file to fetch the log events upto.
	EndOffset *int64 `queryParam:"style=form,explode=true,name=endOffset"`
	// Epoch timestamp of the earliest event (includes rolled files present on disk)
	Et *int64 `queryParam:"style=form,explode=true,name=et"`
	// Filter
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
	// Log ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Maximum number of log lines to retrieve starting from offset.
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Epoch timestamp of the latest event (includes rolled files present on disk)
	Lt *int64 `queryParam:"style=form,explode=true,name=lt"`
}

type GetLogFileContentResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of any objects
	LogFileContents *shared.LogFileContents
	StatusCode      int
	RawResponse     *http.Response
}
