package service

import (
	"net/http"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/internal/errors"
	"github.com/builderscon/octav/octav/model"
	"github.com/builderscon/octav/octav/tools"
	pdebug "github.com/lestrrat/go-pdebug"
)

func (v *ClientSvc) Init() {
}

func (v *ClientSvc) populateRowForCreate(vdb *db.Client, payload *model.CreateClientRequest) error {
	vdb.EID = tools.RandomString(64)
	vdb.Secret = tools.RandomString(64)
	vdb.Name = payload.Name
	return nil
}

func (v *ClientSvc) populateRowForUpdate(vdb *db.Client, payload *model.UpdateClientRequest) error {
	vdb.Secret = payload.Secret
	vdb.Name = payload.Name
	return nil
}

func (v *ClientSvc) Authenticate(clientID, clientSecret string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Client.Authenticate").BindError(&err)
		defer g.End()
	}

	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}
	defer tx.AutoRollback()

	vdb := db.Client{}
	if err := vdb.LoadByEID(tx, clientID); err != nil {
		return errors.Wrap(err, "failed to load client ID")
	}

	if vdb.Secret != clientSecret {
		return errors.WithHTTPCode(errors.New("invalid secret"), http.StatusForbidden)
	}
	return nil
}
