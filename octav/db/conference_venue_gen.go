package db

// Automatically generated by gendb utility. DO NOT EDIT!

import (
	"bytes"
	"database/sql"
	"time"

	"github.com/builderscon/octav/octav/tools"
	"github.com/lestrrat/go-pdebug"
	"github.com/lestrrat/go-sqllib"
	"github.com/pkg/errors"
)

const ConferenceVenueStdSelectColumns = "conference_venues.oid, conference_venues.conference_id, conference_venues.venue_id, conference_venues.created_on, conference_venues.modified_on"
const ConferenceVenueTable = "conference_venues"

type ConferenceVenueList []ConferenceVenue

func (c *ConferenceVenue) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&c.OID, &c.ConferenceID, &c.VenueID, &c.CreatedOn, &c.ModifiedOn)
}

var sqlConferenceVenueUpdateByOIDKey sqllib.Key
var sqlConferenceVenueDeleteByOIDKey sqllib.Key

func init() {
	hooks = append(hooks, func() {
		stmt := tools.GetBuffer()
		defer tools.ReleaseBuffer(stmt)

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(ConferenceVenueTable)
		stmt.WriteString(` WHERE oid = ?`)
		sqlConferenceVenueDeleteByOIDKey = library.Register(stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(ConferenceVenueTable)
		stmt.WriteString(` SET conference_id = ?, venue_id = ? WHERE oid = ?`)
		sqlConferenceVenueUpdateByOIDKey = library.Register(stmt.String())
	})
}

func (c *ConferenceVenue) Create(tx *Tx, opts ...InsertOption) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("db.ConferenceVenue.Create").BindError(&err)
		defer g.End()
		pdebug.Printf("%#v", c)
	}
	c.CreatedOn = time.Now()
	doIgnore := false
	for _, opt := range opts {
		switch opt.(type) {
		case insertIgnoreOption:
			doIgnore = true
		}
	}

	stmt := bytes.Buffer{}
	stmt.WriteString("INSERT ")
	if doIgnore {
		stmt.WriteString("IGNORE ")
	}
	stmt.WriteString("INTO ")
	stmt.WriteString(ConferenceVenueTable)
	stmt.WriteString(` (conference_id, venue_id, created_on, modified_on) VALUES (?, ?, ?, ?)`)
	result, err := tx.Exec(stmt.String(), c.ConferenceID, c.VenueID, c.CreatedOn, c.ModifiedOn)
	if err != nil {
		return err
	}

	lii, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.OID = lii
	return nil
}

func (c ConferenceVenue) Update(tx *Tx) error {
	if c.OID != 0 {
		stmt, err := library.GetStmt(sqlConferenceVenueUpdateByOIDKey)
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.ConferenceID, c.VenueID, c.OID)
		return err
	}
	return errors.New("either OID/EID must be filled")
}

func (c ConferenceVenue) Delete(tx *Tx) error {
	if c.OID != 0 {
		stmt, err := library.GetStmt(sqlConferenceVenueDeleteByOIDKey)
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.OID)
		return err
	}

	return errors.New("column OID must be filled")
}

func (v *ConferenceVenueList) FromRows(rows *sql.Rows, capacity int) error {
	var res []ConferenceVenue
	if capacity > 0 {
		res = make([]ConferenceVenue, 0, capacity)
	} else {
		res = []ConferenceVenue{}
	}

	for rows.Next() {
		vdb := ConferenceVenue{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}
