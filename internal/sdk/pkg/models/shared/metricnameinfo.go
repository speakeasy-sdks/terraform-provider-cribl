// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MetricNameInfoDims struct {
	Count  int64    `json:"count"`
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type MetricNameInfo struct {
	Dims []MetricNameInfoDims `json:"dims"`
	Name string               `json:"name"`
}
