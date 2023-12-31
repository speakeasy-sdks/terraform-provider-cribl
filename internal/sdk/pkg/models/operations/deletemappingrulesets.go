// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteMappingRulesetsRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteMappingRulesetsResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of MappingRuleset objects
	MappingRuleset *shared.MappingRuleset
	StatusCode     int
	RawResponse    *http.Response
}
