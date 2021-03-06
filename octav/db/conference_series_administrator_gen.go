package db

// Automatically generated by gendb utility. DO NOT EDIT!

import (
	"bytes"
	"database/sql"
	"time"

	"github.com/builderscon/octav/octav/tools"
	"github.com/lestrrat/go-pdebug"
	"github.com/pkg/errors"
)

const ConferenceSeriesAdministratorStdSelectColumns = "conference_series_administrators.oid, conference_series_administrators.series_id, conference_series_administrators.user_id, conference_series_administrators.created_on, conference_series_administrators.modified_on"
const ConferenceSeriesAdministratorTable = "conference_series_administrators"

type ConferenceSeriesAdministratorList []ConferenceSeriesAdministrator

func (c *ConferenceSeriesAdministrator) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&c.OID, &c.SeriesID, &c.UserID, &c.CreatedOn, &c.ModifiedOn)
}

func init() {
	hooks = append(hooks, func() {
		stmt := tools.GetBuffer()
		defer tools.ReleaseBuffer(stmt)

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(ConferenceSeriesAdministratorTable)
		stmt.WriteString(` WHERE oid = ?`)
		library.Register("sqlConferenceSeriesAdministratorDeleteByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(ConferenceSeriesAdministratorTable)
		stmt.WriteString(` SET series_id = ?, user_id = ? WHERE oid = ?`)
		library.Register("sqlConferenceSeriesAdministratorUpdateByOIDKey", stmt.String())
	})
}

func (c *ConferenceSeriesAdministrator) Create(tx *Tx, opts ...InsertOption) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("db.ConferenceSeriesAdministrator.Create").BindError(&err)
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
	stmt.WriteString(ConferenceSeriesAdministratorTable)
	stmt.WriteString(` (series_id, user_id, created_on, modified_on) VALUES (?, ?, ?, ?)`)
	result, err := tx.Exec(stmt.String(), c.SeriesID, c.UserID, c.CreatedOn, c.ModifiedOn)
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

func (c ConferenceSeriesAdministrator) Update(tx *Tx) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker(`ConferenceSeriesAdministrator.Update`).BindError(&err)
		defer g.End()
	}
	if c.OID != 0 {
		if pdebug.Enabled {
			pdebug.Printf(`Using OID (%d) as key`, c.OID)
		}
		stmt, err := library.GetStmt("sqlConferenceSeriesAdministratorUpdateByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.SeriesID, c.UserID, c.OID)
		return err
	}
	return errors.New("OID must be filled")
}

func (c ConferenceSeriesAdministrator) Delete(tx *Tx) error {
	if c.OID != 0 {
		stmt, err := library.GetStmt("sqlConferenceSeriesAdministratorDeleteByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.OID)
		return err
	}

	return errors.New("column OID must be filled")
}

func (v *ConferenceSeriesAdministratorList) FromRows(rows *sql.Rows, capacity int) error {
	var res []ConferenceSeriesAdministrator
	if capacity > 0 {
		res = make([]ConferenceSeriesAdministrator, 0, capacity)
	} else {
		res = []ConferenceSeriesAdministrator{}
	}

	for rows.Next() {
		vdb := ConferenceSeriesAdministrator{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}
