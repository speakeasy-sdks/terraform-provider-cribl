// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type SavedJobType string

const (
	SavedJobTypeSavedJobCollection      SavedJobType = "SavedJobCollection"
	SavedJobTypeSavedJobExecutor        SavedJobType = "SavedJobExecutor"
	SavedJobTypeSavedJobScheduledSearch SavedJobType = "SavedJobScheduledSearch"
)

type SavedJob struct {
	SavedJobCollection      *SavedJobCollection
	SavedJobExecutor        *SavedJobExecutor
	SavedJobScheduledSearch *SavedJobScheduledSearch

	Type SavedJobType
}

func CreateSavedJobSavedJobCollection(savedJobCollection SavedJobCollection) SavedJob {
	typ := SavedJobTypeSavedJobCollection

	return SavedJob{
		SavedJobCollection: &savedJobCollection,
		Type:               typ,
	}
}

func CreateSavedJobSavedJobExecutor(savedJobExecutor SavedJobExecutor) SavedJob {
	typ := SavedJobTypeSavedJobExecutor

	return SavedJob{
		SavedJobExecutor: &savedJobExecutor,
		Type:             typ,
	}
}

func CreateSavedJobSavedJobScheduledSearch(savedJobScheduledSearch SavedJobScheduledSearch) SavedJob {
	typ := SavedJobTypeSavedJobScheduledSearch

	return SavedJob{
		SavedJobScheduledSearch: &savedJobScheduledSearch,
		Type:                    typ,
	}
}

func (u *SavedJob) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	savedJobCollection := new(SavedJobCollection)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&savedJobCollection); err == nil {
		u.SavedJobCollection = savedJobCollection
		u.Type = SavedJobTypeSavedJobCollection
		return nil
	}

	savedJobExecutor := new(SavedJobExecutor)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&savedJobExecutor); err == nil {
		u.SavedJobExecutor = savedJobExecutor
		u.Type = SavedJobTypeSavedJobExecutor
		return nil
	}

	savedJobScheduledSearch := new(SavedJobScheduledSearch)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&savedJobScheduledSearch); err == nil {
		u.SavedJobScheduledSearch = savedJobScheduledSearch
		u.Type = SavedJobTypeSavedJobScheduledSearch
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SavedJob) MarshalJSON() ([]byte, error) {
	if u.SavedJobCollection != nil {
		return json.Marshal(u.SavedJobCollection)
	}

	if u.SavedJobExecutor != nil {
		return json.Marshal(u.SavedJobExecutor)
	}

	if u.SavedJobScheduledSearch != nil {
		return json.Marshal(u.SavedJobScheduledSearch)
	}

	return nil, nil
}