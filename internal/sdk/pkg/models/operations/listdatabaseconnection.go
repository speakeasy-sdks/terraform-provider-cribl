// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListDatabaseConnectionRequest struct {
	// type of database connection
	DatabaseType *string `queryParam:"style=form,explode=true,name=databaseType"`
}

type ListDatabaseConnectionResponse struct {
	ContentType string
	// a list of DatabaseConnectionConfig objects
	DatabaseConnectionConfigs *shared.DatabaseConnectionConfigs
	// Unauthorized
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
