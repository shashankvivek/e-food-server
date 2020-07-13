package dao

import (
	"database/sql"
)

type GuestInfoHandler interface {
	AddGuestSessionDetail(db *sql.DB, session_id, extra_info string) (bool, error)
}

type guestInfo struct{}

func CreateGuestInfoHandler() GuestInfoHandler {
	return &guestInfo{}
}

func (g *guestInfo) AddGuestSessionDetail(db *sql.DB, session_id, extra_info string) (bool, error) {
	res, err := db.Exec("INSERT INTO guest values  (?,?)", session_id, extra_info)
	if err != nil {
		return false, err
	}
	insertedRow, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return insertedRow == 1, nil
}
