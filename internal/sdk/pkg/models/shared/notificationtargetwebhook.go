// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// NotificationTargetWebhookAuthenticationType - The authentication method to use for the HTTP request. Defaults to None.
type NotificationTargetWebhookAuthenticationType string

const (
	NotificationTargetWebhookAuthenticationTypeNone  NotificationTargetWebhookAuthenticationType = "none"
	NotificationTargetWebhookAuthenticationTypeBasic NotificationTargetWebhookAuthenticationType = "basic"
	NotificationTargetWebhookAuthenticationTypeToken NotificationTargetWebhookAuthenticationType = "token"
)

func (e NotificationTargetWebhookAuthenticationType) ToPointer() *NotificationTargetWebhookAuthenticationType {
	return &e
}

func (e *NotificationTargetWebhookAuthenticationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "basic":
		fallthrough
	case "token":
		*e = NotificationTargetWebhookAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NotificationTargetWebhookAuthenticationType: %v", v)
	}
}

type NotificationTargetWebhookExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

// NotificationTargetWebhookFailedRequestLoggingMode - Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
type NotificationTargetWebhookFailedRequestLoggingMode string

const (
	NotificationTargetWebhookFailedRequestLoggingModePayload           NotificationTargetWebhookFailedRequestLoggingMode = "payload"
	NotificationTargetWebhookFailedRequestLoggingModePayloadAndHeaders NotificationTargetWebhookFailedRequestLoggingMode = "payloadAndHeaders"
	NotificationTargetWebhookFailedRequestLoggingModeNone              NotificationTargetWebhookFailedRequestLoggingMode = "none"
)

func (e NotificationTargetWebhookFailedRequestLoggingMode) ToPointer() *NotificationTargetWebhookFailedRequestLoggingMode {
	return &e
}

func (e *NotificationTargetWebhookFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "payload":
		fallthrough
	case "payloadAndHeaders":
		fallthrough
	case "none":
		*e = NotificationTargetWebhookFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NotificationTargetWebhookFailedRequestLoggingMode: %v", v)
	}
}

// NotificationTargetWebhookFormat - Specifies how to format events before sending out. Defaults to NDJSON.
type NotificationTargetWebhookFormat string

const (
	NotificationTargetWebhookFormatCustom    NotificationTargetWebhookFormat = "custom"
	NotificationTargetWebhookFormatJSONArray NotificationTargetWebhookFormat = "json_array"
)

func (e NotificationTargetWebhookFormat) ToPointer() *NotificationTargetWebhookFormat {
	return &e
}

func (e *NotificationTargetWebhookFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "custom":
		fallthrough
	case "json_array":
		*e = NotificationTargetWebhookFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NotificationTargetWebhookFormat: %v", v)
	}
}

// NotificationTargetWebhookMethod - The method to use when sending events. Defaults to POST.
type NotificationTargetWebhookMethod string

const (
	NotificationTargetWebhookMethodPost  NotificationTargetWebhookMethod = "POST"
	NotificationTargetWebhookMethodPut   NotificationTargetWebhookMethod = "PUT"
	NotificationTargetWebhookMethodPatch NotificationTargetWebhookMethod = "PATCH"
)

func (e NotificationTargetWebhookMethod) ToPointer() *NotificationTargetWebhookMethod {
	return &e
}

func (e *NotificationTargetWebhookMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POST":
		fallthrough
	case "PUT":
		fallthrough
	case "PATCH":
		*e = NotificationTargetWebhookMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NotificationTargetWebhookMethod: %v", v)
	}
}

type NotificationTargetWebhookType string

const (
	NotificationTargetWebhookTypeWebhook NotificationTargetWebhookType = "webhook"
)

func (e NotificationTargetWebhookType) ToPointer() *NotificationTargetWebhookType {
	return &e
}

func (e *NotificationTargetWebhookType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "webhook":
		*e = NotificationTargetWebhookType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NotificationTargetWebhookType: %v", v)
	}
}

type NotificationTargetWebhook struct {
	// The authentication method to use for the HTTP request. Defaults to None.
	AuthType *NotificationTargetWebhookAuthenticationType `json:"authType,omitempty"`
	// Whether to compress the payload body before sending.
	Compress *bool `json:"compress,omitempty"`
	// Maximum number of ongoing requests before blocking.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// Headers to add to all events. You can also add headers dynamically on a per-event basis in the __headers field, as explained [here](https://docs.cribl.io/stream/destinations-webhook/#internal-fields).
	ExtraHTTPHeaders []NotificationTargetWebhookExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
	FailedRequestLoggingMode *NotificationTargetWebhookFailedRequestLoggingMode `json:"failedRequestLoggingMode,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max body size.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Specifies how to format events before sending out. Defaults to NDJSON.
	Format *NotificationTargetWebhookFormat `json:"format,omitempty"`
	// Unique ID for this output
	ID string `json:"id"`
	// Toggle to No if you want @{product} to close the connection as soon as the outgoing request is sent. Defaults to Yes.
	KeepAlive *bool `json:"keepAlive,omitempty"`
	// Max number of events to include in the request body. Default is 0 (unlimited).
	MaxPayloadEvents *int64 `json:"maxPayloadEvents,omitempty"`
	// Maximum size, in KB, of the request body.
	MaxPayloadSizeKB *int64 `json:"maxPayloadSizeKB,omitempty"`
	// The method to use when sending events. Defaults to POST.
	Method *NotificationTargetWebhookMethod `json:"method,omitempty"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to No.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// List of headers that are safe to log in plain text.
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string `json:"systemFields,omitempty"`
	// Amount of time, in seconds, to wait for a request to complete before aborting it.
	TimeoutSec *int64                        `json:"timeoutSec,omitempty"`
	Type       NotificationTargetWebhookType `json:"type"`
	// URL to send events to. Can be overwritten by an event's __url field.
	URL string `json:"url"`
	// Enable to use round-robin DNS lookup. When a DNS server returns multiple addresses, this will cause Stream to cycle through them in the order returned.
	UseRoundRobinDNS *bool `json:"useRoundRobinDns,omitempty"`
}