// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TableViewSettingsViewMode string

const (
	TableViewSettingsViewModeEvent TableViewSettingsViewMode = "event"
	TableViewSettingsViewModeTable TableViewSettingsViewMode = "table"
)

func (e TableViewSettingsViewMode) ToPointer() *TableViewSettingsViewMode {
	return &e
}

func (e *TableViewSettingsViewMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "event":
		fallthrough
	case "table":
		*e = TableViewSettingsViewMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TableViewSettingsViewMode: %v", v)
	}
}

type TableViewSettings struct {
	ColumnFilterSettings   *ColumnFilterSettings     `json:"columnFilterSettings,omitempty"`
	ColumnFormatSettings   *ColumnFormatSettings     `json:"columnFormatSettings,omitempty"`
	ColumnOrderSettings    *ColumnOrderSettings      `json:"columnOrderSettings,omitempty"`
	ColumnSortSettings     *ColumnSortSettings       `json:"columnSortSettings,omitempty"`
	RowNumberColumnWidth   *int64                    `json:"rowNumberColumnWidth,omitempty"`
	ShowColumnTotals       bool                      `json:"showColumnTotals"`
	ShowColumnTotalsPinned bool                      `json:"showColumnTotalsPinned"`
	ShowRowNumbers         bool                      `json:"showRowNumbers"`
	ShowRowTotals          bool                      `json:"showRowTotals"`
	ShowRowTotalsPinned    bool                      `json:"showRowTotalsPinned"`
	ViewMode               TableViewSettingsViewMode `json:"viewMode"`
}