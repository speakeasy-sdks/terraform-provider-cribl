// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostDatabaseConnectionResponse struct {
	ContentType string
	// a list of DatabaseConnectionConfig objects
	DatabaseConnectionConfigs *shared.DatabaseConnectionConfigs
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}