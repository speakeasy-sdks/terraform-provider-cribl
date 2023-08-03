// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MetricsAggOpts - MetricsAggOpts object
type MetricsAggOpts struct {
	Aggs         AggregationMgrOptions `json:"aggs"`
	AlwaysBounds *bool                 `json:"alwaysBounds,omitempty"`
	Metrics      *MetricsStore         `json:"metrics,omitempty"`
	Where        *string               `json:"where,omitempty"`
}
