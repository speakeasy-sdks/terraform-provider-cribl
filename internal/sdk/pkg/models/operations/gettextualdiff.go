// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetTextualDiffRequest struct {
	// Commit hash (default is HEAD)
	Commit *string `queryParam:"style=form,explode=true,name=commit"`
	// Group ID
	Group *string `queryParam:"style=form,explode=true,name=group"`
}

type GetTextualDiffResponse struct {
	ContentType string
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
	// a list of any objects
	TextualDiff *shared.TextualDiff
}
