package dao

import (
	"database/sql"
)

type Guest struct {
	SessionId string
	ExtraInfo string
}

func AddGuestSessionDetail(db *sql.DB, session_id, extra_info string) (bool, error) {
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
