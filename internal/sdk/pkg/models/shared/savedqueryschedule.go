// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SavedQuerySchedule struct {
	CronSchedule string `json:"cronSchedule"`
	Enabled      bool   `json:"enabled"`
	KeepLastN    int64  `json:"keepLastN"`
	Tz           string `json:"tz"`
}
