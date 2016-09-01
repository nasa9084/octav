package service

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"sync"
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/internal/errors"
	"github.com/builderscon/octav/octav/model"
	"github.com/lestrrat/go-pdebug"
)

var _ = time.Time{}

var sessionTypeSvc *SessionTypeSvc
var sessionTypeOnce sync.Once

func SessionType() *SessionTypeSvc {
	sessionTypeOnce.Do(sessionTypeSvc.Init)
	return sessionTypeSvc
}

func (v *SessionTypeSvc) LookupFromPayload(tx *db.Tx, m *model.SessionType, payload model.LookupSessionTypeRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.SessionType.LookupFromPayload").BindError(&err)
		defer g.End()
	}
	if err = v.Lookup(tx, m, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.SessionType from database")
	}
	if err := v.Decorate(tx, m, payload.TrustedCall, payload.Lang.String); err != nil {
		return errors.Wrap(err, "failed to load associated data for model.SessionType from database")
	}
	return nil
}
func (v *SessionTypeSvc) Lookup(tx *db.Tx, m *model.SessionType, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.SessionType.Lookup").BindError(&err)
		defer g.End()
	}

	r := model.SessionType{}
	if err = r.Load(tx, id); err != nil {
		return errors.Wrap(err, "failed to load model.SessionType from database")
	}
	*m = r
	return nil
}

// Create takes in the transaction, the incoming payload, and a reference to
// a database row. The database row is initialized/populated so that the
// caller can use it afterwards.
func (v *SessionTypeSvc) Create(tx *db.Tx, vdb *db.SessionType, payload model.CreateSessionTypeRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.SessionType.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Create(tx); err != nil {
		return err
	}

	if err := payload.L10N.CreateLocalizedStrings(tx, "SessionType", vdb.EID); err != nil {
		return err
	}
	return nil
}

func (v *SessionTypeSvc) Update(tx *db.Tx, vdb *db.SessionType, payload model.UpdateSessionTypeRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.SessionType.Update (%s)", vdb.EID).BindError(&err)
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
			ParentType: "SessionType",
			ParentID:   vdb.EID,
			Language:   l,
			Name:       k,
			Localized:  x,
		}
		return ls.Upsert(tx)
	})
}

func (v *SessionTypeSvc) ReplaceL10NStrings(tx *db.Tx, m *model.SessionType, lang string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("service.SessionType.ReplaceL10NStrings lang = %s", lang)
		defer g.End()
	}
	switch lang {
	case "", "en":
		if len(m.Name) > 0 && len(m.Abstract) > 0 {
			return nil
		}
		for _, extralang := range []string{`ja`} {
			rows, err := tx.Query(`SELECT localized FROM localized_strings WHERE parent_type = ? AND parent_id = ? AND language = ?`, "SessionType", m.ID, extralang)
			if err != nil {
				if errors.IsSQLNoRows(err) {
					break
				}
				return errors.Wrap(err, `failed to excute query`)
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
				case "name":
					if len(m.Name) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'name' (fallback en -> %s", l.Language)
						}
						m.Name = l.Localized
					}
				case "abstract":
					if len(m.Abstract) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'abstract' (fallback en -> %s", l.Language)
						}
						m.Abstract = l.Localized
					}
				}
			}
		}
		return nil
	case "all":
		rows, err := tx.Query(`SELECT oid, parent_id, parent_type, name, language, localized FROM localized_strings WHERE parent_type = ? AND parent_id = ?`, "SessionType", m.ID)
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
		rows, err := tx.Query(`SELECT oid, parent_id, parent_type, name, language, localized FROM localized_strings WHERE parent_type = ? AND parent_id = ? AND language = ?`, "SessionType", m.ID, lang)
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
			case "name":
				if pdebug.Enabled {
					pdebug.Printf("Replacing for key 'name'")
				}
				m.Name = l.Localized
			case "abstract":
				if pdebug.Enabled {
					pdebug.Printf("Replacing for key 'abstract'")
				}
				m.Abstract = l.Localized
			}
		}
	}
	return nil
}

func (v *SessionTypeSvc) Delete(tx *db.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("SessionType.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.SessionType{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	if err := db.DeleteLocalizedStringsForParent(tx, id, "SessionType"); err != nil {
		return err
	}
	return nil
}

func (v *SessionTypeSvc) LoadList(tx *db.Tx, vdbl *db.SessionTypeList, since string, limit int) error {
	return vdbl.LoadSinceEID(tx, since, limit)
}
