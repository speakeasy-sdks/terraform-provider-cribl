// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Expression struct {
	MaxCache                     int64             `json:"MAX_CACHE"`
	Cache                        Map               `json:"cache"`
	DeclaredVariables            []string          `json:"declaredVariables"`
	IsSafe                       bool              `json:"isSafe"`
	ModifiedExpression           string            `json:"modifiedExpression"`
	Opt                          ExpressionOptions `json:"opt"`
	OriginalExpression           string            `json:"originalExpression"`
	PartialExpression            string            `json:"partialExpression"`
	ReferencedCriblExport        bool              `json:"referencedCriblExport"`
	ReplaceIdentifiersExpression string            `json:"replaceIdentifiersExpression"`
	ReplaceLiteralExpression     string            `json:"replaceLiteralExpression"`
}
