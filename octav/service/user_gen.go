// Automatically generated by genmodel utility. DO NOT EDIT!
package service

import (
	"errors"
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/lestrrat/go-pdebug"
)

var _ = time.Time{}

// Create takes in the transaction, the incoming payload, and a reference to
// a database row. The database row is initialized/populated so that the
// caller can use it afterwards
func (v *User) Create(tx *db.Tx, vdb *db.User, payload CreateUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Create(tx); err != nil {
		return err
	}

	return nil
}

func (v *User) Update(tx *db.Tx, vdb *db.User, payload UpdateUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.Update (%s)", vdb.EID).BindError(&err)
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
	return nil
}

func (v *User) Delete(tx *db.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("User.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.User{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	return nil
}

func (v *User) LoadList(tx *db.Tx, vdbl *db.UserList, since string, limit int) error {
	return vdbl.LoadSinceEID(tx, since, limit)
}
