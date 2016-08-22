package service

import (
	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/model"
	"github.com/builderscon/octav/octav/tools"
	"github.com/pkg/errors"
)

func (v *Venue) populateRowForCreate(vdb *db.Venue, payload model.CreateVenueRequest) error {
	vdb.EID = tools.UUID()
	vdb.Name = payload.Name.String
	vdb.Address = payload.Address.String
	vdb.Latitude = payload.Latitude.Float
	vdb.Longitude = payload.Longitude.Float
	return nil
}

func (v *Venue) populateRowForUpdate(vdb *db.Venue, payload model.UpdateVenueRequest) error {
	if payload.Name.Valid() {
	vdb.Name = payload.Name.String
	}

	if payload.Address.Valid() {
	vdb.Address = payload.Address.String
	}

	if payload.Latitude.Valid() {
		vdb.Latitude = payload.Latitude.Float
	}

	if payload.Longitude.Valid() {
		vdb.Longitude = payload.Longitude.Float
	}
	return nil
}

func (v *Venue) LoadRooms(tx *db.Tx, cdl *model.RoomList, venueID string) error {
	var vdbl db.RoomList
	if err := db.LoadVenueRooms(tx, &vdbl, venueID); err != nil {
		return err
	}

	res := make(model.RoomList, len(vdbl))
	for i, vdb := range vdbl {
		var u model.Room
		if err := u.FromRow(vdb); err != nil {
			return err
		}
		res[i] = u
	}
	*cdl = res
	return nil
}

func (v *Venue) Decorate(tx *db.Tx, venue *model.Venue, trustedCall bool, lang string) error {
	if err := v.LoadRooms(tx, &venue.Rooms, venue.ID); err != nil {
		return err
	}

	if lang != "" {
		sr := Room{}
		for i := range venue.Rooms {
			if err := sr.ReplaceL10NStrings(tx, &venue.Rooms[i], lang); err != nil {
				return errors.Wrap(err, "failed to replace L10N strings")
			}
		}

		if err := v.ReplaceL10NStrings(tx, venue, lang); err != nil {
			return errors.Wrap(err, "failed to replace L10N strings")
		}
	}

	return nil
}

func (v *Venue) CreateFromPayload(tx *db.Tx, venue *model.Venue, payload model.CreateVenueRequest) error {
	su := User{}
	if err := su.IsAdministrator(tx, payload.UserID); err != nil {
		return errors.Wrap(err, "creating venues require administrator privileges")
	}

	vdb := db.Venue{}
	if err := v.Create(tx, &vdb, payload); err != nil {
		return errors.Wrap(err, "failed to store in database")
	}

	r := model.Venue{}
	if err := r.FromRow(vdb); err != nil {
		return errors.Wrap(err, "failed to populate model from database")
	}
	*venue = r

	return nil
}

func (v *Venue) UpdateFromPayload(tx *db.Tx, venue *model.Venue, payload model.UpdateVenueRequest) error {
	vdb := db.Venue{}
	if err := vdb.LoadByEID(tx, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load from database")
	}

	// TODO: We must protect the API server from changing important
	// fields like conference_id, speaker_id, room_id, etc from regular
	// users, but allow administrators to do anything they want
	if err := v.Update(tx, &vdb, payload); err != nil {
		return errors.Wrap(err, "failed to update database")
	}

	r := model.Venue{}
	if err := r.FromRow(vdb); err != nil {
		return errors.Wrap(err, "failed to populate model from database")
	}
	*venue = r

	return nil
}

func (v *Venue) DeleteFromPayload(tx *db.Tx, payload model.DeleteVenueRequest) error {
	su := User{}
	if err := su.IsAdministrator(tx, payload.UserID); err != nil {
		return errors.Wrap(err, "deleting venues require administrator privileges")
	}

	return errors.Wrap(v.Delete(tx, payload.ID), "failed to delete from database")
}

func (v *Venue) ListFromPayload(tx *db.Tx, result *model.VenueList, payload model.ListVenueRequest) error {
	vdbl := db.VenueList{}
	if err := vdbl.LoadSinceEID(tx, payload.Since.String, int(payload.Limit.Int)); err != nil {
		return errors.Wrap(err, "failed to load from database")
	}

	l := make(model.VenueList, len(vdbl))
	for i, vdb := range vdbl {
		if err := (l[i]).FromRow(vdb); err != nil {
			return errors.Wrap(err, "failed to populate model from database")
		}

		if err := v.Decorate(tx, &l[i], payload.TrustedCall, payload.Lang.String); err != nil {
			return errors.Wrap(err, "failed to decorate venue with associated data")
		}
	}

	*result = l
	return nil
}

