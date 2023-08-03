// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputSplunkAuthenticationMethod - Enter a token directly, or provide a secret referencing a token
type OutputSplunkAuthenticationMethod string

const (
	OutputSplunkAuthenticationMethodSecret OutputSplunkAuthenticationMethod = "secret"
	OutputSplunkAuthenticationMethodManual OutputSplunkAuthenticationMethod = "manual"
)

func (e OutputSplunkAuthenticationMethod) ToPointer() *OutputSplunkAuthenticationMethod {
	return &e
}

func (e *OutputSplunkAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = OutputSplunkAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkAuthenticationMethod: %v", v)
	}
}

// OutputSplunkMaxS2SVersion - The highest S2S protocol version to advertise during handshake.
type OutputSplunkMaxS2SVersion string

const (
	OutputSplunkMaxS2SVersionV3 OutputSplunkMaxS2SVersion = "v3"
	OutputSplunkMaxS2SVersionV4 OutputSplunkMaxS2SVersion = "v4"
)

func (e OutputSplunkMaxS2SVersion) ToPointer() *OutputSplunkMaxS2SVersion {
	return &e
}

func (e *OutputSplunkMaxS2SVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v3":
		fallthrough
	case "v4":
		*e = OutputSplunkMaxS2SVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkMaxS2SVersion: %v", v)
	}
}

// OutputSplunkNestedFieldSerialization - Specifies how to serialize nested fields into index-time fields.
type OutputSplunkNestedFieldSerialization string

const (
	OutputSplunkNestedFieldSerializationJSON OutputSplunkNestedFieldSerialization = "json"
	OutputSplunkNestedFieldSerializationNone OutputSplunkNestedFieldSerialization = "none"
)

func (e OutputSplunkNestedFieldSerialization) ToPointer() *OutputSplunkNestedFieldSerialization {
	return &e
}

func (e *OutputSplunkNestedFieldSerialization) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "none":
		*e = OutputSplunkNestedFieldSerialization(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkNestedFieldSerialization: %v", v)
	}
}

// OutputSplunkBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputSplunkBackpressureBehavior string

const (
	OutputSplunkBackpressureBehaviorQueue OutputSplunkBackpressureBehavior = "queue"
	OutputSplunkBackpressureBehaviorDrop  OutputSplunkBackpressureBehavior = "drop"
	OutputSplunkBackpressureBehaviorBlock OutputSplunkBackpressureBehavior = "block"
)

func (e OutputSplunkBackpressureBehavior) ToPointer() *OutputSplunkBackpressureBehavior {
	return &e
}

func (e *OutputSplunkBackpressureBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "queue":
		fallthrough
	case "drop":
		fallthrough
	case "block":
		*e = OutputSplunkBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkBackpressureBehavior: %v", v)
	}
}

// OutputSplunkCompression - Codec to use to compress the persisted data.
type OutputSplunkCompression string

const (
	OutputSplunkCompressionNone OutputSplunkCompression = "none"
	OutputSplunkCompressionGzip OutputSplunkCompression = "gzip"
)

func (e OutputSplunkCompression) ToPointer() *OutputSplunkCompression {
	return &e
}

func (e *OutputSplunkCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputSplunkCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkCompression: %v", v)
	}
}

type OutputSplunkPqControls struct {
}

// OutputSplunkQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputSplunkQueueFullBehavior string

const (
	OutputSplunkQueueFullBehaviorBlock OutputSplunkQueueFullBehavior = "block"
	OutputSplunkQueueFullBehaviorDrop  OutputSplunkQueueFullBehavior = "drop"
)

func (e OutputSplunkQueueFullBehavior) ToPointer() *OutputSplunkQueueFullBehavior {
	return &e
}

func (e *OutputSplunkQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputSplunkQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkQueueFullBehavior: %v", v)
	}
}

// OutputSplunkTLSSettingsClientSideMaximumTLSVersion - Maximum TLS version to use when connecting
type OutputSplunkTLSSettingsClientSideMaximumTLSVersion string

const (
	OutputSplunkTLSSettingsClientSideMaximumTLSVersionTlSv1  OutputSplunkTLSSettingsClientSideMaximumTLSVersion = "TLSv1"
	OutputSplunkTLSSettingsClientSideMaximumTLSVersionTlSv11 OutputSplunkTLSSettingsClientSideMaximumTLSVersion = "TLSv1.1"
	OutputSplunkTLSSettingsClientSideMaximumTLSVersionTlSv12 OutputSplunkTLSSettingsClientSideMaximumTLSVersion = "TLSv1.2"
	OutputSplunkTLSSettingsClientSideMaximumTLSVersionTlSv13 OutputSplunkTLSSettingsClientSideMaximumTLSVersion = "TLSv1.3"
)

func (e OutputSplunkTLSSettingsClientSideMaximumTLSVersion) ToPointer() *OutputSplunkTLSSettingsClientSideMaximumTLSVersion {
	return &e
}

