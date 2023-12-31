// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputTcpjsonAuthenticationMethod - Enter a token directly, or provide a secret referencing a token
type OutputTcpjsonAuthenticationMethod string

const (
	OutputTcpjsonAuthenticationMethodSecret OutputTcpjsonAuthenticationMethod = "secret"
	OutputTcpjsonAuthenticationMethodManual OutputTcpjsonAuthenticationMethod = "manual"
)

func (e OutputTcpjsonAuthenticationMethod) ToPointer() *OutputTcpjsonAuthenticationMethod {
	return &e
}

func (e *OutputTcpjsonAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = OutputTcpjsonAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputTcpjsonAuthenticationMethod: %v", v)
	}
}

// OutputTcpjsonCompression - Codec to use to compress the data before sending
type OutputTcpjsonCompression string

const (
	OutputTcpjsonCompressionNone OutputTcpjsonCompression = "none"
	OutputTcpjsonCompressionGzip OutputTcpjsonCompression = "gzip"
)

func (e OutputTcpjsonCompression) ToPointer() *OutputTcpjsonCompression {
	return &e
}

func (e *OutputTcpjsonCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputTcpjsonCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputTcpjsonCompression: %v", v)
	}
}

// OutputTcpjsonHostsTLS - Whether to inherit TLS configs from group setting or disable TLS.
type OutputTcpjsonHostsTLS string

const (
	OutputTcpjsonHostsTLSInherit OutputTcpjsonHostsTLS = "inherit"
	OutputTcpjsonHostsTLSOff     OutputTcpjsonHostsTLS = "off"
)

func (e OutputTcpjsonHostsTLS) ToPointer() *OutputTcpjsonHostsTLS {
	return &e
}

func (e *OutputTcpjsonHostsTLS) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "inherit":
		fallthrough
	case "off":
		*e = OutputTcpjsonHostsTLS(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputTcpjsonHostsTLS: %v", v)
	}
}

type OutputTcpjsonHosts struct {
	// The hostname of the receiver.
	Host string `json:"host"`
	// The port to connect to on the provided host.
	Port int64 `json:"port"`
	// Servername to use if establishing a TLS connection. If not specified, defaults to connection host (iff not an IP); otherwise, to the global TLS settings.
	Servername *string `json:"servername,omitempty"`
	// Whether to inherit TLS configs from group setting or disable TLS.
	TLS *OutputTcpjsonHostsTLS `json:"tls,omitempty"`
	// The weight to use for load-balancing purposes.
	Weight *int64 `json:"weight,omitempty"`
}

// OutputTcpjsonBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputTcpjsonBackpressureBehavior string

const (
	OutputTcpjsonBackpressureBehaviorQueue OutputTcpjsonBackpressureBehavior = "queue"
	OutputTcpjsonBackpressureBehaviorDrop  OutputTcpjsonBackpressureBehavior = "drop"
	OutputTcpjsonBackpressureBehaviorBlock OutputTcpjsonBackpressureBehavior = "block"
)

func (e OutputTcpjsonBackpressureBehavior) ToPointer() *OutputTcpjsonBackpressureBehavior {
	return &e
}

func (e *OutputTcpjsonBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputTcpjsonBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputTcpjsonBackpressureBehavior: %v", v)
	}
}

type OutputTcpjsonOptionalFieldsInGeneralSection string

const (
	OutputTcpjsonOptionalFieldsInGeneralSectionLoadBalanced OutputTcpjsonOptionalFieldsInGeneralSection = "loadBalanced"
	OutputTcpjsonOptionalFieldsInGeneralSectionHosts        OutputTcpjsonOptionalFieldsInGeneralSection = "hosts"
)

func (e OutputTcpjsonOptionalFieldsInGeneralSection) ToPointer() *OutputTcpjsonOptionalFieldsInGeneralSection {
	return &e
}

func (e *OutputTcpjsonOptionalFieldsInGeneralSection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "loadBalanced":
		fallthrough
	case "hosts":
		*e = OutputTcpjsonOptionalFieldsInGeneralSection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputTcpjsonOptionalFieldsInGeneralSection: %v", v)
	}
}

type OutputTcpjsonPqControls struct {
}

// OutputTcpjsonQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputTcpjsonQueueFullBehavior string

const (
	OutputTcpjsonQueueFullBehaviorBlock OutputTcpjsonQueueFullBehavior = "block"
	OutputTcpjsonQueueFullBehaviorDrop  OutputTcpjsonQueueFullBehavior = "drop"
)

func (e OutputTcpjsonQueueFullBehavior) ToPointer() *OutputTcpjsonQueueFullBehavior {
	return &e
}

func (e *OutputTcpjsonQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputTcpjsonQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputTcpjsonQueueFullBehavior: %v", v)
	}
}

// OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion - Maximum TLS version to use when connecting
type OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion string

