// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type BulletinMessageMetadata struct {
}

type BulletinMessageSeverity string

const (
	BulletinMessageSeverityInfo  BulletinMessageSeverity = "info"
	BulletinMessageSeverityWarn  BulletinMessageSeverity = "warn"
	BulletinMessageSeverityError BulletinMessageSeverity = "error"
	BulletinMessageSeverityFatal BulletinMessageSeverity = "fatal"
)

func (e BulletinMessageSeverity) ToPointer() *BulletinMessageSeverity {
	return &e
}

func (e *BulletinMessageSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "info":
		fallthrough
	case "warn":
		fallthrough
	case "error":
		fallthrough
	case "fatal":
		*e = BulletinMessageSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BulletinMessageSeverity: %v", v)
	}
}

// BulletinMessage - New BulletinMessage object
type BulletinMessage struct {
	Group    *string                   `json:"group,omitempty"`
	ID       string                    `json:"id"`
	Metadata []BulletinMessageMetadata `json:"metadata,omitempty"`
	Severity *BulletinMessageSeverity  `json:"severity,omitempty"`
	Text     string                    `json:"text"`
	Time     *int64                    `json:"time,omitempty"`
	Title    *string                   `json:"title,omitempty"`
}
