// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConfigGroupGit struct {
	Commit       *string  `json:"commit,omitempty"`
	LocalChanges *int64   `json:"localChanges,omitempty"`
	Log          []Commit `json:"log,omitempty"`
}

// ConfigGroup - New ConfigGroup object
type ConfigGroup struct {
	ConfigVersion       string          `json:"configVersion"`
	Description         *string         `json:"description,omitempty"`
	EstimatedIngestRate *int64          `json:"estimatedIngestRate,omitempty"`
	Git                 *ConfigGroupGit `json:"git,omitempty"`
	ID                  string          `json:"id"`
	Inherits            *string         `json:"inherits,omitempty"`
	IsFleet             *bool           `json:"isFleet,omitempty"`
	IsSearch            *bool           `json:"isSearch,omitempty"`
	Name                *string         `json:"name,omitempty"`
	OnPrem              *bool           `json:"onPrem,omitempty"`
	Provisioned         *bool           `json:"provisioned,omitempty"`
	SourceGroupID       *string         `json:"sourceGroupId,omitempty"`
	Tags                *string         `json:"tags,omitempty"`
	WorkerCount         *int64          `json:"workerCount,omitempty"`
	WorkerRemoteAccess  *bool           `json:"workerRemoteAccess,omitempty"`
}
