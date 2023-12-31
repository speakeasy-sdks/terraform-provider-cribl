// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListJobInfosRequest struct {
	// Filter by collector id, e.g. "collectorId=Prometheus-in"
	CollectorID *string `queryParam:"style=form,explode=true,name=collectorId"`
	// Filter by job id, e.g. "id=1608713335.3&id=1608713326.1"
	ID *string `queryParam:"style=form,explode=true,name=id"`
	// Maximum number of items to return
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Pagination offset
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
	// Filter by job run type, one of "adhoc", "scheduled", "system"
	RunType *string `queryParam:"style=form,explode=true,name=runType"`
	// Filter by current job state, e.g. "running"
	State *string `queryParam:"style=form,explode=true,name=state"`
}

type ListJobInfosResponse struct {
	ContentType string
	// Unauthorized
	Error *shared.Error
	// a list of JobInfo objects
	JobInfos    *shared.JobInfos
	StatusCode  int
	RawResponse *http.Response
}
