// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetNamesOpts - GetNamesOpts object
type GetNamesOpts struct {
	DimKeyFilter     interface{} `json:"dimKeyFilter,omitempty"`
	DimValueFilter   interface{} `json:"dimValueFilter,omitempty"`
	Earliest         *int64      `json:"earliest,omitempty"`
	FilterExpr       *Expression `json:"filterExpr,omitempty"`
	MaxValues        *int64      `json:"maxValues,omitempty"`
	MetricNameFilter interface{} `json:"metricNameFilter,omitempty"`
}