// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputConfluentCloudAcknowledgments - Control the number of required acknowledgments.
type OutputConfluentCloudAcknowledgments int64

const (
	OutputConfluentCloudAcknowledgmentsOne    OutputConfluentCloudAcknowledgments = 1
	OutputConfluentCloudAcknowledgmentsZero   OutputConfluentCloudAcknowledgments = 0
	OutputConfluentCloudAcknowledgmentsMinus1 OutputConfluentCloudAcknowledgments = -1
)

func (e OutputConfluentCloudAcknowledgments) ToPointer() *OutputConfluentCloudAcknowledgments {
	return &e
}

func (e *OutputConfluentCloudAcknowledgments) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 0:
		fallthrough
	case -1:
		*e = OutputConfluentCloudAcknowledgments(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudAcknowledgments: %v", v)
	}
}

// OutputConfluentCloudCompression - Codec to use to compress the data before sending to Kafka
type OutputConfluentCloudCompression string

const (
	OutputConfluentCloudCompressionNone   OutputConfluentCloudCompression = "none"
	OutputConfluentCloudCompressionGzip   OutputConfluentCloudCompression = "gzip"
	OutputConfluentCloudCompressionSnappy OutputConfluentCloudCompression = "snappy"
	OutputConfluentCloudCompressionLz4    OutputConfluentCloudCompression = "lz4"
)

func (e OutputConfluentCloudCompression) ToPointer() *OutputConfluentCloudCompression {
	return &e
}

func (e *OutputConfluentCloudCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		fallthrough
	case "snappy":
		fallthrough
	case "lz4":
		*e = OutputConfluentCloudCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudCompression: %v", v)
	}
}

// OutputConfluentCloudRecordDataFormat - Format to use to serialize events before writing to Kafka.
type OutputConfluentCloudRecordDataFormat string

const (
	OutputConfluentCloudRecordDataFormatJSON OutputConfluentCloudRecordDataFormat = "json"
	OutputConfluentCloudRecordDataFormatRaw  OutputConfluentCloudRecordDataFormat = "raw"
)

func (e OutputConfluentCloudRecordDataFormat) ToPointer() *OutputConfluentCloudRecordDataFormat {
	return &e
}

func (e *OutputConfluentCloudRecordDataFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "raw":
		*e = OutputConfluentCloudRecordDataFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudRecordDataFormat: %v", v)
	}
}

// OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion - Maximum TLS version to use when connecting
type OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion string

const (
	OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersionTlSv1  OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion = "TLSv1"
	OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersionTlSv11 OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion = "TLSv1.1"
	OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersionTlSv12 OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion = "TLSv1.2"
	OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersionTlSv13 OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion = "TLSv1.3"
)

func (e OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion) ToPointer() *OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion {
	return &e
}

func (e *OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion: %v", v)
	}
}

// OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion - Minimum TLS version to use when connecting
type OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion string

const (
	OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersionTlSv1  OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion = "TLSv1"
	OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersionTlSv11 OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion = "TLSv1.1"
	OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersionTlSv12 OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion = "TLSv1.2"
	OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersionTlSv13 OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion = "TLSv1.3"
)

func (e OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion) ToPointer() *OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion {
	return &e
}

func (e *OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion: %v", v)
	}
}

type OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSide struct {
	// Path on client in which to find CA certificates to verify the server's cert. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on client in which to find certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string `json:"certificateName,omitempty"`
	Disabled        *bool   `json:"disabled,omitempty"`
	// Maximum TLS version to use when connecting
	MaxVersion *OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMaximumTLSVersion `json:"maxVersion,omitempty"`
	// Minimum TLS version to use when connecting
	MinVersion *OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSideMinimumTLSVersion `json:"minVersion,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on client in which to find the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to No.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// Server name for the SNI (Server Name Indication) TLS extension. It must be a host name, and not an IP address.
	Servername *string `json:"servername,omitempty"`
}

type OutputConfluentCloudKafkaSchemaRegistryAuthentication struct {
	// Used when __keySchemaIdOut is not present, to transform key values, leave blank if key transformation is not required by default.
	DefaultKeySchemaID *int64 `json:"defaultKeySchemaId,omitempty"`
	// Used when __valueSchemaIdOut is not present, to transform _raw, leave blank if value transformation is not required by default.
	DefaultValueSchemaID *int64 `json:"defaultValueSchemaId,omitempty"`
	// Enable Schema Registry
	Disabled bool `json:"disabled"`
	// URL for access to the Confluent Schema Registry, i.e.: http://localhost:8081
	SchemaRegistryURL *string                                                                     `json:"schemaRegistryURL,omitempty"`
	TLS               *OutputConfluentCloudKafkaSchemaRegistryAuthenticationTLSSettingsClientSide `json:"tls,omitempty"`
}

// OutputConfluentCloudBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputConfluentCloudBackpressureBehavior string

const (
	OutputConfluentCloudBackpressureBehaviorQueue OutputConfluentCloudBackpressureBehavior = "queue"
	OutputConfluentCloudBackpressureBehaviorDrop  OutputConfluentCloudBackpressureBehavior = "drop"
	OutputConfluentCloudBackpressureBehaviorBlock OutputConfluentCloudBackpressureBehavior = "block"
)

func (e OutputConfluentCloudBackpressureBehavior) ToPointer() *OutputConfluentCloudBackpressureBehavior {
	return &e
}

func (e *OutputConfluentCloudBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputConfluentCloudBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudBackpressureBehavior: %v", v)
	}
}

// OutputConfluentCloudCompression1 - Codec to use to compress the persisted data.
type OutputConfluentCloudCompression1 string

const (
	OutputConfluentCloudCompression1None OutputConfluentCloudCompression1 = "none"
	OutputConfluentCloudCompression1Gzip OutputConfluentCloudCompression1 = "gzip"
)

func (e OutputConfluentCloudCompression1) ToPointer() *OutputConfluentCloudCompression1 {
	return &e
}

func (e *OutputConfluentCloudCompression1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputConfluentCloudCompression1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudCompression1: %v", v)
	}
}

type OutputConfluentCloudPqControls struct {
}

// OutputConfluentCloudQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputConfluentCloudQueueFullBehavior string

const (
	OutputConfluentCloudQueueFullBehaviorBlock OutputConfluentCloudQueueFullBehavior = "block"
	OutputConfluentCloudQueueFullBehaviorDrop  OutputConfluentCloudQueueFullBehavior = "drop"
)

func (e OutputConfluentCloudQueueFullBehavior) ToPointer() *OutputConfluentCloudQueueFullBehavior {
	return &e
}

func (e *OutputConfluentCloudQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputConfluentCloudQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudQueueFullBehavior: %v", v)
	}
}

// OutputConfluentCloudAuthenticationSASLMechanism - SASL authentication mechanism to use.
type OutputConfluentCloudAuthenticationSASLMechanism string

const (
	OutputConfluentCloudAuthenticationSASLMechanismPlain       OutputConfluentCloudAuthenticationSASLMechanism = "plain"
	OutputConfluentCloudAuthenticationSASLMechanismScramSha256 OutputConfluentCloudAuthenticationSASLMechanism = "scram-sha-256"
	OutputConfluentCloudAuthenticationSASLMechanismScramSha512 OutputConfluentCloudAuthenticationSASLMechanism = "scram-sha-512"
	OutputConfluentCloudAuthenticationSASLMechanismKerberos    OutputConfluentCloudAuthenticationSASLMechanism = "kerberos"
)

func (e OutputConfluentCloudAuthenticationSASLMechanism) ToPointer() *OutputConfluentCloudAuthenticationSASLMechanism {
	return &e
}

func (e *OutputConfluentCloudAuthenticationSASLMechanism) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "plain":
		fallthrough
	case "scram-sha-256":
		fallthrough
	case "scram-sha-512":
		fallthrough
	case "kerberos":
		*e = OutputConfluentCloudAuthenticationSASLMechanism(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudAuthenticationSASLMechanism: %v", v)
	}
}

// OutputConfluentCloudAuthentication - Authentication parameters to use when connecting to brokers. Using TLS is highly recommended.
type OutputConfluentCloudAuthentication struct {
	// Enable Authentication
	Disabled bool `json:"disabled"`
	// SASL authentication mechanism to use.
	Mechanism *OutputConfluentCloudAuthenticationSASLMechanism `json:"mechanism,omitempty"`
}

// OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion - Maximum TLS version to use when connecting
type OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion string

const (
	OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersionTlSv1  OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion = "TLSv1"
	OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersionTlSv11 OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion = "TLSv1.1"
	OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersionTlSv12 OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion = "TLSv1.2"
	OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersionTlSv13 OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion = "TLSv1.3"
)

func (e OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion) ToPointer() *OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion {
	return &e
}

func (e *OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion: %v", v)
	}
}

// OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion - Minimum TLS version to use when connecting
type OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion string

const (
	OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersionTlSv1  OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion = "TLSv1"
	OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersionTlSv11 OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion = "TLSv1.1"
	OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersionTlSv12 OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion = "TLSv1.2"
	OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersionTlSv13 OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion = "TLSv1.3"
)

func (e OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion) ToPointer() *OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion {
	return &e
}

func (e *OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion: %v", v)
	}
}

type OutputConfluentCloudTLSSettingsClientSide struct {
	// Path on client in which to find CA certificates to verify the server's cert. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on client in which to find certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string `json:"certificateName,omitempty"`
	Disabled        *bool   `json:"disabled,omitempty"`
	// Maximum TLS version to use when connecting
	MaxVersion *OutputConfluentCloudTLSSettingsClientSideMaximumTLSVersion `json:"maxVersion,omitempty"`
	// Minimum TLS version to use when connecting
	MinVersion *OutputConfluentCloudTLSSettingsClientSideMinimumTLSVersion `json:"minVersion,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on client in which to find the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to No.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// Server name for the SNI (Server Name Indication) TLS extension. It must be a host name, and not an IP address.
	Servername *string `json:"servername,omitempty"`
}

