// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InputLokiAuthenticationType - Loki logs authentication type
type InputLokiAuthenticationType string

const (
	InputLokiAuthenticationTypeOauth             InputLokiAuthenticationType = "oauth"
	InputLokiAuthenticationTypeBasic             InputLokiAuthenticationType = "basic"
	InputLokiAuthenticationTypeCredentialsSecret InputLokiAuthenticationType = "credentialsSecret"
	InputLokiAuthenticationTypeToken             InputLokiAuthenticationType = "token"
	InputLokiAuthenticationTypeTextSecret        InputLokiAuthenticationType = "textSecret"
)

func (e InputLokiAuthenticationType) ToPointer() *InputLokiAuthenticationType {
	return &e
}

func (e *InputLokiAuthenticationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth":
		fallthrough
	case "basic":
		fallthrough
	case "credentialsSecret":
		fallthrough
	case "token":
		fallthrough
	case "textSecret":
		*e = InputLokiAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputLokiAuthenticationType: %v", v)
	}
}

type InputLokiConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

type InputLokiMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

type InputLokiOauthHeaders struct {
	// OAuth header name
	Name string `json:"name"`
	// OAuth header value
	Value string `json:"value"`
}

type InputLokiOauthParams struct {
	// OAuth parameter name
	Name string `json:"name"`
	// OAuth parameter value
	Value string `json:"value"`
}

// InputLokiPqCompression - Codec to use to compress the persisted data.
type InputLokiPqCompression string

const (
	InputLokiPqCompressionNone InputLokiPqCompression = "none"
	InputLokiPqCompressionGzip InputLokiPqCompression = "gzip"
)

func (e InputLokiPqCompression) ToPointer() *InputLokiPqCompression {
	return &e
}

func (e *InputLokiPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputLokiPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputLokiPqCompression: %v", v)
	}
}

// InputLokiPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputLokiPqMode string

const (
	InputLokiPqModeSmart  InputLokiPqMode = "smart"
	InputLokiPqModeAlways InputLokiPqMode = "always"
)

func (e InputLokiPqMode) ToPointer() *InputLokiPqMode {
	return &e
}

func (e *InputLokiPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputLokiPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputLokiPqMode: %v", v)
	}
}

type InputLokiPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputLokiPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputLokiPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

// InputLokiTLSSettingsServerSideMaximumTLSVersion - Maximum TLS version to accept from connections.
type InputLokiTLSSettingsServerSideMaximumTLSVersion string

const (
	InputLokiTLSSettingsServerSideMaximumTLSVersionTlSv1  InputLokiTLSSettingsServerSideMaximumTLSVersion = "TLSv1"
	InputLokiTLSSettingsServerSideMaximumTLSVersionTlSv11 InputLokiTLSSettingsServerSideMaximumTLSVersion = "TLSv1.1"
	InputLokiTLSSettingsServerSideMaximumTLSVersionTlSv12 InputLokiTLSSettingsServerSideMaximumTLSVersion = "TLSv1.2"
	InputLokiTLSSettingsServerSideMaximumTLSVersionTlSv13 InputLokiTLSSettingsServerSideMaximumTLSVersion = "TLSv1.3"
)

func (e InputLokiTLSSettingsServerSideMaximumTLSVersion) ToPointer() *InputLokiTLSSettingsServerSideMaximumTLSVersion {
	return &e
}

func (e *InputLokiTLSSettingsServerSideMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputLokiTLSSettingsServerSideMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputLokiTLSSettingsServerSideMaximumTLSVersion: %v", v)
	}
}

// InputLokiTLSSettingsServerSideMinimumTLSVersion - Minimum TLS version to accept from connections.
type InputLokiTLSSettingsServerSideMinimumTLSVersion string

const (
	InputLokiTLSSettingsServerSideMinimumTLSVersionTlSv1  InputLokiTLSSettingsServerSideMinimumTLSVersion = "TLSv1"
	InputLokiTLSSettingsServerSideMinimumTLSVersionTlSv11 InputLokiTLSSettingsServerSideMinimumTLSVersion = "TLSv1.1"
	InputLokiTLSSettingsServerSideMinimumTLSVersionTlSv12 InputLokiTLSSettingsServerSideMinimumTLSVersion = "TLSv1.2"
	InputLokiTLSSettingsServerSideMinimumTLSVersionTlSv13 InputLokiTLSSettingsServerSideMinimumTLSVersion = "TLSv1.3"
)

func (e InputLokiTLSSettingsServerSideMinimumTLSVersion) ToPointer() *InputLokiTLSSettingsServerSideMinimumTLSVersion {
	return &e
}

func (e *InputLokiTLSSettingsServerSideMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputLokiTLSSettingsServerSideMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputLokiTLSSettingsServerSideMinimumTLSVersion: %v", v)
	}
}

