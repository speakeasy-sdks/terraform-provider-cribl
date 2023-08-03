// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputWebhookAuthenticationType - The authentication method to use for the HTTP request. Defaults to None.
type OutputWebhookAuthenticationType string

const (
	OutputWebhookAuthenticationTypeOauth             OutputWebhookAuthenticationType = "oauth"
	OutputWebhookAuthenticationTypeBasic             OutputWebhookAuthenticationType = "basic"
	OutputWebhookAuthenticationTypeCredentialsSecret OutputWebhookAuthenticationType = "credentialsSecret"
	OutputWebhookAuthenticationTypeToken             OutputWebhookAuthenticationType = "token"
	OutputWebhookAuthenticationTypeTextSecret        OutputWebhookAuthenticationType = "textSecret"
	OutputWebhookAuthenticationTypeNone              OutputWebhookAuthenticationType = "none"
)

func (e OutputWebhookAuthenticationType) ToPointer() *OutputWebhookAuthenticationType {
	return &e
}

func (e *OutputWebhookAuthenticationType) UnmarshalJSON(data []byte) error {
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
		fallthrough
	case "none":
		*e = OutputWebhookAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWebhookAuthenticationType: %v", v)
	}
}

type OutputWebhookExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

// OutputWebhookFailedRequestLoggingMode - Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
type OutputWebhookFailedRequestLoggingMode string

const (
	OutputWebhookFailedRequestLoggingModePayload           OutputWebhookFailedRequestLoggingMode = "payload"
	OutputWebhookFailedRequestLoggingModePayloadAndHeaders OutputWebhookFailedRequestLoggingMode = "payloadAndHeaders"
	OutputWebhookFailedRequestLoggingModeNone              OutputWebhookFailedRequestLoggingMode = "none"
)

func (e OutputWebhookFailedRequestLoggingMode) ToPointer() *OutputWebhookFailedRequestLoggingMode {
	return &e
}

func (e *OutputWebhookFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputWebhookFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWebhookFailedRequestLoggingMode: %v", v)
	}
}

// OutputWebhookFormat - Specifies how to format events before sending out. Defaults to NDJSON.
type OutputWebhookFormat string

const (
	OutputWebhookFormatCustom    OutputWebhookFormat = "custom"
	OutputWebhookFormatJSONArray OutputWebhookFormat = "json_array"
)

func (e OutputWebhookFormat) ToPointer() *OutputWebhookFormat {
	return &e
}

func (e *OutputWebhookFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "custom":
		fallthrough
	case "json_array":
		*e = OutputWebhookFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWebhookFormat: %v", v)
	}
}

// OutputWebhookMethod - The method to use when sending events. Defaults to POST.
type OutputWebhookMethod string

const (
	OutputWebhookMethodPost  OutputWebhookMethod = "POST"
	OutputWebhookMethodPut   OutputWebhookMethod = "PUT"
	OutputWebhookMethodPatch OutputWebhookMethod = "PATCH"
)

func (e OutputWebhookMethod) ToPointer() *OutputWebhookMethod {
	return &e
}

func (e *OutputWebhookMethod) UnmarshalJSON(data []byte) error {
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
		*e = OutputWebhookMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWebhookMethod: %v", v)
	}
}

type OutputWebhookOauthHeaders struct {
	// OAuth header name
	Name string `json:"name"`
	// OAuth header value
	Value string `json:"value"`
}

type OutputWebhookOauthParams struct {
	// OAuth parameter name
	Name string `json:"name"`
	// OAuth parameter value
	Value string `json:"value"`
}

// OutputWebhookBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputWebhookBackpressureBehavior string

const (
	OutputWebhookBackpressureBehaviorQueue OutputWebhookBackpressureBehavior = "queue"
	OutputWebhookBackpressureBehaviorDrop  OutputWebhookBackpressureBehavior = "drop"
	OutputWebhookBackpressureBehaviorBlock OutputWebhookBackpressureBehavior = "block"
)

func (e OutputWebhookBackpressureBehavior) ToPointer() *OutputWebhookBackpressureBehavior {
	return &e
}

func (e *OutputWebhookBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputWebhookBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWebhookBackpressureBehavior: %v", v)
	}
}

// OutputWebhookCompression - Codec to use to compress the persisted data.
type OutputWebhookCompression string

const (
	OutputWebhookCompressionNone OutputWebhookCompression = "none"
	OutputWebhookCompressionGzip OutputWebhookCompression = "gzip"
)

func (e OutputWebhookCompression) ToPointer() *OutputWebhookCompression {
	return &e
}

