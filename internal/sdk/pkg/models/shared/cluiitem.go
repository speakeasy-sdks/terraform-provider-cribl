// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CluiItem struct {
	Category CluiCategory `json:"category"`
	GroupID  *string      `json:"groupId,omitempty"`
	ID       *string      `json:"id,omitempty"`
	Name     *string      `json:"name,omitempty"`
	PackID   *string      `json:"packId,omitempty"`
	SubType  *string      `json:"subType,omitempty"`
	Type     CluiType     `json:"type"`
}
