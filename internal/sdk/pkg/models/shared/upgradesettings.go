// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpgradeSettings struct {
	AutomaticUpgradeCheckPeriod *string              `json:"automaticUpgradeCheckPeriod,omitempty"`
	DisableAutomaticUpgrade     bool                 `json:"disableAutomaticUpgrade"`
	PackageUrls                 []UpgradePackageUrls `json:"packageUrls,omitempty"`
	UpgradeSource               string               `json:"upgradeSource"`
}
