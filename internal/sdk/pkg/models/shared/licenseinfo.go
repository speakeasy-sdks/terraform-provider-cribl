// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type LicenseInfoType string

const (
	LicenseInfoTypeProd  LicenseInfoType = "prod"
	LicenseInfoTypeTrial LicenseInfoType = "trial"
	LicenseInfoTypeFree  LicenseInfoType = "free"
)

func (e LicenseInfoType) ToPointer() *LicenseInfoType {
	return &e
}

func (e *LicenseInfoType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prod":
		fallthrough
	case "trial":
		fallthrough
	case "free":
		*e = LicenseInfoType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LicenseInfoType: %v", v)
	}
}

type LicenseInfo struct {
	EffectiveClass       LicenseType            `json:"effectiveClass"`
	Email                *string                `json:"email,omitempty"`
	Expiration           int64                  `json:"expiration"`
	GUID                 string                 `json:"guid"`
	IsRunningInSaaS      bool                   `json:"isRunningInSaaS"`
	IsSplunkApp          *bool                  `json:"isSplunkApp,omitempty"`
	LicenseEnforceReason string                 `json:"licenseEnforceReason"`
	Limits               map[string]interface{} `json:"limits"`
	PhoneHome            bool                   `json:"phoneHome"`
	PhoneHomeGrace       int64                  `json:"phoneHomeGrace"`
	Quota                int64                  `json:"quota"`
	Type                 LicenseInfoType        `json:"type"`
}
