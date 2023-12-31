// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ProviderType string

const (
	ProviderTypeMmHeap    ProviderType = "mm-heap"
	ProviderTypeQuicksort ProviderType = "quicksort"
	ProviderTypeMerge     ProviderType = "merge"
	ProviderTypeInvalid   ProviderType = "invalid"
)

func (e ProviderType) ToPointer() *ProviderType {
	return &e
}

func (e *ProviderType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mm-heap":
		fallthrough
	case "quicksort":
		fallthrough
	case "merge":
		fallthrough
	case "invalid":
		*e = ProviderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProviderType: %v", v)
	}
}
