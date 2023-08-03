// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetProcessRunningDetailRequest struct {
	// Unique ID
	Pid string `pathParam:"style=simple,explode=false,name=pid"`
}

type GetProcessRunningDetailResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of Process objects
	Processes   *shared.Processes
	StatusCode  int
	RawResponse *http.Response
}