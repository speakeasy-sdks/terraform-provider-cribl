// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GitCommit - a list of GitCommitSummary objects
type GitCommit struct {
	// number of items present in the items array
	Count *int64             `json:"count,omitempty"`
	Items []GitCommitSummary `json:"items,omitempty"`
}
