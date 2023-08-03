// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListChangedFilesRequest struct {
	// Commit ID
	ID *string `queryParam:"style=form,explode=true,name=ID"`
	// Group ID
	Group *string `queryParam:"style=form,explode=true,name=group"`
}

type ListChangedFilesResponse struct {
	// a list of GitFilesResponse objects
	ChangedFiles *shared.ChangedFiles
	ContentType  string
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
