package service

import (
	"context"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/model"
	"github.com/builderscon/octav/octav/tools"
	"github.com/pkg/errors"
)

func (v *FeaturedSpeakerSvc) Init() {}

func (v *FeaturedSpeakerSvc) populateRowForCreate(vdb *db.FeaturedSpeaker, payload *model.CreateFeaturedSpeakerRequest) error {
	vdb.EID = tools.UUID()
	vdb.ConferenceID = payload.ConferenceID
	vdb.DisplayName = payload.DisplayName
	vdb.Description = payload.Description

	if payload.AvatarURL.Valid() {
		vdb.AvatarURL.Valid = true
		vdb.AvatarURL.String = payload.AvatarURL.String
	}

	if payload.SpeakerID.Valid() {
		vdb.SpeakerID.Valid = true
		vdb.SpeakerID.String = payload.SpeakerID.String
	}

	return nil
}

func (v *FeaturedSpeakerSvc) populateRowForUpdate(vdb *db.FeaturedSpeaker, payload *model.UpdateFeaturedSpeakerRequest) error {
	if payload.DisplayName.Valid() {
		vdb.DisplayName = payload.DisplayName.String
	}

	if payload.Description.Valid() {
		vdb.Description = payload.Description.String
	}

	if payload.SpeakerID.Valid() {
		vdb.SpeakerID.Valid = true
		vdb.SpeakerID.String = payload.SpeakerID.String
	}

	if payload.AvatarURL.Valid() {
		vdb.AvatarURL.Valid = true
		vdb.AvatarURL.String = payload.AvatarURL.String
	}

	return nil
}

func (v *FeaturedSpeakerSvc) CreateFromPayload(tx *db.Tx, payload *model.AddFeaturedSpeakerRequest, result *model.FeaturedSpeaker) error {
	su := User()
	if err := su.IsConferenceAdministrator(tx, payload.ConferenceID, payload.UserID); err != nil {
		return errors.Wrap(err, "creating a featured speaker requires conference administrator privilege")
	}

	vdb := db.FeaturedSpeaker{}
	if err := v.Create(tx, &vdb, &model.CreateFeaturedSpeakerRequest{payload}); err != nil {
		return errors.Wrap(err, "failed to store in database")
	}

	c := model.FeaturedSpeaker{}
	if err := c.FromRow(vdb); err != nil {
		return errors.Wrap(err, "failed to populate model from database")
	}

	*result = c
	return nil
}

func (v *FeaturedSpeakerSvc) PreUpdateFromPayloadHook(ctx context.Context, tx *db.Tx, vdb *db.FeaturedSpeaker, payload *model.UpdateFeaturedSpeakerRequest) (err error) {
	su := User()
	if err := su.IsConferenceAdministrator(tx, vdb.ConferenceID, payload.UserID); err != nil {
		return errors.Wrap(err, "updating a featured speaker requires conference administrator privilege")
	}
	return nil
}

func (v *FeaturedSpeakerSvc) DeleteFromPayload(tx *db.Tx, payload *model.DeleteFeaturedSpeakerRequest) error {
	var m db.FeaturedSpeaker
	if err := m.LoadByEID(tx, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load featured speaker from database")
	}

	su := User()
	if err := su.IsConferenceAdministrator(tx, m.ConferenceID, payload.UserID); err != nil {
		return errors.Wrap(err, "deleting venues require administrator privileges")
	}

	return errors.Wrap(v.Delete(tx, m.EID), "failed to delete from database")
}

func (v *FeaturedSpeakerSvc) ListFromPayload(tx *db.Tx, result *model.FeaturedSpeakerList, payload *model.ListFeaturedSpeakersRequest) error {
	var vdbl db.FeaturedSpeakerList
	if err := vdbl.LoadByConferenceSinceEID(tx, payload.ConferenceID, payload.Since.String, int(payload.Limit.Int)); err != nil {
		return errors.Wrap(err, "failed to load featured speakers from database")
	}

	l := make(model.FeaturedSpeakerList, len(vdbl))
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

func (v *FeaturedSpeakerSvc) Decorate(tx *db.Tx, speaker *model.FeaturedSpeaker, trustedCall bool, lang string) error {
	if lang == "" {
		return nil
	}

	if err := v.ReplaceL10NStrings(tx, speaker, lang); err != nil {
		return errors.Wrap(err, "failed to replace L10N strings")
	}

	return nil
}
