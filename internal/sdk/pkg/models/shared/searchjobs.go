// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SearchJobs - a list of SearchJob objects
type SearchJobs struct {
	// number of items present in the items array
	Count *int64      `json:"count,omitempty"`
	Items []SearchJob `json:"items,omitempty"`
}