package db

// Automatically generated by gendb utility. DO NOT EDIT!

import (
	"bytes"
	"database/sql"
	"errors"
)

const ConferenceDateStdSelectColumns = "conference_dates.oid, conference_dates.conference_id, conference_dates.date, conference_dates.open, conference_dates.close"
const ConferenceDateTable = "conference_dates"

type ConferenceDateList []ConferenceDate

func (c *ConferenceDate) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&c.OID, &c.ConferenceID, &c.Date, &c.Open, &c.Close)
}

func (c *ConferenceDate) Create(tx *Tx, opts ...InsertOption) error {
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
	stmt.WriteString(ConferenceDateTable)
	stmt.WriteString(` (conference_id, date, open, close) VALUES (?, ?, ?, ?)`)
	result, err := tx.Exec(stmt.String(), c.ConferenceID, c.Date, c.Open, c.Close)
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

func (c ConferenceDate) Update(tx *Tx) error {
	if c.OID != 0 {
		_, err := tx.Exec(`UPDATE `+ConferenceDateTable+` SET conference_id = ?, date = ?, open = ?, close = ? WHERE oid = ?`, c.ConferenceID, c.Date, c.Open, c.Close, c.OID)
		return err
	}
	return errors.New("either OID/EID must be filled")
}

func (c ConferenceDate) Delete(tx *Tx) error {
	if c.OID != 0 {
		_, err := tx.Exec(`DELETE FROM `+ConferenceDateTable+` WHERE oid = ?`, c.OID)
		return err
	}

	return errors.New("column OID must be filled")
}

func (v *ConferenceDateList) FromRows(rows *sql.Rows, capacity int) error {
	var res []ConferenceDate
	if capacity > 0 {
		res = make([]ConferenceDate, 0, capacity)
	} else {
		res = []ConferenceDate{}
	}

	for rows.Next() {
		vdb := ConferenceDate{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}