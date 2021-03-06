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

type rawUser struct {
	ID         string `json:"id"`
	AuthVia    string `json:"auth_via,omitempty"`
	AuthUserID string `json:"auth_user_id,omitempty"`
	AvatarURL  string `json:"avatar_url,omitempty"`
	FirstName  string `json:"first_name,omitempty" l10n:"true"`
	LastName   string `json:"last_name,omitempty" l10n:"true"`
	Lang       string `json:"lang"`
	Nickname   string `json:"nickname"`
	Email      string `json:"email,omitempty"`
	TshirtSize string `json:"tshirt_size,omitempty"`
	IsAdmin    bool   `json:"is_admin"`
	Timezone   string `json:"timezone"`
}

func (v User) MarshalJSON() ([]byte, error) {
	var raw rawUser
	raw.ID = v.ID
	raw.AuthVia = v.AuthVia
	raw.AuthUserID = v.AuthUserID
	raw.AvatarURL = v.AvatarURL
	raw.FirstName = v.FirstName
	raw.LastName = v.LastName
	raw.Lang = v.Lang
	raw.Nickname = v.Nickname
	raw.Email = v.Email
	raw.TshirtSize = v.TshirtSize
	raw.IsAdmin = v.IsAdmin
	raw.Timezone = v.Timezone
	buf, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	return MarshalJSONWithL10N(buf, v.LocalizedFields)
}

func (v *User) Load(tx *db.Tx, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("model.User.Load %s", id).BindError(&err)
		defer g.End()
	}
	vdb := db.User{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(&vdb); err != nil {
		return err
	}
	return nil
}

func (v *User) FromRow(vdb *db.User) error {
	v.ID = vdb.EID
	v.AuthVia = vdb.AuthVia
	v.AuthUserID = vdb.AuthUserID
	if vdb.AvatarURL.Valid {
		v.AvatarURL = vdb.AvatarURL.String
	}
	if vdb.FirstName.Valid {
		v.FirstName = vdb.FirstName.String
	}
	if vdb.LastName.Valid {
		v.LastName = vdb.LastName.String
	}
	v.Lang = vdb.Lang
	v.Nickname = vdb.Nickname
	if vdb.Email.Valid {
		v.Email = vdb.Email.String
	}
	if vdb.TshirtSize.Valid {
		v.TshirtSize = vdb.TshirtSize.String
	}
	v.IsAdmin = vdb.IsAdmin
	v.Timezone = vdb.Timezone
	return nil
}

func (v *User) ToRow(vdb *db.User) error {
	vdb.EID = v.ID
	vdb.AuthVia = v.AuthVia
	vdb.AuthUserID = v.AuthUserID
	vdb.AvatarURL.Valid = true
	vdb.AvatarURL.String = v.AvatarURL
	vdb.FirstName.Valid = true
	vdb.FirstName.String = v.FirstName
	vdb.LastName.Valid = true
	vdb.LastName.String = v.LastName
	vdb.Lang = v.Lang
	vdb.Nickname = v.Nickname
	vdb.Email.Valid = true
	vdb.Email.String = v.Email
	vdb.TshirtSize.Valid = true
	vdb.TshirtSize.String = v.TshirtSize
	vdb.IsAdmin = v.IsAdmin
	vdb.Timezone = v.Timezone
	return nil
}
