// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DatabaseConnectionConfig - New DatabaseConnectionConfig object
type DatabaseConnectionConfig struct {
	AuthType          string                  `json:"authType"`
	ConfigObj         *string                 `json:"configObj,omitempty"`
	ConnectionString  *string                 `json:"connectionString,omitempty"`
	ConnectionTimeout *int64                  `json:"connectionTimeout,omitempty"`
	DatabaseType      *DatabaseConnectionType `json:"databaseType,omitempty"`
	Description       string                  `json:"description"`
	ID                string                  `json:"id"`
	RequestTimeout    *int64                  `json:"requestTimeout,omitempty"`
	Tags              *string                 `json:"tags,omitempty"`
}