func (e *OutputWebhookCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputWebhookCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWebhookCompression: %v", v)
	}
}

type OutputWebhookPqControls struct {
}

// OutputWebhookQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputWebhookQueueFullBehavior string

const (
	OutputWebhookQueueFullBehaviorBlock OutputWebhookQueueFullBehavior = "block"
	OutputWebhookQueueFullBehaviorDrop  OutputWebhookQueueFullBehavior = "drop"
)

func (e OutputWebhookQueueFullBehavior) ToPointer() *OutputWebhookQueueFullBehavior {
	return &e
}

func (e *OutputWebhookQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputWebhookQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWebhookQueueFullBehavior: %v", v)
	}
}

// OutputWebhookTLSSettingsClientSideMaximumTLSVersion - Maximum TLS version to use when connecting
type OutputWebhookTLSSettingsClientSideMaximumTLSVersion string

const (
	OutputWebhookTLSSettingsClientSideMaximumTLSVersionTlSv1  OutputWebhookTLSSettingsClientSideMaximumTLSVersion = "TLSv1"
	OutputWebhookTLSSettingsClientSideMaximumTLSVersionTlSv11 OutputWebhookTLSSettingsClientSideMaximumTLSVersion = "TLSv1.1"
	OutputWebhookTLSSettingsClientSideMaximumTLSVersionTlSv12 OutputWebhookTLSSettingsClientSideMaximumTLSVersion = "TLSv1.2"
	OutputWebhookTLSSettingsClientSideMaximumTLSVersionTlSv13 OutputWebhookTLSSettingsClientSideMaximumTLSVersion = "TLSv1.3"
)

func (e OutputWebhookTLSSettingsClientSideMaximumTLSVersion) ToPointer() *OutputWebhookTLSSettingsClientSideMaximumTLSVersion {
	return &e
}

func (e *OutputWebhookTLSSettingsClientSideMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputWebhookTLSSettingsClientSideMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWebhookTLSSettingsClientSideMaximumTLSVersion: %v", v)
	}
}

// OutputWebhookTLSSettingsClientSideMinimumTLSVersion - Minimum TLS version to use when connecting
type OutputWebhookTLSSettingsClientSideMinimumTLSVersion string

const (
	OutputWebhookTLSSettingsClientSideMinimumTLSVersionTlSv1  OutputWebhookTLSSettingsClientSideMinimumTLSVersion = "TLSv1"
	OutputWebhookTLSSettingsClientSideMinimumTLSVersionTlSv11 OutputWebhookTLSSettingsClientSideMinimumTLSVersion = "TLSv1.1"
	OutputWebhookTLSSettingsClientSideMinimumTLSVersionTlSv12 OutputWebhookTLSSettingsClientSideMinimumTLSVersion = "TLSv1.2"
	OutputWebhookTLSSettingsClientSideMinimumTLSVersionTlSv13 OutputWebhookTLSSettingsClientSideMinimumTLSVersion = "TLSv1.3"
)

func (e OutputWebhookTLSSettingsClientSideMinimumTLSVersion) ToPointer() *OutputWebhookTLSSettingsClientSideMinimumTLSVersion {
	return &e
}

func (e *OutputWebhookTLSSettingsClientSideMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputWebhookTLSSettingsClientSideMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWebhookTLSSettingsClientSideMinimumTLSVersion: %v", v)
	}
}

