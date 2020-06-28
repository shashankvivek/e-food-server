package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

type BroadCategoryEntity struct {
	bcId          string
	bcName        string
	bcDescription string
	bcImageUrl    string
	bcIsActive    bool
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
		err = rows.Scan(&cat.bcId, &cat.bcName, &cat.bcDescription, &cat.bcImageUrl, &cat.bcIsActive)
		if err != nil {
			return nil, errors.New("error in scanning rows")
		}
		retVal = append(retVal, cat)
	}
	return retVal, nil
}
