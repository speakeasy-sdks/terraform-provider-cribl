// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListWorkerEdgeNodesRequest struct {
	// Filter expression evaluated against nodes
	FilterExp *string `queryParam:"style=form,explode=true,name=filterExp"`
	// Maximum number of nodes to return
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Pagination offset
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
	// Sorting expression evaluated against nodes
	SortExp *string `queryParam:"style=form,explode=true,name=sortExp"`
}

type ListWorkerEdgeNodesResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of MasterWorkerEntry objects
	MasterWorkerEntries *shared.MasterWorkerEntries
	StatusCode          int
	RawResponse         *http.Response
}
