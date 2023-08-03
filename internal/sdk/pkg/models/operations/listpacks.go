// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListPacksResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of PackInfo objects
	PackInfos   *shared.PackInfos
	StatusCode  int
	RawResponse *http.Response
}
