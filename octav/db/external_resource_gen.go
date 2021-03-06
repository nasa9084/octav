package db

// Automatically generated by gendb utility. DO NOT EDIT!

import (
	"bytes"
	"database/sql"
	"strconv"

	"github.com/builderscon/octav/octav/tools"
	"github.com/lestrrat/go-pdebug"
	"github.com/pkg/errors"
)

const ExternalResourceStdSelectColumns = "external_resources.oid, external_resources.eid, external_resources.conference_id, external_resources.description, external_resources.image_url, external_resources.title, external_resources.url, external_resources.sort_order"
const ExternalResourceTable = "external_resources"

type ExternalResourceList []ExternalResource

func (e *ExternalResource) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&e.OID, &e.EID, &e.ConferenceID, &e.Description, &e.ImageURL, &e.Title, &e.URL, &e.SortOrder)
}

func init() {
	hooks = append(hooks, func() {
		stmt := tools.GetBuffer()
		defer tools.ReleaseBuffer(stmt)

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(ExternalResourceTable)
		stmt.WriteString(` WHERE oid = ?`)
		library.Register("sqlExternalResourceDeleteByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(ExternalResourceTable)
		stmt.WriteString(` SET eid = ?, conference_id = ?, description = ?, image_url = ?, title = ?, url = ?, sort_order = ? WHERE oid = ?`)
		library.Register("sqlExternalResourceUpdateByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`SELECT `)
		stmt.WriteString(ExternalResourceStdSelectColumns)
		stmt.WriteString(` FROM `)
		stmt.WriteString(ExternalResourceTable)
		stmt.WriteString(` WHERE `)
		stmt.WriteString(ExternalResourceTable)
		stmt.WriteString(`.eid = ?`)
		library.Register("sqlExternalResourceLoadByEIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(ExternalResourceTable)
		stmt.WriteString(` WHERE eid = ?`)
		library.Register("sqlExternalResourceDeleteByEIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(ExternalResourceTable)
		stmt.WriteString(` SET eid = ?, conference_id = ?, description = ?, image_url = ?, title = ?, url = ?, sort_order = ? WHERE eid = ?`)
		library.Register("sqlExternalResourceUpdateByEIDKey", stmt.String())
	})
}

func (e *ExternalResource) LoadByEID(tx *Tx, eid string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker(`ExternalResource.LoadByEID %s`, eid).BindError(&err)
		defer g.End()
	}
	stmt, err := library.GetStmt("sqlExternalResourceLoadByEIDKey")
	if err != nil {
		return errors.Wrap(err, `failed to get statement`)
	}
	row := tx.Stmt(stmt).QueryRow(eid)
	if err := e.Scan(row); err != nil {
		return err
	}
	return nil
}

func (e *ExternalResource) Create(tx *Tx, opts ...InsertOption) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("db.ExternalResource.Create").BindError(&err)
		defer g.End()
		pdebug.Printf("%#v", e)
	}
	if e.EID == "" {
		return errors.New("create: non-empty EID required")
	}

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
	stmt.WriteString(ExternalResourceTable)
	stmt.WriteString(` (eid, conference_id, description, image_url, title, url, sort_order) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	result, err := tx.Exec(stmt.String(), e.EID, e.ConferenceID, e.Description, e.ImageURL, e.Title, e.URL, e.SortOrder)
	if err != nil {
		return err
	}

	lii, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.OID = lii
	return nil
}

func (e ExternalResource) Update(tx *Tx) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker(`ExternalResource.Update`).BindError(&err)
		defer g.End()
	}
	if e.OID != 0 {
		if pdebug.Enabled {
			pdebug.Printf(`Using OID (%d) as key`, e.OID)
		}
		stmt, err := library.GetStmt("sqlExternalResourceUpdateByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(e.EID, e.ConferenceID, e.Description, e.ImageURL, e.Title, e.URL, e.SortOrder, e.OID)
		return err
	}
	if e.EID != "" {
		if pdebug.Enabled {
			pdebug.Printf(`Using EID (%s) as key`, e.EID)
		}
		stmt, err := library.GetStmt("sqlExternalResourceUpdateByEIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(e.EID, e.ConferenceID, e.Description, e.ImageURL, e.Title, e.URL, e.SortOrder, e.EID)
		return err
	}
	return errors.New("either OID/EID must be filled")
}

func (e ExternalResource) Delete(tx *Tx) error {
	if e.OID != 0 {
		stmt, err := library.GetStmt("sqlExternalResourceDeleteByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(e.OID)
		return err
	}

	if e.EID != "" {
		stmt, err := library.GetStmt("sqlExternalResourceDeleteByEIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(e.EID)
		return err
	}

	return errors.New("either OID/EID must be filled")
}

func (v *ExternalResourceList) FromRows(rows *sql.Rows, capacity int) error {
	var res []ExternalResource
	if capacity > 0 {
		res = make([]ExternalResource, 0, capacity)
	} else {
		res = []ExternalResource{}
	}

	for rows.Next() {
		vdb := ExternalResource{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}

func (v *ExternalResourceList) LoadSinceEID(tx *Tx, since string, limit int) error {
	var s int64
	if id := since; id != "" {
		vdb := ExternalResource{}
		if err := vdb.LoadByEID(tx, id); err != nil {
			return err
		}

		s = vdb.OID
	}
	return v.LoadSince(tx, s, limit)
}

func (v *ExternalResourceList) LoadSince(tx *Tx, since int64, limit int) error {
	rows, err := tx.Query(`SELECT `+ExternalResourceStdSelectColumns+` FROM `+ExternalResourceTable+` WHERE external_resources.oid > ? ORDER BY oid ASC LIMIT `+strconv.Itoa(limit), since)
	if err != nil {
		return err
	}

	if err := v.FromRows(rows, limit); err != nil {
		return err
	}
	return nil
}
