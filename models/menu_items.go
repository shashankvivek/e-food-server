package models

import (
	"database/sql"
	"e-food/integration/mysql"
	"errors"
	"fmt"
)

type CategoryJoin struct {
	BC_Id       string
	BC_Name     string
	BC_IsActive bool
	SC_Id       string
	SC_Name     string
	SC_IsActive bool
}

type MenuItems struct {
	BC_Id         string
	BC_Name       string
	BC_IsActive   bool
	SubCategories []SubMenuItems
}

type SubMenuItems struct {
	SC_Id       string
	SC_Name     string
	SC_IsActive bool
}

func GetMenuItems(dbClient *sql.DB) (*[]CategoryJoin, error) {
	q := fmt.Sprintf("select b.BC_ID ,BC_NAME,BC_IsActive,SC_Id,SC_Name,SC_IsActive "+
		"from %s as B INNER JOIN %s as S where b.BC_Id = s.BC_Id", mysql.BroadCategoryTable, mysql.SubCategoryTable)
	fmt.Println(q)
	rows, err := dbClient.Query(q)
	if err != nil {
		return nil, errors.New("error fetching broad category list")
	}
	defer rows.Close()
	var retVal []CategoryJoin
	if err := rows.Err(); err != nil {
		return nil, errors.New("error in rows of  broad category entity ")
	}
	for rows.Next() {
		cat := CategoryJoin{}
		err = rows.Scan(&cat.BC_Id, &cat.BC_Name, &cat.BC_IsActive, &cat.SC_Id, &cat.SC_Name, &cat.SC_IsActive)
		if err != nil {
			return nil, errors.New("error in scanning rows")
		}
		retVal = append(retVal, cat)
	}
	return &retVal, nil
}