const (
	OutputTcpjsonTLSSettingsClientSideMaximumTLSVersionTlSv1  OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion = "TLSv1"
	OutputTcpjsonTLSSettingsClientSideMaximumTLSVersionTlSv11 OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion = "TLSv1.1"
	OutputTcpjsonTLSSettingsClientSideMaximumTLSVersionTlSv12 OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion = "TLSv1.2"
	OutputTcpjsonTLSSettingsClientSideMaximumTLSVersionTlSv13 OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion = "TLSv1.3"
)

func (e OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion) ToPointer() *OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion {
	return &e
}

func (e *OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion: %v", v)
	}
}

// OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion - Minimum TLS version to use when connecting
type OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion string

const (
	OutputTcpjsonTLSSettingsClientSideMinimumTLSVersionTlSv1  OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion = "TLSv1"
	OutputTcpjsonTLSSettingsClientSideMinimumTLSVersionTlSv11 OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion = "TLSv1.1"
	OutputTcpjsonTLSSettingsClientSideMinimumTLSVersionTlSv12 OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion = "TLSv1.2"
	OutputTcpjsonTLSSettingsClientSideMinimumTLSVersionTlSv13 OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion = "TLSv1.3"
)

func (e OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion) ToPointer() *OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion {
	return &e
}

func (e *OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion: %v", v)
	}
}

type OutputTcpjsonTLSSettingsClientSide struct {
	// Path on client in which to find CA certificates to verify the server's cert. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on client in which to find certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string `json:"certificateName,omitempty"`
	Disabled        *bool   `json:"disabled,omitempty"`
	// Maximum TLS version to use when connecting
	MaxVersion *OutputTcpjsonTLSSettingsClientSideMaximumTLSVersion `json:"maxVersion,omitempty"`
	// Minimum TLS version to use when connecting
	MinVersion *OutputTcpjsonTLSSettingsClientSideMinimumTLSVersion `json:"minVersion,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on client in which to find the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to No.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// Server name for the SNI (Server Name Indication) TLS extension. It must be a host name, and not an IP address.
	Servername *string `json:"servername,omitempty"`
}

type OutputTcpjsonType string

const (
	OutputTcpjsonTypeTcpjson OutputTcpjsonType = "tcpjson"
)

func (e OutputTcpjsonType) ToPointer() *OutputTcpjsonType {
	return &e
}

func (e *OutputTcpjsonType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tcpjson":
		*e = OutputTcpjsonType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputTcpjsonType: %v", v)
	}
}

type OutputTcpjson struct {
	// Optional authentication token to include as part of the connection header
	AuthToken *string `json:"authToken,omitempty"`
	// Enter a token directly, or provide a secret referencing a token
	AuthType *OutputTcpjsonAuthenticationMethod `json:"authType,omitempty"`
	// Codec to use to compress the data before sending
	Compression *OutputTcpjsonCompression `json:"compression,omitempty"`
	// Amount of time (milliseconds) to wait for the connection to establish before retrying
	ConnectionTimeout *int64 `json:"connectionTimeout,omitempty"`
	// Re-resolve any hostnames every this many seconds and pick up destinations from A records.
	DNSResolvePeriodSec *int64 `json:"dnsResolvePeriodSec,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Exclude all IPs of the current host from the list of any resolved hostnames.
	ExcludeSelf *bool `json:"excludeSelf,omitempty"`
	// The hostname of the receiver
	Host *string `json:"host,omitempty"`
	// Set of hosts to load-balance data to.
	Hosts []OutputTcpjsonHosts `json:"hosts,omitempty"`
	// Unique ID for this output
	ID string `json:"id"`
	// How far back in time to keep traffic stats for load balancing purposes.
	LoadBalanceStatsPeriodSec *int64 `json:"loadBalanceStatsPeriodSec,omitempty"`
	// Use load-balanced destinations
	LoadBalanced *bool `json:"loadBalanced,omitempty"`
	// Maximum number of concurrent connections (per worker process). A random set of IPs will be picked on every DNS resolution period. Use 0 for unlimited.
	MaxConcurrentSenders *int64 `json:"maxConcurrentSenders,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure                 *OutputTcpjsonBackpressureBehavior           `json:"onBackpressure,omitempty"`
	OptionalFieldsInGeneralSection *OutputTcpjsonOptionalFieldsInGeneralSection `json:"optionalFieldsInGeneralSection,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// The port to connect to on the provided host
	Port *int64 `json:"port,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputTcpjsonCompression `json:"pqCompress,omitempty"`
	PqControls *OutputTcpjsonPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputTcpjsonQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
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
	ThrottleRatePerSec *string                             `json:"throttleRatePerSec,omitempty"`
	TLS                *OutputTcpjsonTLSSettingsClientSide `json:"tls,omitempty"`
	// The number of minutes before the internally generated authentication token expires, valid values between 1 and 60
	TokenTTLMinutes *int64            `json:"tokenTTLMinutes,omitempty"`
	Type            OutputTcpjsonType `json:"type"`
	// Amount of time (milliseconds) to wait for a write to complete before assuming connection is dead
	WriteTimeout *int64 `json:"writeTimeout,omitempty"`
}