type OutputConfluentCloudType string

const (
	OutputConfluentCloudTypeConfluentCloud OutputConfluentCloudType = "confluent_cloud"
)

func (e OutputConfluentCloudType) ToPointer() *OutputConfluentCloudType {
	return &e
}

func (e *OutputConfluentCloudType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "confluent_cloud":
		*e = OutputConfluentCloudType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudType: %v", v)
	}
}

type OutputConfluentCloud struct {
	// Control the number of required acknowledgments.
	Ack *OutputConfluentCloudAcknowledgments `json:"ack,omitempty"`
	// Maximum time to wait for Kafka to respond to an authentication request
	AuthenticationTimeout *int64 `json:"authenticationTimeout,omitempty"`
	// List of Confluent Cloud brokers to connect to, e.g., yourAccount.confluent.cloud:9092.
	Brokers []string `json:"brokers"`
	// Codec to use to compress the data before sending to Kafka
	Compression *OutputConfluentCloudCompression `json:"compression,omitempty"`
	// Maximum time to wait for a connection to complete successfully
	ConnectionTimeout *int64 `json:"connectionTimeout,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// The maximum number of events you want the Destination to allow in a batch before forcing a flush
	FlushEventCount *int64 `json:"flushEventCount,omitempty"`
	// The maximum amount of time you want the Destination to wait before forcing a flush. Shorter intervals tend to result in smaller batches being sent.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Format to use to serialize events before writing to Kafka.
	Format *OutputConfluentCloudRecordDataFormat `json:"format,omitempty"`
	// Unique ID for this output
	ID                  *string                                                `json:"id,omitempty"`
	KafkaSchemaRegistry *OutputConfluentCloudKafkaSchemaRegistryAuthentication `json:"kafkaSchemaRegistry,omitempty"`
	// Maximum size of each record batch before compression. The value must not exceed the Kafka brokers' message.max.bytes setting.
	MaxRecordSizeKB *int64 `json:"maxRecordSizeKB,omitempty"`
	// If messages are failing, you can set the maximum number of retries as high as 100 to prevent loss of data.
	MaxRetries *int64 `json:"maxRetries,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputConfluentCloudBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputConfluentCloudCompression1 `json:"pqCompress,omitempty"`
	PqControls *OutputConfluentCloudPqControls   `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputConfluentCloudQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Specifies a time window during which @{product} can reauthenticate if needed. Creates the window measuring backwards from the moment when credentials are set to expire.
	ReauthenticationThreshold *int64 `json:"reauthenticationThreshold,omitempty"`
	// Maximum time to wait for Kafka to respond to a request
	RequestTimeout *int64 `json:"requestTimeout,omitempty"`
	// Authentication parameters to use when connecting to brokers. Using TLS is highly recommended.
	Sasl *OutputConfluentCloudAuthentication `json:"sasl,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string                                   `json:"systemFields,omitempty"`
	TLS          *OutputConfluentCloudTLSSettingsClientSide `json:"tls,omitempty"`
	// The topic to publish events to. Can be overridden using the __topicOut field.
	Topic string                    `json:"topic"`
	Type  *OutputConfluentCloudType `json:"type,omitempty"`
}
