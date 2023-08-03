// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RedisCacheLimits struct {
	ClientTrackingMechanism *string `json:"clientTrackingMechanism,omitempty"`
	EnableServerAssist      *bool   `json:"enableServerAssist,omitempty"`
	ID                      string  `json:"id"`
	KeyTTLSecs              *int64  `json:"keyTTLSecs,omitempty"`
	MaxCacheSize            *int64  `json:"maxCacheSize,omitempty"`
	MaxNumKeys              *int64  `json:"maxNumKeys,omitempty"`
	ServicePeriodSecs       *int64  `json:"servicePeriodSecs,omitempty"`
}
