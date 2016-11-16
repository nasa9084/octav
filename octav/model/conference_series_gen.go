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

type rawConferenceSeries struct {
	ID    string `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title" l10n:"true"`
}

func (v ConferenceSeries) MarshalJSON() ([]byte, error) {
	var raw rawConferenceSeries
	raw.ID = v.ID
	raw.Slug = v.Slug
	raw.Title = v.Title
	buf, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	return MarshalJSONWithL10N(buf, v.LocalizedFields)
}

func (v *ConferenceSeries) Load(tx *db.Tx, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("model.ConferenceSeries.Load %s", id).BindError(&err)
		defer g.End()
	}
	vdb := db.ConferenceSeries{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(vdb); err != nil {
		return err
	}
	return nil
}

func (v *ConferenceSeries) FromRow(vdb db.ConferenceSeries) error {
	v.ID = vdb.EID
	v.Slug = vdb.Slug
	v.Title = vdb.Title
	return nil
}

func (v *ConferenceSeries) ToRow(vdb *db.ConferenceSeries) error {
	vdb.EID = v.ID
	vdb.Slug = v.Slug
	vdb.Title = v.Title
	return nil
}
