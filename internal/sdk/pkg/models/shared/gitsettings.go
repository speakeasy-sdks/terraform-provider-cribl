// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GitSettings - a list of string objects
type GitSettings struct {
	AuthType                 *string     `json:"authType,omitempty"`
	AutoAction               *string     `json:"autoAction,omitempty"`
	AutoActionMessage        *string     `json:"autoActionMessage,omitempty"`
	AutoActionSchedule       *string     `json:"autoActionSchedule,omitempty"`
	Branch                   *string     `json:"branch,omitempty"`
	CommitDeploySingleAction *bool       `json:"commitDeploySingleAction,omitempty"`
	DefaultCommitMessage     *string     `json:"defaultCommitMessage,omitempty"`
	GitOps                   *GitOpsType `json:"gitOps,omitempty"`
	Password                 *string     `json:"password,omitempty"`
	Remote                   *string     `json:"remote,omitempty"`
	SSHKey                   *string     `json:"sshKey,omitempty"`
	StrictHostKeyChecking    *bool       `json:"strictHostKeyChecking,omitempty"`
	Timeout                  *int64      `json:"timeout,omitempty"`
	User                     *string     `json:"user,omitempty"`
}
