// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetHealthInfoResponse struct {
	ContentType string
	// Healthy
	HealthStatus *shared.HealthStatus
	StatusCode   int
	RawResponse  *http.Response
}