package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

type BroadCategoryEntity struct {
	BC_Id          string
	BC_Name        string
	BC_Description string
	BC_ImageUrl    string
	BC_IsActive    bool
}

// Insert and update can be used by sellers
func InsertMenuItem(broadCategory *BroadCategoryEntity) (string, error) {
	return "new_id", nil
}

func GetBroadCategoryList(dbClient *sql.DB) ([]BroadCategoryEntity, error) {
	q := fmt.Sprintf("SELECT * FROM  %s", BroadCategoryTable)
	rows, err := dbClient.Query(q)
	if err != nil {
		return nil, errors.New("error fetching broad category list")
	}
	defer rows.Close()
	var retVal []BroadCategoryEntity
	if err := rows.Err(); err != nil {
		return nil, errors.New("error in rows of  broad category entity ")
	}
	for rows.Next() {
		cat := BroadCategoryEntity{}
		err = rows.Scan(&cat.BC_Id, &cat.BC_Name, &cat.BC_Description, &cat.BC_ImageUrl, &cat.BC_IsActive)
		if err != nil {
			return nil, errors.New("error in scanning rows")
		}
		retVal = append(retVal, cat)
	}
	return retVal, nil
}
