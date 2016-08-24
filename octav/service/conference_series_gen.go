package service

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/model"
	"github.com/lestrrat/go-pdebug"
	"github.com/pkg/errors"
)

var _ = time.Time{}

func (v *ConferenceSeries) LookupFromPayload(tx *db.Tx, m *model.ConferenceSeries, payload model.LookupConferenceSeriesRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceSeries.LookupFromPayload").BindError(&err)
		defer g.End()
	}
	if err = v.Lookup(tx, m, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.ConferenceSeries from database")
	}
	if err := v.Decorate(tx, m, payload.TrustedCall, payload.Lang.String); err != nil {
		return errors.Wrap(err, "failed to load associated data for model.ConferenceSeries from database")
	}
	return nil
}
func (v *ConferenceSeries) Lookup(tx *db.Tx, m *model.ConferenceSeries, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceSeries.Lookup").BindError(&err)
		defer g.End()
	}

	r := model.ConferenceSeries{}
	if err = r.Load(tx, id); err != nil {
		return errors.Wrap(err, "failed to load model.ConferenceSeries from database")
	}
	*m = r
	return nil
}

// Create takes in the transaction, the incoming payload, and a reference to
// a database row. The database row is initialized/populated so that the
// caller can use it afterwards.
func (v *ConferenceSeries) Create(tx *db.Tx, vdb *db.ConferenceSeries, payload model.CreateConferenceSeriesRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceSeries.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Create(tx); err != nil {
		return err
	}

	if err := payload.L10N.CreateLocalizedStrings(tx, "ConferenceSeries", vdb.EID); err != nil {
		return err
	}
	return nil
}

func (v *ConferenceSeries) Update(tx *db.Tx, vdb *db.ConferenceSeries, payload model.UpdateConferenceSeriesRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceSeries.Update (%s)", vdb.EID).BindError(&err)
		defer g.End()
	}

	if vdb.EID == "" {
		return errors.New("vdb.EID is required (did you forget to call vdb.Load(tx) before hand?)")
	}

	if err := v.populateRowForUpdate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Update(tx); err != nil {
		return err
	}

	return payload.L10N.Foreach(func(l, k, x string) error {
		if pdebug.Enabled {
			pdebug.Printf("Updating l10n string for '%s' (%s)", k, l)
		}
		ls := db.LocalizedString{
			ParentType: "ConferenceSeries",
			ParentID:   vdb.EID,
			Language:   l,
			Name:       k,
			Localized:  x,
		}
		return ls.Upsert(tx)
	})
}

func (v *ConferenceSeries) ReplaceL10NStrings(tx *db.Tx, m *model.ConferenceSeries, lang string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceSeries.ReplaceL10NStrings lang = %s", lang)
		defer g.End()
	}
	switch lang {
	case "en":
		return nil
	case "all":
		rows, err := tx.Query(`SELECT oid, parent_id, parent_type, name, language, localized FROM localized_strings WHERE parent_type = ? AND parent_id = ?`, "ConferenceSeries", m.ID)
		if err != nil {
			return err
		}

		var l db.LocalizedString
		for rows.Next() {
			if err := l.Scan(rows); err != nil {
				return err
			}
			if len(l.Localized) == 0 {
				continue
			}
			if pdebug.Enabled {
				pdebug.Printf("Adding key '%s#%s'", l.Name, l.Language)
			}
			m.LocalizedFields.Set(l.Language, l.Name, l.Localized)
		}
	default:
		rows, err := tx.Query(`SELECT oid, parent_id, parent_type, name, language, localized FROM localized_strings WHERE parent_type = ? AND parent_id = ? AND language = ?`, "ConferenceSeries", m.ID, lang)
		if err != nil {
			return err
		}

		var l db.LocalizedString
		for rows.Next() {
			if err := l.Scan(rows); err != nil {
				return err
			}
			if len(l.Localized) == 0 {
				continue
			}

			switch l.Name {
			case "title":
				if pdebug.Enabled {
					pdebug.Printf("Replacing for key 'title'")
				}
				m.Title = l.Localized
			}
		}
	}
	return nil
}

func (v *ConferenceSeries) Delete(tx *db.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("ConferenceSeries.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.ConferenceSeries{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	if err := db.DeleteLocalizedStringsForParent(tx, id, "ConferenceSeries"); err != nil {
		return err
	}
	return nil
}

func (v *ConferenceSeries) LoadList(tx *db.Tx, vdbl *db.ConferenceSeriesList, since string, limit int) error {
	return vdbl.LoadSinceEID(tx, since, limit)
}
