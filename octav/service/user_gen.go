package service

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"context"
	"sync"
	"time"

	"github.com/builderscon/octav/octav/cache"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/internal/errors"
	"github.com/builderscon/octav/octav/model"
	"github.com/lestrrat/go-pdebug"
)

var _ = time.Time{}
var _ = cache.WithExpires(time.Minute)
var _ = context.Background
var _ = errors.Wrap
var _ = model.User{}
var _ = db.User{}
var _ = pdebug.Enabled

var userSvc UserSvc
var userOnce sync.Once

func User() *UserSvc {
	userOnce.Do(userSvc.Init)
	return &userSvc
}

func (v *UserSvc) LookupFromPayload(tx *db.Tx, m *model.User, payload *model.LookupUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.LookupFromPayload").BindError(&err)
		defer g.End()
	}
	if err = v.Lookup(tx, m, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.User from database")
	}
	if err := v.Decorate(tx, m, payload.TrustedCall, payload.Lang.String); err != nil {
		return errors.Wrap(err, "failed to load associated data for model.User from database")
	}
	return nil
}

func (v *UserSvc) Lookup(tx *db.Tx, m *model.User, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.Lookup").BindError(&err)
		defer g.End()
	}

	var r model.User
	c := Cache()
	key := c.Key("User", id)
	var cacheMiss bool
	_, err = c.GetOrSet(key, &r, func() (interface{}, error) {
		if pdebug.Enabled {
			cacheMiss = true
		}
		if err := r.Load(tx, id); err != nil {
			return nil, errors.Wrap(err, "failed to load model.User from database")
		}
		return &r, nil
	}, cache.WithExpires(time.Hour))
	if pdebug.Enabled {
		cacheSt := `HIT`
		if cacheMiss {
			cacheSt = `MISS`
		}
		pdebug.Printf(`CACHE %s: %s`, cacheSt, key)
	}
	if err = v.PostLookupFromPayloadHook(tx, &r); err != nil {
		return errors.Wrap(err, "failed to execute PostLookupFromPayloadHook")
	}
	*m = r
	return nil
}

// Create takes in the transaction, the incoming payload, and a reference to
// a database row. The database row is initialized/populated so that the
// caller can use it afterwards.
func (v *UserSvc) Create(tx *db.Tx, vdb *db.User, payload *model.CreateUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return errors.Wrap(err, `failed to populate row`)
	}

	if err := vdb.Create(tx, payload.DatabaseOptions...); err != nil {
		return errors.Wrap(err, `failed to insert into database`)
	}

	if err := payload.LocalizedFields.CreateLocalizedStrings(tx, "User", vdb.EID); err != nil {
		return errors.Wrap(err, `failed to populate localized strings`)
	}
	return nil
}

func (v *UserSvc) Update(tx *db.Tx, vdb *db.User) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.Update (%s)", vdb.EID).BindError(&err)
		defer g.End()
	}

	if vdb.EID == `` {
		return errors.New("vdb.EID is required (did you forget to call vdb.Load(tx) before hand?)")
	}

	if err := vdb.Update(tx); err != nil {
		return errors.Wrap(err, `failed to update database`)
	}
	c := Cache()
	key := c.Key("User", vdb.EID)
	if pdebug.Enabled {
		pdebug.Printf(`CACHE DEL %s`, key)
	}
	cerr := c.Delete(key)
	if pdebug.Enabled {
		if cerr != nil {
			pdebug.Printf(`CACHE ERR: %s`, cerr)
		}
	}
	return nil
}

func (v *UserSvc) UpdateFromPayload(ctx context.Context, tx *db.Tx, payload *model.UpdateUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.UpdateFromPayload (%s)", payload.ID).BindError(&err)
		defer g.End()
	}
	var vdb db.User
	if err := vdb.LoadByEID(tx, payload.ID); err != nil {
		return errors.Wrap(err, `failed to load from database`)
	}

	if err := v.PreUpdateFromPayloadHook(ctx, tx, &vdb, payload); err != nil {
		return errors.Wrap(err, `failed to execute PreUpdateFromPayloadHook`)
	}

	if err := v.populateRowForUpdate(&vdb, payload); err != nil {
		return errors.Wrap(err, `failed to populate row data`)
	}

	if err := v.Update(tx, &vdb); err != nil {
		return errors.Wrap(err, `failed to update row in database`)
	}

	ls := LocalizedString()
	if err := ls.UpdateFields(tx, "User", vdb.EID, payload.LocalizedFields); err != nil {
		return errors.Wrap(err, `failed to update localized fields`)
	}
	return nil
}

func (v *UserSvc) ReplaceL10NStrings(tx *db.Tx, m *model.User, lang string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.ReplaceL10NStrings lang = %s", lang)
		defer g.End()
	}
	ls := LocalizedString()
	list := make([]db.LocalizedString, 0, 2)
	switch lang {
	case "", "en":
		if len(m.FirstName) > 0 && len(m.LastName) > 0 {
			return nil
		}
		for _, extralang := range []string{`ja`} {
			list = list[:0]
			if err := ls.LookupFields(tx, "User", m.ID, extralang, &list); err != nil {
				return errors.Wrap(err, `failed to lookup localized fields`)
			}

			for _, l := range list {
				switch l.Name {
				case "first_name":
					if len(m.FirstName) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'first_name' (fallback en -> %s", l.Language)
						}
						m.FirstName = l.Localized
					}
				case "last_name":
					if len(m.LastName) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'last_name' (fallback en -> %s", l.Language)
						}
						m.LastName = l.Localized
					}
				}
			}
		}
		return nil
	case "all":
		for _, extralang := range []string{`ja`} {
			list = list[:0]
			if err := ls.LookupFields(tx, "User", m.ID, extralang, &list); err != nil {
				return errors.Wrap(err, `failed to lookup localized fields`)
			}

			for _, l := range list {
				if pdebug.Enabled {
					pdebug.Printf("Adding key '%s#%s'", l.Name, l.Language)
				}
				m.LocalizedFields.Set(l.Language, l.Name, l.Localized)
			}
		}
	default:
		for _, extralang := range []string{`ja`} {
			list = list[:0]
			if err := ls.LookupFields(tx, "User", m.ID, extralang, &list); err != nil {
				return errors.Wrap(err, `failed to lookup localized fields`)
			}

			for _, l := range list {
				switch l.Name {
				case "first_name":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'first_name'")
					}
					m.FirstName = l.Localized
				case "last_name":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'last_name'")
					}
					m.LastName = l.Localized
				}
			}
		}
	}
	return nil
}

func (v *UserSvc) Delete(tx *db.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("User.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.User{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return errors.Wrap(err, `failed to delete from database`)
	}
	c := Cache()
	key := c.Key("User", id)
	c.Delete(key)
	if pdebug.Enabled {
		pdebug.Printf(`CACHE DEL %s`, key)
	}
	if err := db.DeleteLocalizedStringsForParent(tx, id, "User"); err != nil {
		return errors.Wrap(err, `failed to delete localized strings`)
	}
	return nil
}
