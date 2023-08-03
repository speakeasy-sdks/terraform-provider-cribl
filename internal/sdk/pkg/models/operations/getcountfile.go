// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetCountFileRequest struct {
	// Group ID
	Group *string `queryParam:"style=form,explode=true,name=group"`
}

type GetCountFileResponse struct {
	ContentType string
	// a list of any objects
	CountFile *shared.CountFile
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
