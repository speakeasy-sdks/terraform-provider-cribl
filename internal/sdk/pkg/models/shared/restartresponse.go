// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type RestartResponseStatus string

const (
	RestartResponseStatusRestarting RestartResponseStatus = "Restarting"
	RestartResponseStatusError      RestartResponseStatus = "Error"
)

func (e RestartResponseStatus) ToPointer() *RestartResponseStatus {
	return &e
}

func (e *RestartResponseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Restarting":
		fallthrough
	case "Error":
		*e = RestartResponseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RestartResponseStatus: %v", v)
	}
}

type RestartResponse struct {
	ID      string                `json:"id"`
	Message *string               `json:"message,omitempty"`
	Status  RestartResponseStatus `json:"status"`
}
