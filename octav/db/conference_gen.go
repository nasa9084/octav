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

const ConferenceStdSelectColumns = "conferences.oid, conferences.eid, conferences.cover_url, conferences.redirect_url, conferences.series_id, conferences.slug, conferences.status, conferences.sub_title, conferences.title, conferences.blog_feedback_available, conferences.timetable_available, conferences.timezone, conferences.created_by, conferences.created_on, conferences.modified_on"
const ConferenceTable = "conferences"

type ConferenceList []Conference

func (c *Conference) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&c.OID, &c.EID, &c.CoverURL, &c.RedirectURL, &c.SeriesID, &c.Slug, &c.Status, &c.SubTitle, &c.Title, &c.BlogFeedbackAvailable, &c.TimetableAvailable, &c.Timezone, &c.CreatedBy, &c.CreatedOn, &c.ModifiedOn)
}

func init() {
	hooks = append(hooks, func() {
		stmt := tools.GetBuffer()
		defer tools.ReleaseBuffer(stmt)

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(ConferenceTable)
		stmt.WriteString(` WHERE oid = ?`)
		library.Register("sqlConferenceDeleteByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(ConferenceTable)
		stmt.WriteString(` SET eid = ?, cover_url = ?, redirect_url = ?, series_id = ?, slug = ?, status = ?, sub_title = ?, title = ?, blog_feedback_available = ?, timetable_available = ?, timezone = ?, created_by = ? WHERE oid = ?`)
		library.Register("sqlConferenceUpdateByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`SELECT `)
		stmt.WriteString(ConferenceStdSelectColumns)
		stmt.WriteString(` FROM `)
		stmt.WriteString(ConferenceTable)
		stmt.WriteString(` WHERE `)
		stmt.WriteString(ConferenceTable)
		stmt.WriteString(`.eid = ?`)
		library.Register("sqlConferenceLoadByEIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(ConferenceTable)
		stmt.WriteString(` WHERE eid = ?`)
		library.Register("sqlConferenceDeleteByEIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(ConferenceTable)
		stmt.WriteString(` SET eid = ?, cover_url = ?, redirect_url = ?, series_id = ?, slug = ?, status = ?, sub_title = ?, title = ?, blog_feedback_available = ?, timetable_available = ?, timezone = ?, created_by = ? WHERE eid = ?`)
		library.Register("sqlConferenceUpdateByEIDKey", stmt.String())
	})
}

func (c *Conference) LoadByEID(tx *Tx, eid string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker(`Conference.LoadByEID %s`, eid).BindError(&err)
		defer g.End()
	}
	stmt, err := library.GetStmt("sqlConferenceLoadByEIDKey")
	if err != nil {
		return errors.Wrap(err, `failed to get statement`)
	}
	row := tx.Stmt(stmt).QueryRow(eid)
	if err := c.Scan(row); err != nil {
		return err
	}
	return nil
}

func (c *Conference) Create(tx *Tx, opts ...InsertOption) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("db.Conference.Create").BindError(&err)
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
	stmt.WriteString(ConferenceTable)
	stmt.WriteString(` (eid, cover_url, redirect_url, series_id, slug, status, sub_title, title, blog_feedback_available, timetable_available, timezone, created_by, created_on, modified_on) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	result, err := tx.Exec(stmt.String(), c.EID, c.CoverURL, c.RedirectURL, c.SeriesID, c.Slug, c.Status, c.SubTitle, c.Title, c.BlogFeedbackAvailable, c.TimetableAvailable, c.Timezone, c.CreatedBy, c.CreatedOn, c.ModifiedOn)
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

func (c Conference) Update(tx *Tx) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker(`Conference.Update`).BindError(&err)
		defer g.End()
	}
	if c.OID != 0 {
		if pdebug.Enabled {
			pdebug.Printf(`Using OID (%d) as key`, c.OID)
		}
		stmt, err := library.GetStmt("sqlConferenceUpdateByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.EID, c.CoverURL, c.RedirectURL, c.SeriesID, c.Slug, c.Status, c.SubTitle, c.Title, c.BlogFeedbackAvailable, c.TimetableAvailable, c.Timezone, c.CreatedBy, c.OID)
		return err
	}
	if c.EID != "" {
		if pdebug.Enabled {
			pdebug.Printf(`Using EID (%s) as key`, c.EID)
		}
		stmt, err := library.GetStmt("sqlConferenceUpdateByEIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.EID, c.CoverURL, c.RedirectURL, c.SeriesID, c.Slug, c.Status, c.SubTitle, c.Title, c.BlogFeedbackAvailable, c.TimetableAvailable, c.Timezone, c.CreatedBy, c.EID)
		return err
	}
	return errors.New("either OID/EID must be filled")
}

func (c Conference) Delete(tx *Tx) error {
	if c.OID != 0 {
		stmt, err := library.GetStmt("sqlConferenceDeleteByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.OID)
		return err
	}

	if c.EID != "" {
		stmt, err := library.GetStmt("sqlConferenceDeleteByEIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.EID)
		return err
	}

	return errors.New("either OID/EID must be filled")
}

func (v *ConferenceList) FromRows(rows *sql.Rows, capacity int) error {
	var res []Conference
	if capacity > 0 {
		res = make([]Conference, 0, capacity)
	} else {
		res = []Conference{}
	}

	for rows.Next() {
		vdb := Conference{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}

func (v *ConferenceList) LoadSinceEID(tx *Tx, since string, limit int) error {
	var s int64
	if id := since; id != "" {
		vdb := Conference{}
		if err := vdb.LoadByEID(tx, id); err != nil {
			return err
		}

		s = vdb.OID
	}
	return v.LoadSince(tx, s, limit)
}

func (v *ConferenceList) LoadSince(tx *Tx, since int64, limit int) error {
	rows, err := tx.Query(`SELECT `+ConferenceStdSelectColumns+` FROM `+ConferenceTable+` WHERE conferences.oid > ? ORDER BY oid ASC LIMIT `+strconv.Itoa(limit), since)
	if err != nil {
		return err
	}

	if err := v.FromRows(rows, limit); err != nil {
		return err
	}
	return nil
}
