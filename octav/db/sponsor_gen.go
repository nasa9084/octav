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

const SponsorStdSelectColumns = "sponsors.oid, sponsors.eid, sponsors.conference_id, sponsors.name, sponsors.logo_url1, sponsors.logo_url2, sponsors.logo_url3, sponsors.url, sponsors.group_name, sponsors.sort_order, sponsors.created_on, sponsors.modified_on"
const SponsorTable = "sponsors"

type SponsorList []Sponsor

func (s *Sponsor) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&s.OID, &s.EID, &s.ConferenceID, &s.Name, &s.LogoURL1, &s.LogoURL2, &s.LogoURL3, &s.URL, &s.GroupName, &s.SortOrder, &s.CreatedOn, &s.ModifiedOn)
}

func init() {
	hooks = append(hooks, func() {
		stmt := tools.GetBuffer()
		defer tools.ReleaseBuffer(stmt)

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(SponsorTable)
		stmt.WriteString(` WHERE oid = ?`)
		library.Register("sqlSponsorDeleteByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(SponsorTable)
		stmt.WriteString(` SET eid = ?, conference_id = ?, name = ?, logo_url1 = ?, logo_url2 = ?, logo_url3 = ?, url = ?, group_name = ?, sort_order = ? WHERE oid = ?`)
		library.Register("sqlSponsorUpdateByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`SELECT `)
		stmt.WriteString(SponsorStdSelectColumns)
		stmt.WriteString(` FROM `)
		stmt.WriteString(SponsorTable)
		stmt.WriteString(` WHERE `)
		stmt.WriteString(SponsorTable)
		stmt.WriteString(`.eid = ?`)
		library.Register("sqlSponsorLoadByEIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(SponsorTable)
		stmt.WriteString(` WHERE eid = ?`)
		library.Register("sqlSponsorDeleteByEIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(SponsorTable)
		stmt.WriteString(` SET eid = ?, conference_id = ?, name = ?, logo_url1 = ?, logo_url2 = ?, logo_url3 = ?, url = ?, group_name = ?, sort_order = ? WHERE eid = ?`)
		library.Register("sqlSponsorUpdateByEIDKey", stmt.String())
	})
}

func (s *Sponsor) LoadByEID(tx *Tx, eid string) error {
	stmt, err := library.GetStmt("sqlSponsorLoadByEIDKey")
	if err != nil {
		return errors.Wrap(err, `failed to get statement`)
	}
	row := tx.Stmt(stmt).QueryRow(eid)
	if err := s.Scan(row); err != nil {
		return err
	}
	return nil
}

func (s *Sponsor) Create(tx *Tx, opts ...InsertOption) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("db.Sponsor.Create").BindError(&err)
		defer g.End()
		pdebug.Printf("%#v", s)
	}
	if s.EID == "" {
		return errors.New("create: non-empty EID required")
	}

	s.CreatedOn = time.Now()
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
	stmt.WriteString(SponsorTable)
	stmt.WriteString(` (eid, conference_id, name, logo_url1, logo_url2, logo_url3, url, group_name, sort_order, created_on, modified_on) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	result, err := tx.Exec(stmt.String(), s.EID, s.ConferenceID, s.Name, s.LogoURL1, s.LogoURL2, s.LogoURL3, s.URL, s.GroupName, s.SortOrder, s.CreatedOn, s.ModifiedOn)
	if err != nil {
		return err
	}

	lii, err := result.LastInsertId()
	if err != nil {
		return err
	}

	s.OID = lii
	return nil
}

func (s Sponsor) Update(tx *Tx) error {
	if s.OID != 0 {
		stmt, err := library.GetStmt("sqlSponsorUpdateByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(s.EID, s.ConferenceID, s.Name, s.LogoURL1, s.LogoURL2, s.LogoURL3, s.URL, s.GroupName, s.SortOrder, s.OID)
		return err
	}
	if s.EID != "" {
		stmt, err := library.GetStmt("sqlSponsorUpdateByEIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(s.EID, s.ConferenceID, s.Name, s.LogoURL1, s.LogoURL2, s.LogoURL3, s.URL, s.GroupName, s.SortOrder, s.EID)
		return err
	}
	return errors.New("either OID/EID must be filled")
}

func (s Sponsor) Delete(tx *Tx) error {
	if s.OID != 0 {
		stmt, err := library.GetStmt("sqlSponsorDeleteByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(s.OID)
		return err
	}

	if s.EID != "" {
		stmt, err := library.GetStmt("sqlSponsorDeleteByEIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(s.EID)
		return err
	}

	return errors.New("either OID/EID must be filled")
}

func (v *SponsorList) FromRows(rows *sql.Rows, capacity int) error {
	var res []Sponsor
	if capacity > 0 {
		res = make([]Sponsor, 0, capacity)
	} else {
		res = []Sponsor{}
	}

	for rows.Next() {
		vdb := Sponsor{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}

func (v *SponsorList) LoadSinceEID(tx *Tx, since string, limit int) error {
	var s int64
	if id := since; id != "" {
		vdb := Sponsor{}
		if err := vdb.LoadByEID(tx, id); err != nil {
			return err
		}

		s = vdb.OID
	}
	return v.LoadSince(tx, s, limit)
}

func (v *SponsorList) LoadSince(tx *Tx, since int64, limit int) error {
	rows, err := tx.Query(`SELECT `+SponsorStdSelectColumns+` FROM `+SponsorTable+` WHERE sponsors.oid > ? ORDER BY oid ASC LIMIT `+strconv.Itoa(limit), since)
	if err != nil {
		return err
	}

	if err := v.FromRows(rows, limit); err != nil {
		return err
	}
	return nil
}
