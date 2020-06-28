package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

type BroadCategoryEntity struct {
	BroadCategoryId string
	Name            string
	Description     string
	ImageUrl        string
	IsActive        bool
}

// Insert and update can be used by sellers
func InsertMenuItem(broadCategory *BroadCategoryEntity) (string, error) {
	return "new_id", nil
}

func GetBroadCategoryList(dbClient *sql.DB) ([]BroadCategoryEntity, error) {
	var q = fmt.Sprintf("SELECT * FROM %s.%s", databaseSchema, broadCategoryTable)
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
		err = rows.Scan(&cat.BroadCategoryId, &cat.Name, &cat.Description, &cat.ImageUrl, &cat.IsActive)
		if err != nil {
			return nil, errors.New("error in scanning rows")
		}
		retVal = append(retVal, cat)
	}
	return retVal, nil
}