type OutputWebhookTLSSettingsClientSide struct {
	// Path on client in which to find CA certificates to verify the server's cert. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on client in which to find certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string `json:"certificateName,omitempty"`
	Disabled        *bool   `json:"disabled,omitempty"`
	// Maximum TLS version to use when connecting
	MaxVersion *OutputWebhookTLSSettingsClientSideMaximumTLSVersion `json:"maxVersion,omitempty"`
	// Minimum TLS version to use when connecting
	MinVersion *OutputWebhookTLSSettingsClientSideMinimumTLSVersion `json:"minVersion,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on client in which to find the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Server name for the SNI (Server Name Indication) TLS extension. It must be a host name, and not an IP address.
	Servername *string `json:"servername,omitempty"`
}

type OutputWebhookType string

const (
	OutputWebhookTypeWebhook OutputWebhookType = "webhook"
)

func (e OutputWebhookType) ToPointer() *OutputWebhookType {
	return &e
}

func (e *OutputWebhookType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "webhook":
		*e = OutputWebhookType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWebhookType: %v", v)
	}
}

type OutputWebhook struct {
	// JavaScript expression to compute the Authorization header value to pass in requests. The value `${token}` is used to reference the token obtained from authentication, e.g.: `Bearer ${token}`.
	AuthHeaderExpr *string `json:"authHeaderExpr,omitempty"`
	// The authentication method to use for the HTTP request. Defaults to None.
	AuthType *OutputWebhookAuthenticationType `json:"authType,omitempty"`
	// Whether to compress the payload body before sending.
	Compress *bool `json:"compress,omitempty"`
	// Maximum number of ongoing requests before blocking.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Content type to use for request. Defaults to application/x-ndjson. Any content types set in Advanced Settings > Extra HTTP headers will override this entry.
	CustomContentType *string `json:"customContentType,omitempty"`
	// Whether or not to drop events when the source expression evaluates to null.
	CustomDropWhenNull *bool `json:"customDropWhenNull,omitempty"`
	// Delimiter string to insert between individual events. Defaults to newline character.
	CustomEventDelimiter *string `json:"customEventDelimiter,omitempty"`
	// Expression specifying how to format the payload for each batch. To reference the events to send, use the `${events}` variable. Example expression: `{ "items" : [${events}] }` would send the batch inside a JSON object.
	CustomPayloadExpression *string `json:"customPayloadExpression,omitempty"`
	// Expression to evaluate on events, to generate output notifications. E.g.: `notification=${_raw}`. See docs for other fields available. If empty, we send the full notification event as stringified JSON
	CustomSourceExpression *string `json:"customSourceExpression,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Headers to add to all events. You can also add headers dynamically on a per-event basis in the __headers field, as explained [here](https://docs.cribl.io/stream/destinations-webhook/#internal-fields).
	ExtraHTTPHeaders []OutputWebhookExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
	FailedRequestLoggingMode *OutputWebhookFailedRequestLoggingMode `json:"failedRequestLoggingMode,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max body size.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Specifies how to format events before sending out. Defaults to NDJSON.
	Format *OutputWebhookFormat `json:"format,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// Toggle to No if you want @{product} to close the connection as soon as the outgoing request is sent. Defaults to Yes.
	KeepAlive *bool `json:"keepAlive,omitempty"`
	// URL for OAuth
	LoginURL *string `json:"loginUrl,omitempty"`
	// Max number of events to include in the request body. Default is 0 (unlimited).
	MaxPayloadEvents *int64 `json:"maxPayloadEvents,omitempty"`
	// Maximum size, in KB, of the request body.
	MaxPayloadSizeKB *int64 `json:"maxPayloadSizeKB,omitempty"`
	// The method to use when sending events. Defaults to POST.
	Method *OutputWebhookMethod `json:"method,omitempty"`
	// Additional headers to send in the OAuth login request. @{product} will automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthHeaders []OutputWebhookOauthHeaders `json:"oauthHeaders,omitempty"`
	// Additional parameters to send in the OAuth login request. @{product} will combine the secret with these parameters, and will send the URL-encoded result in a POST request to the endpoint specified in the 'Login URL'. We'll automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthParams []OutputWebhookOauthParams `json:"oauthParams,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputWebhookBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Password for Basic authentication
	Password *string `json:"password,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputWebhookCompression `json:"pqCompress,omitempty"`
	PqControls *OutputWebhookPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputWebhookQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to No.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// List of headers that are safe to log in plain text.
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Secret parameter value to pass in request body
	Secret *string `json:"secret,omitempty"`
	// Secret parameter name to pass in request body
	SecretParamName *string `json:"secretParamName,omitempty"`
	Spacer          *string `json:"spacer,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string `json:"systemFields,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// Amount of time, in seconds, to wait for a request to complete before aborting it.
	TimeoutSec *int64                              `json:"timeoutSec,omitempty"`
	TLS        *OutputWebhookTLSSettingsClientSide `json:"tls,omitempty"`
	// Bearer token to include in the authorization header
	Token *string `json:"token,omitempty"`
	// Name of the auth token attribute in the OAuth response. Can be top-level (e.g., 'token'); or nested, using a period (e.g., 'data.token').
	TokenAttributeName *string `json:"tokenAttributeName,omitempty"`
	// How often the OAuth token should be refreshed.
	TokenTimeoutSecs *int64            `json:"tokenTimeoutSecs,omitempty"`
	Type             OutputWebhookType `json:"type"`
	// URL to send events to. Can be overwritten by an event's __url field.
	URL string `json:"url"`
	// Enable to use round-robin DNS lookup. When a DNS server returns multiple addresses, this will cause Stream to cycle through them in the order returned.
	UseRoundRobinDNS *bool `json:"useRoundRobinDns,omitempty"`
	// Username for Basic authentication
	Username *string `json:"username,omitempty"`
}
