// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IAMTrustPolicyStatementConditionStringEquals struct {
}

type IAMTrustPolicyStatementCondition struct {
	StringEquals *IAMTrustPolicyStatementConditionStringEquals `json:"StringEquals,omitempty"`
}

type IAMTrustPolicyStatementPrincipal struct {
	Aws string `json:"AWS"`
}

type IAMTrustPolicyStatement struct {
	Action    []string                          `json:"Action"`
	Condition *IAMTrustPolicyStatementCondition `json:"Condition,omitempty"`
	Effect    string                            `json:"Effect"`
	Principal IAMTrustPolicyStatementPrincipal  `json:"Principal"`
}

type IAMTrustPolicy struct {
	Statement []IAMTrustPolicyStatement `json:"Statement"`
	Version   string                    `json:"Version"`
}