func (e *OutputSplunkTLSSettingsClientSideMaximumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = OutputSplunkTLSSettingsClientSideMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkTLSSettingsClientSideMaximumTLSVersion: %v", v)
	}
}

// OutputSplunkTLSSettingsClientSideMinimumTLSVersion - Minimum TLS version to use when connecting
type OutputSplunkTLSSettingsClientSideMinimumTLSVersion string

const (
	OutputSplunkTLSSettingsClientSideMinimumTLSVersionTlSv1  OutputSplunkTLSSettingsClientSideMinimumTLSVersion = "TLSv1"
	OutputSplunkTLSSettingsClientSideMinimumTLSVersionTlSv11 OutputSplunkTLSSettingsClientSideMinimumTLSVersion = "TLSv1.1"
	OutputSplunkTLSSettingsClientSideMinimumTLSVersionTlSv12 OutputSplunkTLSSettingsClientSideMinimumTLSVersion = "TLSv1.2"
	OutputSplunkTLSSettingsClientSideMinimumTLSVersionTlSv13 OutputSplunkTLSSettingsClientSideMinimumTLSVersion = "TLSv1.3"
)

func (e OutputSplunkTLSSettingsClientSideMinimumTLSVersion) ToPointer() *OutputSplunkTLSSettingsClientSideMinimumTLSVersion {
	return &e
}

func (e *OutputSplunkTLSSettingsClientSideMinimumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = OutputSplunkTLSSettingsClientSideMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkTLSSettingsClientSideMinimumTLSVersion: %v", v)
	}
}

type OutputSplunkTLSSettingsClientSide struct {
	// Path on client in which to find CA certificates to verify the server's cert. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on client in which to find certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string `json:"certificateName,omitempty"`
	Disabled        *bool   `json:"disabled,omitempty"`
	// Maximum TLS version to use when connecting
	MaxVersion *OutputSplunkTLSSettingsClientSideMaximumTLSVersion `json:"maxVersion,omitempty"`
	// Minimum TLS version to use when connecting
	MinVersion *OutputSplunkTLSSettingsClientSideMinimumTLSVersion `json:"minVersion,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on client in which to find the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to No.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// Server name for the SNI (Server Name Indication) TLS extension. It must be a host name, and not an IP address.
	Servername *string `json:"servername,omitempty"`
}

type OutputSplunkType string

const (
	OutputSplunkTypeSplunk OutputSplunkType = "splunk"
)

func (e OutputSplunkType) ToPointer() *OutputSplunkType {
	return &e
}

func (e *OutputSplunkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "splunk":
		*e = OutputSplunkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkType: %v", v)
	}
}

type OutputSplunk struct {
	// Shared secret token to use when establishing a connection to a Splunk indexer.
	AuthToken *string `json:"authToken,omitempty"`
	// Enter a token directly, or provide a secret referencing a token
	AuthType *OutputSplunkAuthenticationMethod `json:"authType,omitempty"`
	// Amount of time (milliseconds) to wait for the connection to establish before retrying
	ConnectionTimeout *int64 `json:"connectionTimeout,omitempty"`
	// Check if indexer is shutting down and stop sending data. This helps minimize data loss during shutdown.
	EnableACK *bool `json:"enableACK,omitempty"`
	// Output metrics in multiple-metric format in a single event. Supported in Splunk 8.0 and above.
	EnableMultiMetrics *bool `json:"enableMultiMetrics,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// The hostname of the receiver
	Host string `json:"host"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// Maximum number of times healthcheck can fail before we close connection. If set to 0 (disabled), and the connection to Splunk is forcibly closed, some data loss might occur.
	MaxFailedHealthChecks *int64 `json:"maxFailedHealthChecks,omitempty"`
	// The highest S2S protocol version to advertise during handshake.
	MaxS2Sversion *OutputSplunkMaxS2SVersion `json:"maxS2Sversion,omitempty"`
	// Specifies how to serialize nested fields into index-time fields.
	NestedFields *OutputSplunkNestedFieldSerialization `json:"nestedFields,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputSplunkBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// The port to connect to on the provided host
	Port int64 `json:"port"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputSplunkCompression `json:"pqCompress,omitempty"`
	PqControls *OutputSplunkPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputSplunkQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string `json:"systemFields,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// Rate (in bytes per second) to throttle while writing to an output. Also takes values with multiple-byte units, such as KB, MB, GB, etc. (E.g., 42 MB.) Default value of 0 specifies no throttling.
	ThrottleRatePerSec *string                            `json:"throttleRatePerSec,omitempty"`
	TLS                *OutputSplunkTLSSettingsClientSide `json:"tls,omitempty"`
	Type               *OutputSplunkType                  `json:"type,omitempty"`
	// Amount of time (milliseconds) to wait for a write to complete before assuming connection is dead
	WriteTimeout *int64 `json:"writeTimeout,omitempty"`
}