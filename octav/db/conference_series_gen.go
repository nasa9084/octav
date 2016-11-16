package db

// Automatically generated by gendb utility. DO NOT EDIT!

import (
	"bytes"
	"database/sql"
	"strconv"
	"time"

	"github.com/builderscon/octav/octav/tools"
	"github.com/lestrrat/go-pdebug"
	"github.com/pkg/errors"
)

const ConferenceSeriesStdSelectColumns = "conference_series.oid, conference_series.eid, conference_series.slug, conference_series.title, conference_series.created_on, conference_series.modified_on"
const ConferenceSeriesTable = "conference_series"

type ConferenceSeriesList []ConferenceSeries

func (c *ConferenceSeries) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&c.OID, &c.EID, &c.Slug, &c.Title, &c.CreatedOn, &c.ModifiedOn)
}

func init() {
	hooks = append(hooks, func() {
		stmt := tools.GetBuffer()
		defer tools.ReleaseBuffer(stmt)

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(ConferenceSeriesTable)
		stmt.WriteString(` WHERE oid = ?`)
		library.Register("sqlConferenceSeriesDeleteByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(ConferenceSeriesTable)
		stmt.WriteString(` SET eid = ?, slug = ?, title = ? WHERE oid = ?`)
		library.Register("sqlConferenceSeriesUpdateByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`SELECT `)
		stmt.WriteString(ConferenceSeriesStdSelectColumns)
		stmt.WriteString(` FROM `)
		stmt.WriteString(ConferenceSeriesTable)
		stmt.WriteString(` WHERE `)
		stmt.WriteString(ConferenceSeriesTable)
		stmt.WriteString(`.eid = ?`)
		library.Register("sqlConferenceSeriesLoadByEIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(ConferenceSeriesTable)
		stmt.WriteString(` WHERE eid = ?`)
		library.Register("sqlConferenceSeriesDeleteByEIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(ConferenceSeriesTable)
		stmt.WriteString(` SET eid = ?, slug = ?, title = ? WHERE eid = ?`)
		library.Register("sqlConferenceSeriesUpdateByEIDKey", stmt.String())
	})
}

func (c *ConferenceSeries) LoadByEID(tx *Tx, eid string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker(`ConferenceSeries.LoadByEID %s`, eid).BindError(&err)
		defer g.End()
	}
	stmt, err := library.GetStmt("sqlConferenceSeriesLoadByEIDKey")
	if err != nil {
		return errors.Wrap(err, `failed to get statement`)
	}
	row := tx.Stmt(stmt).QueryRow(eid)
	if err := c.Scan(row); err != nil {
		return err
	}
	return nil
}

func (c *ConferenceSeries) Create(tx *Tx, opts ...InsertOption) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("db.ConferenceSeries.Create").BindError(&err)
		defer g.End()
		pdebug.Printf("%#v", c)
	}
	if c.EID == "" {
		return errors.New("create: non-empty EID required")
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
	stmt.WriteString(ConferenceSeriesTable)
	stmt.WriteString(` (eid, slug, title, created_on, modified_on) VALUES (?, ?, ?, ?, ?)`)
	result, err := tx.Exec(stmt.String(), c.EID, c.Slug, c.Title, c.CreatedOn, c.ModifiedOn)
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

func (c ConferenceSeries) Update(tx *Tx) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker(`ConferenceSeries.Update`).BindError(&err)
		defer g.End()
	}
	if c.OID != 0 {
		if pdebug.Enabled {
			pdebug.Printf(`Using OID (%d) as key`, c.OID)
		}
		stmt, err := library.GetStmt("sqlConferenceSeriesUpdateByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.EID, c.Slug, c.Title, c.OID)
		return err
	}
	if c.EID != "" {
		if pdebug.Enabled {
			pdebug.Printf(`Using EID (%s) as key`, c.EID)
		}
		stmt, err := library.GetStmt("sqlConferenceSeriesUpdateByEIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.EID, c.Slug, c.Title, c.EID)
		return err
	}
	return errors.New("either OID/EID must be filled")
}

func (c ConferenceSeries) Delete(tx *Tx) error {
	if c.OID != 0 {
		stmt, err := library.GetStmt("sqlConferenceSeriesDeleteByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.OID)
		return err
	}

	if c.EID != "" {
		stmt, err := library.GetStmt("sqlConferenceSeriesDeleteByEIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.EID)
		return err
	}

	return errors.New("either OID/EID must be filled")
}

func (v *ConferenceSeriesList) FromRows(rows *sql.Rows, capacity int) error {
	var res []ConferenceSeries
	if capacity > 0 {
		res = make([]ConferenceSeries, 0, capacity)
	} else {
		res = []ConferenceSeries{}
	}

	for rows.Next() {
		vdb := ConferenceSeries{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}

func (v *ConferenceSeriesList) LoadSinceEID(tx *Tx, since string, limit int) error {
	var s int64
	if id := since; id != "" {
		vdb := ConferenceSeries{}
		if err := vdb.LoadByEID(tx, id); err != nil {
			return err
		}

		s = vdb.OID
	}
	return v.LoadSince(tx, s, limit)
}

func (v *ConferenceSeriesList) LoadSince(tx *Tx, since int64, limit int) error {
	rows, err := tx.Query(`SELECT `+ConferenceSeriesStdSelectColumns+` FROM `+ConferenceSeriesTable+` WHERE conference_series.oid > ? ORDER BY oid ASC LIMIT `+strconv.Itoa(limit), since)
	if err != nil {
		return err
	}

	if err := v.FromRows(rows, limit); err != nil {
		return err
	}
	return nil
}
