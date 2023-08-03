// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetVersioningResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of GitInfo objects
	GitInfos    *shared.GitInfos
	StatusCode  int
	RawResponse *http.Response
}