type InputLokiTLSSettingsServerSide struct {
	// Path on server containing CA certificates to use. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on server containing certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string     `json:"certificateName,omitempty"`
	CommonNameRegex interface{} `json:"commonNameRegex,omitempty"`
	Disabled        *bool       `json:"disabled,omitempty"`
	// Maximum TLS version to accept from connections.
	MaxVersion *InputLokiTLSSettingsServerSideMaximumTLSVersion `json:"maxVersion,omitempty"`
	// Minimum TLS version to accept from connections.
	MinVersion *InputLokiTLSSettingsServerSideMinimumTLSVersion `json:"minVersion,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on server containing the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath        *string     `json:"privKeyPath,omitempty"`
	RejectUnauthorized interface{} `json:"rejectUnauthorized,omitempty"`
	// Whether to require clients to present their certificates. Used to perform client authentication using SSL certs.
	RequestCert *bool `json:"requestCert,omitempty"`
}

type InputLokiType string

const (
	InputLokiTypeLoki InputLokiType = "loki"
)

func (e InputLokiType) ToPointer() *InputLokiType {
	return &e
}

func (e *InputLokiType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "loki":
		*e = InputLokiType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputLokiType: %v", v)
	}
}

type InputLoki struct {
	// How often request activity is logged at the `info` level. A value of 1 would log every request, 10 every 10th request, etc.
	ActivityLogSampleRate *int64 `json:"activityLogSampleRate,omitempty"`
	// JavaScript expression to compute the Authorization header value to pass in requests. The value `${token}` is used to reference the token obtained from authentication, e.g.: `Bearer ${token}`.
	AuthHeaderExpr *string `json:"authHeaderExpr,omitempty"`
	// Loki logs authentication type
	AuthType *InputLokiAuthenticationType `json:"authType,omitempty"`
	// Toggle this to Yes to add request headers to events, in the __headers field.
	CaptureHeaders *bool `json:"captureHeaders,omitempty"`
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputLokiConnections `json:"connections,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
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
	// After the last response is sent, @{product} will wait this long for additional data before closing the socket connection. Minimum 1 sec.; maximum 600 sec. (10 min.).
	KeepAliveTimeout *int64 `json:"keepAliveTimeout,omitempty"`
	// URL for OAuth
	LoginURL *string `json:"loginUrl,omitempty"`
	// Absolute path on which to listen for Loki logs requests. Defaults to /loki/api/v1/push, which will (in this example) expand as: 'http://<your‑upstream‑URL>:<your‑port>/loki/api/v1/push'.
	LokiAPI string `json:"lokiAPI"`
	// Maximum number of active requests per Worker Process. Use 0 for unlimited.
	MaxActiveReq *int64 `json:"maxActiveReq,omitempty"`
	// Fields to add to events from this input.
	Metadata []InputLokiMetadata `json:"metadata,omitempty"`
	// Additional headers to send in the OAuth login request. @{product} will automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthHeaders []InputLokiOauthHeaders `json:"oauthHeaders,omitempty"`
	// Additional parameters to send in the OAuth login request. @{product} will combine the secret with these parameters, and will send the URL-encoded result in a POST request to the endpoint specified in the 'Login URL'. We'll automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthParams []InputLokiOauthParams `json:"oauthParams,omitempty"`
	// Password for Basic authentication
	Password *string `json:"password,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string `json:"pipeline,omitempty"`
	// Port to listen on.
	Port int64        `json:"port"`
	Pq   *InputLokiPq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// How long to wait for an incoming request to complete before aborting it. Use 0 to disable.
	RequestTimeout *int64 `json:"requestTimeout,omitempty"`
	// Secret parameter value to pass in request body
	Secret *string `json:"secret,omitempty"`
	// Secret parameter name to pass in request body
	SecretParamName *string `json:"secretParamName,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// How long @{product} should wait before assuming that an inactive socket has timed out. To wait forever, set to 0.
	SocketTimeout *int64  `json:"socketTimeout,omitempty"`
	Spacer        *string `json:"spacer,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string                         `json:"textSecret,omitempty"`
	TLS        *InputLokiTLSSettingsServerSide `json:"tls,omitempty"`
	// Bearer token to include in the authorization header
	Token *string `json:"token,omitempty"`
	// Name of the auth token attribute in the OAuth response. Can be top-level (e.g., 'token'); or nested, using a period (e.g., 'data.token').
	TokenAttributeName *string `json:"tokenAttributeName,omitempty"`
	// How often the OAuth token should be refreshed.
	TokenTimeoutSecs *int64         `json:"tokenTimeoutSecs,omitempty"`
	Type             *InputLokiType `json:"type,omitempty"`
	// Username for Basic authentication
	Username *string `json:"username,omitempty"`
}
