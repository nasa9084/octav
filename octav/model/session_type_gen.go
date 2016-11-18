package model

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"encoding/json"
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/lestrrat/go-pdebug"
)

var _ = pdebug.Enabled
var _ = time.Time{}

type rawSessionType struct {
	ID                    string     `json:"id"`
	ConferenceID          string     `json:"conference_id"`
	Name                  string     `json:"name" l10n:"true"`
	Abstract              string     `json:"abstract" l10n:"true"`
	Duration              int        `json:"duration"`
	SubmissionStart       *time.Time `json:"submission_start,omitempty"`
	SubmissionEnd         *time.Time `json:"submission_end,omitempty"`
	IsDefault             bool       `json:"is_default"`
	IsAcceptingSubmission bool       `json:"is_accepting_submission"`
}

func (v SessionType) MarshalJSON() ([]byte, error) {
	var raw rawSessionType
	raw.ID = v.ID
	raw.ConferenceID = v.ConferenceID
	raw.Name = v.Name
	raw.Abstract = v.Abstract
	raw.Duration = v.Duration
	if !v.SubmissionStart.IsZero() {
		raw.SubmissionStart = &v.SubmissionStart
	}
	if !v.SubmissionEnd.IsZero() {
		raw.SubmissionEnd = &v.SubmissionEnd
	}
	raw.IsDefault = v.IsDefault
	raw.IsAcceptingSubmission = v.IsAcceptingSubmission
	buf, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	return MarshalJSONWithL10N(buf, v.LocalizedFields)
}

func (v *SessionType) Load(tx *db.Tx, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("model.SessionType.Load %s", id).BindError(&err)
		defer g.End()
	}
	vdb := db.SessionType{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(&vdb); err != nil {
		return err
	}
	return nil
}

func (v *SessionType) FromRow(vdb *db.SessionType) error {
	v.ID = vdb.EID
	v.ConferenceID = vdb.ConferenceID
	v.Name = vdb.Name
	v.Abstract = vdb.Abstract
	v.Duration = vdb.Duration
	if vdb.SubmissionStart.Valid {
		v.SubmissionStart = vdb.SubmissionStart.Time
	}
	if vdb.SubmissionEnd.Valid {
		v.SubmissionEnd = vdb.SubmissionEnd.Time
	}
	v.IsDefault = vdb.IsDefault
	return nil
}

func (v *SessionType) ToRow(vdb *db.SessionType) error {
	vdb.EID = v.ID
	vdb.ConferenceID = v.ConferenceID
	vdb.Name = v.Name
	vdb.Abstract = v.Abstract
	vdb.Duration = v.Duration
	vdb.SubmissionStart.Valid = true
	vdb.SubmissionStart.Time = v.SubmissionStart
	vdb.SubmissionEnd.Valid = true
	vdb.SubmissionEnd.Time = v.SubmissionEnd
	vdb.IsDefault = v.IsDefault
	return nil
}
