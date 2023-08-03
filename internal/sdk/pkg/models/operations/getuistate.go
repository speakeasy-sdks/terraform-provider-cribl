// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetUIStateRequest struct {
	// UI state key
	Key string `pathParam:"style=simple,explode=false,name=key"`
}

type GetUIStateResponse struct {
	ContentType string
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
	// a list of any objects
	UIStates *shared.UIStates
}