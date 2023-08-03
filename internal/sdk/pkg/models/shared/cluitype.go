// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CluiType string

const (
	CluiTypeInput      CluiType = "input"
	CluiTypeOutput     CluiType = "output"
	CluiTypeRoute      CluiType = "route"
	CluiTypePipeline   CluiType = "pipeline"
	CluiTypeKnowledge  CluiType = "knowledge"
	CluiTypeCollector  CluiType = "collector"
	CluiTypePack       CluiType = "pack"
	CluiTypeMonitoring CluiType = "monitoring"
)

func (e CluiType) ToPointer() *CluiType {
	return &e
}

func (e *CluiType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "input":
		fallthrough
	case "output":
		fallthrough
	case "route":
		fallthrough
	case "pipeline":
		fallthrough
	case "knowledge":
		fallthrough
	case "collector":
		fallthrough
	case "pack":
		fallthrough
	case "monitoring":
		*e = CluiType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CluiType: %v", v)
	}
}
