// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InputTcpjsonAuthenticationMethod - Enter a token directly, or provide a secret referencing a token
type InputTcpjsonAuthenticationMethod string

const (
	InputTcpjsonAuthenticationMethodSecret InputTcpjsonAuthenticationMethod = "secret"
	InputTcpjsonAuthenticationMethodManual InputTcpjsonAuthenticationMethod = "manual"
)

func (e InputTcpjsonAuthenticationMethod) ToPointer() *InputTcpjsonAuthenticationMethod {
	return &e
}

func (e *InputTcpjsonAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = InputTcpjsonAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputTcpjsonAuthenticationMethod: %v", v)
	}
}

type InputTcpjsonConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

type InputTcpjsonMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

// InputTcpjsonPqCompression - Codec to use to compress the persisted data.
type InputTcpjsonPqCompression string

const (
	InputTcpjsonPqCompressionNone InputTcpjsonPqCompression = "none"
	InputTcpjsonPqCompressionGzip InputTcpjsonPqCompression = "gzip"
)

func (e InputTcpjsonPqCompression) ToPointer() *InputTcpjsonPqCompression {
	return &e
}

func (e *InputTcpjsonPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputTcpjsonPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputTcpjsonPqCompression: %v", v)
	}
}

// InputTcpjsonPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputTcpjsonPqMode string

const (
	InputTcpjsonPqModeSmart  InputTcpjsonPqMode = "smart"
	InputTcpjsonPqModeAlways InputTcpjsonPqMode = "always"
)

func (e InputTcpjsonPqMode) ToPointer() *InputTcpjsonPqMode {
	return &e
}

func (e *InputTcpjsonPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputTcpjsonPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputTcpjsonPqMode: %v", v)
	}
}

type InputTcpjsonPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputTcpjsonPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputTcpjsonPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

// InputTcpjsonTLSSettingsServerSideMaximumTLSVersion - Maximum TLS version to accept from connections.
type InputTcpjsonTLSSettingsServerSideMaximumTLSVersion string

const (
	InputTcpjsonTLSSettingsServerSideMaximumTLSVersionTlSv1  InputTcpjsonTLSSettingsServerSideMaximumTLSVersion = "TLSv1"
	InputTcpjsonTLSSettingsServerSideMaximumTLSVersionTlSv11 InputTcpjsonTLSSettingsServerSideMaximumTLSVersion = "TLSv1.1"
	InputTcpjsonTLSSettingsServerSideMaximumTLSVersionTlSv12 InputTcpjsonTLSSettingsServerSideMaximumTLSVersion = "TLSv1.2"
	InputTcpjsonTLSSettingsServerSideMaximumTLSVersionTlSv13 InputTcpjsonTLSSettingsServerSideMaximumTLSVersion = "TLSv1.3"
)

func (e InputTcpjsonTLSSettingsServerSideMaximumTLSVersion) ToPointer() *InputTcpjsonTLSSettingsServerSideMaximumTLSVersion {
	return &e
}

func (e *InputTcpjsonTLSSettingsServerSideMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputTcpjsonTLSSettingsServerSideMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputTcpjsonTLSSettingsServerSideMaximumTLSVersion: %v", v)
	}
}

// InputTcpjsonTLSSettingsServerSideMinimumTLSVersion - Minimum TLS version to accept from connections.
type InputTcpjsonTLSSettingsServerSideMinimumTLSVersion string

const (
	InputTcpjsonTLSSettingsServerSideMinimumTLSVersionTlSv1  InputTcpjsonTLSSettingsServerSideMinimumTLSVersion = "TLSv1"
	InputTcpjsonTLSSettingsServerSideMinimumTLSVersionTlSv11 InputTcpjsonTLSSettingsServerSideMinimumTLSVersion = "TLSv1.1"
	InputTcpjsonTLSSettingsServerSideMinimumTLSVersionTlSv12 InputTcpjsonTLSSettingsServerSideMinimumTLSVersion = "TLSv1.2"
	InputTcpjsonTLSSettingsServerSideMinimumTLSVersionTlSv13 InputTcpjsonTLSSettingsServerSideMinimumTLSVersion = "TLSv1.3"
)

func (e InputTcpjsonTLSSettingsServerSideMinimumTLSVersion) ToPointer() *InputTcpjsonTLSSettingsServerSideMinimumTLSVersion {
	return &e
}

func (e *InputTcpjsonTLSSettingsServerSideMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputTcpjsonTLSSettingsServerSideMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputTcpjsonTLSSettingsServerSideMinimumTLSVersion: %v", v)
	}
}

type InputTcpjsonTLSSettingsServerSide struct {
	// Path on server containing CA certificates to use. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on server containing certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string     `json:"certificateName,omitempty"`
	CommonNameRegex interface{} `json:"commonNameRegex,omitempty"`
	Disabled        *bool       `json:"disabled,omitempty"`
	// Maximum TLS version to accept from connections.
	MaxVersion *InputTcpjsonTLSSettingsServerSideMaximumTLSVersion `json:"maxVersion,omitempty"`
	// Minimum TLS version to accept from connections.
	MinVersion *InputTcpjsonTLSSettingsServerSideMinimumTLSVersion `json:"minVersion,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on server containing the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath        *string     `json:"privKeyPath,omitempty"`
	RejectUnauthorized interface{} `json:"rejectUnauthorized,omitempty"`
	// Whether to require clients to present their certificates. Used to perform client authentication using SSL certs.
	RequestCert *bool `json:"requestCert,omitempty"`
}

type InputTcpjsonType string

const (
	InputTcpjsonTypeTcpjson InputTcpjsonType = "tcpjson"
)

func (e InputTcpjsonType) ToPointer() *InputTcpjsonType {
	return &e
}

func (e *InputTcpjsonType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tcpjson":
		*e = InputTcpjsonType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputTcpjsonType: %v", v)
	}
}

type InputTcpjson struct {
	// Shared secret to be provided by any client (in authToken header field). If empty, unauthed access is permitted.
	AuthToken *string `json:"authToken,omitempty"`
	// Enter a token directly, or provide a secret referencing a token
	AuthType *InputTcpjsonAuthenticationMethod `json:"authType,omitempty"`
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputTcpjsonConnections `json:"connections,omitempty"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// Enable if the connection is proxied by a device that supports Proxy Protocol V1 or V2.
	EnableProxyHeader *bool `json:"enableProxyHeader,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Address to bind on. Defaults to 0.0.0.0 (all addresses).
	Host string `json:"host"`
	// Unique ID for this input
	ID *string `json:"id,omitempty"`
	// Regex matching IP addresses that are allowed to establish a connection.
	IPWhitelistRegex *string `json:"ipWhitelistRegex,omitempty"`
	// Maximum number of active connections allowed per Worker Process, use 0 for unlimited
	MaxActiveCxn *int64 `json:"maxActiveCxn,omitempty"`
	// Fields to add to events from this input.
	Metadata []InputTcpjsonMetadata `json:"metadata,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string `json:"pipeline,omitempty"`
	// Port to listen on.
	Port int64           `json:"port"`
	Pq   *InputTcpjsonPq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string                            `json:"textSecret,omitempty"`
	TLS        *InputTcpjsonTLSSettingsServerSide `json:"tls,omitempty"`
	Type       *InputTcpjsonType                  `json:"type,omitempty"`
}
