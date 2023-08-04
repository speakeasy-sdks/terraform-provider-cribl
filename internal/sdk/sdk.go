// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"Cribl/internal/sdk/pkg/utils"
	"fmt"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://logstream.{organizationID}.cribl.cloud/",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          *shared.Security
	ServerURL         string
	ServerIndex       int
	ServerDefaults    []map[string]string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

// CriblTerraform - Cribl API Reference: This API Reference lists available REST endpoints, along with their supported operations for accessing, creating, updating, or deleting resources. See our complementary product documentation at [docs.cribl.io](http://docs.cribl.io).
type CriblTerraform struct {
	AppscopeConfigs *appscopeConfigs
	// Auth - Actions related to authentication. Do not use the /auth endpoints in Cribl.Cloud deployments. Instead, obtain a Bearer token as described here: https://docs.cribl.io/stream/api-tutorials/#criblcloud-free-tier
	Auth      *auth
	Authorize *authorize
	// Certificates - Actions related to certificates
	Certificates *certificates
	Changelog    *changelog
	Clui         *clui
	// Collectors - Actions related to collectors
	Collectors          *collectors
	Conditions          *conditions
	DatabaseConnections *databaseConnections
	Datasets            *datasets
	// Diag - Actions related to diagnostics
	Diag        *diag
	Distributed *distributed
	// Edge - Actions enabled in Edge mode
	Edge           *edge
	EdgeContainers *edgeContainers
	EdgeEvents     *edgeEvents
	EdgeFiles      *edgeFiles
	EdgeLs         *edgeLs
	EdgeProcesses  *edgeProcesses
	// EventBreakerRules - Actions related to event breaker rules
	EventBreakerRules *eventBreakerRules
	Events            *events
	// Executors - Actions related to executors
	Executors *executors
	// Expressions - Actions related to expressions
	Expressions   *expressions
	Features      *features
	FileSampler   *fileSampler
	FleetMappings *fleetMappings
	// Functions - Actions related to functions
	Functions       *functions
	Git             *git
	GlobalVariables *globalVariables
	Grokfiles       *grokfiles
	Groups          *groups
	// Health - Actions related to REST server health
	Health *health
	Jobs   *jobs
	// Keys - Actions related to encryption keys
	Keys *keys
	// Licenses - Actions related to licenses. The <code>/licenses</code> endpoints do not apply to Cribl.Cloud deployments.
	Licenses *licenses
	Logger   *logger
	// Logging - Actions related to logging
	Logging *logging
	// Lookups - Actions related to lookups
	Lookups  *lookups
	Mappings *mappings
	// Messages - Actions related to messages
	Messages *messages
	// Metrics - Actions related to metrics
	Metrics             *metrics
	NotificationTargets *notificationTargets
	// Outputs - Actions related to outputs
	Outputs        *outputs
	Packs          *packs
	Parquetschemas *parquetschemas
	// Parsers - Actions related to parsers
	Parsers *parsers
	// Pipelines - Actions related to pipelines
	Pipelines *pipelines
	Policies  *policies
	// Preview - Actions related to data preview
	Preview   *preview
	Processes *processes
	Profiler  *profiler
	// Regexes - Actions related to regular expressions
	Regexes *regexes
	Roles   *roles
	// Routes - Actions related to routes
	Routes *routes
	// Samples - Actions related to samples
	Samples      *samples
	SavedJobs    *savedJobs
	SavedQueries *savedQueries
	Schemas      *schemas
	// Scripts - Actions related to scripts
	Scripts *scripts
	// Search - Actions related to search
	Search   *search
	Secrets  *secrets
	Security *security
	// System - Actions related to system settings
	System        *system
	TrustPolicies *trustPolicies
	UIState       *uiState
	// Users - Actions related to users
	Users      *users
	Versioning *versioning
	// Workers - Actions related to workers
	Workers *workers

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CriblTerraform)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *CriblTerraform) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *CriblTerraform) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *CriblTerraform) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithOrganizationID allows setting the $name variable for url substitution
func WithOrganizationID(organizationID string) SDKOption {
	return func(sdk *CriblTerraform) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["organizationID"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["organizationID"] = fmt.Sprintf("%v", organizationID)
		}
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CriblTerraform) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CriblTerraform) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *CriblTerraform {
	sdk := &CriblTerraform{
		sdkConfiguration: sdkConfiguration{
			Language:          "terraform",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "1.1.0",
			GenVersion:        "2.81.1",
			ServerDefaults: []map[string]string{
				{
					"organizationID": "api",
				},
			},
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.AppscopeConfigs = newAppscopeConfigs(sdk.sdkConfiguration)

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.Authorize = newAuthorize(sdk.sdkConfiguration)

	sdk.Certificates = newCertificates(sdk.sdkConfiguration)

	sdk.Changelog = newChangelog(sdk.sdkConfiguration)

	sdk.Clui = newClui(sdk.sdkConfiguration)

	sdk.Collectors = newCollectors(sdk.sdkConfiguration)

	sdk.Conditions = newConditions(sdk.sdkConfiguration)

	sdk.DatabaseConnections = newDatabaseConnections(sdk.sdkConfiguration)

	sdk.Datasets = newDatasets(sdk.sdkConfiguration)

	sdk.Diag = newDiag(sdk.sdkConfiguration)

	sdk.Distributed = newDistributed(sdk.sdkConfiguration)

	sdk.Edge = newEdge(sdk.sdkConfiguration)

	sdk.EdgeContainers = newEdgeContainers(sdk.sdkConfiguration)

	sdk.EdgeEvents = newEdgeEvents(sdk.sdkConfiguration)

	sdk.EdgeFiles = newEdgeFiles(sdk.sdkConfiguration)

	sdk.EdgeLs = newEdgeLs(sdk.sdkConfiguration)

	sdk.EdgeProcesses = newEdgeProcesses(sdk.sdkConfiguration)

	sdk.EventBreakerRules = newEventBreakerRules(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	sdk.Executors = newExecutors(sdk.sdkConfiguration)

	sdk.Expressions = newExpressions(sdk.sdkConfiguration)

	sdk.Features = newFeatures(sdk.sdkConfiguration)

	sdk.FileSampler = newFileSampler(sdk.sdkConfiguration)

	sdk.FleetMappings = newFleetMappings(sdk.sdkConfiguration)

	sdk.Functions = newFunctions(sdk.sdkConfiguration)

	sdk.Git = newGit(sdk.sdkConfiguration)

	sdk.GlobalVariables = newGlobalVariables(sdk.sdkConfiguration)

	sdk.Grokfiles = newGrokfiles(sdk.sdkConfiguration)

	sdk.Groups = newGroups(sdk.sdkConfiguration)

	sdk.Health = newHealth(sdk.sdkConfiguration)

	sdk.Jobs = newJobs(sdk.sdkConfiguration)

	sdk.Keys = newKeys(sdk.sdkConfiguration)

	sdk.Licenses = newLicenses(sdk.sdkConfiguration)

	sdk.Logger = newLogger(sdk.sdkConfiguration)

	sdk.Logging = newLogging(sdk.sdkConfiguration)

	sdk.Lookups = newLookups(sdk.sdkConfiguration)

	sdk.Mappings = newMappings(sdk.sdkConfiguration)

	sdk.Messages = newMessages(sdk.sdkConfiguration)

	sdk.Metrics = newMetrics(sdk.sdkConfiguration)

	sdk.NotificationTargets = newNotificationTargets(sdk.sdkConfiguration)

	sdk.Outputs = newOutputs(sdk.sdkConfiguration)

	sdk.Packs = newPacks(sdk.sdkConfiguration)

	sdk.Parquetschemas = newParquetschemas(sdk.sdkConfiguration)

	sdk.Parsers = newParsers(sdk.sdkConfiguration)

	sdk.Pipelines = newPipelines(sdk.sdkConfiguration)

	sdk.Policies = newPolicies(sdk.sdkConfiguration)

	sdk.Preview = newPreview(sdk.sdkConfiguration)

	sdk.Processes = newProcesses(sdk.sdkConfiguration)

	sdk.Profiler = newProfiler(sdk.sdkConfiguration)

	sdk.Regexes = newRegexes(sdk.sdkConfiguration)

	sdk.Roles = newRoles(sdk.sdkConfiguration)

	sdk.Routes = newRoutes(sdk.sdkConfiguration)

	sdk.Samples = newSamples(sdk.sdkConfiguration)

	sdk.SavedJobs = newSavedJobs(sdk.sdkConfiguration)

	sdk.SavedQueries = newSavedQueries(sdk.sdkConfiguration)

	sdk.Schemas = newSchemas(sdk.sdkConfiguration)

	sdk.Scripts = newScripts(sdk.sdkConfiguration)

	sdk.Search = newSearch(sdk.sdkConfiguration)

	sdk.Secrets = newSecrets(sdk.sdkConfiguration)

	sdk.Security = newSecurity(sdk.sdkConfiguration)

	sdk.System = newSystem(sdk.sdkConfiguration)

	sdk.TrustPolicies = newTrustPolicies(sdk.sdkConfiguration)

	sdk.UIState = newUIState(sdk.sdkConfiguration)

	sdk.Users = newUsers(sdk.sdkConfiguration)

	sdk.Versioning = newVersioning(sdk.sdkConfiguration)

	sdk.Workers = newWorkers(sdk.sdkConfiguration)

	return sdk
}
