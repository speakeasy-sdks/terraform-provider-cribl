// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DistributedSummaryGroups struct {
	Count        int64 `json:"count"`
	Destinations int64 `json:"destinations"`
	Pipelines    int64 `json:"pipelines"`
	Routes       int64 `json:"routes"`
	Sources      int64 `json:"sources"`
}

type DistributedSummaryWorkers struct {
	Alive            int64 `json:"alive"`
	ConfVersions     int64 `json:"confVersions"`
	Count            int64 `json:"count"`
	Groups           int64 `json:"groups"`
	SoftwareVersions int64 `json:"softwareVersions"`
	Unhealthy        int64 `json:"unhealthy"`
}

type DistributedSummary struct {
	Groups  DistributedSummaryGroups  `json:"groups"`
	Workers DistributedSummaryWorkers `json:"workers"`
}