// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateEventBreakerRequest struct {
	// Event Breaker Ruleset object to be updated
	EventBreakerRuleset *shared.EventBreakerRuleset `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateEventBreakerResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of Event Breaker Ruleset objects
	EventBreakerRulesets *shared.EventBreakerRulesets
	StatusCode           int
	RawResponse          *http.Response
}