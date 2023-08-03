// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SearchConfig struct {
	Datasets                  []string      `json:"datasets"`
	HasSendOperator           bool          `json:"hasSendOperator"`
	OrderedFieldNames         []string      `json:"orderedFieldNames"`
	SearchTerms               []SearchTerm  `json:"searchTerms"`
	SortFields                []SortByField `json:"sortFields,omitempty"`
	SuppressEventMarking      bool          `json:"suppressEventMarking"`
	UseFormattedVisualization bool          `json:"useFormattedVisualization"`
}
