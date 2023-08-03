// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// NotificationTargetPagerDutySeverity - Default value for message severity, will be overwritten by value of __severity if set. Defaults to info.
type NotificationTargetPagerDutySeverity string

const (
	NotificationTargetPagerDutySeverityError    NotificationTargetPagerDutySeverity = "error"
	NotificationTargetPagerDutySeverityWarning  NotificationTargetPagerDutySeverity = "warning"
	NotificationTargetPagerDutySeverityInfo     NotificationTargetPagerDutySeverity = "info"
	NotificationTargetPagerDutySeverityCritical NotificationTargetPagerDutySeverity = "critical"
)

func (e NotificationTargetPagerDutySeverity) ToPointer() *NotificationTargetPagerDutySeverity {
	return &e
}

func (e *NotificationTargetPagerDutySeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "error":
		fallthrough
	case "warning":
		fallthrough
	case "info":
		fallthrough
	case "critical":
		*e = NotificationTargetPagerDutySeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NotificationTargetPagerDutySeverity: %v", v)
	}
}

type NotificationTargetPagerDutyType string

const (
	NotificationTargetPagerDutyTypePagerDuty NotificationTargetPagerDutyType = "pager_duty"
)

func (e NotificationTargetPagerDutyType) ToPointer() *NotificationTargetPagerDutyType {
	return &e
}

func (e *NotificationTargetPagerDutyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pager_duty":
		*e = NotificationTargetPagerDutyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NotificationTargetPagerDutyType: %v", v)
	}
}

type NotificationTargetPagerDuty struct {
	// Optional, default class value
	Class *string `json:"class,omitempty"`
	// Optional, default component value
	Component *string `json:"component,omitempty"`
	// Optional, default group value
	Group *string `json:"group,omitempty"`
	// Unique ID for this output
	ID string `json:"id"`
	// This is the 32 character Integration Key for an integration on a service or on a global ruleset.
	RoutingKey string `json:"routingKey"`
	// Default value for message severity, will be overwritten by value of __severity if set. Defaults to info.
	Severity *NotificationTargetPagerDutySeverity `json:"severity,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string                        `json:"systemFields,omitempty"`
	Type         NotificationTargetPagerDutyType `json:"type"`
}
