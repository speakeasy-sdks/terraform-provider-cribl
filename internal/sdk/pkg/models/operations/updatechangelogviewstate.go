// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateChangelogViewStateResponse struct {
	// a list of ChangelogState objects
	ChangelogStateses *shared.ChangelogStateses
	ContentType       string
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}