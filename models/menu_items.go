package models

import (
	"database/sql"
	"e-food/integration/mysql"
	"errors"
	"fmt"
)

type CategoryJoin struct {
	bcId       string
	bcName     string
	bcIsActive bool
	scId       string
	scName     string
	scIsActive bool
}

func GetMenuItems(dbClient *sql.DB) (Categories, error) {
	q := fmt.Sprintf("select b.bcID ,bcNAME,bcIsActive,scId,scName,scIsActive "+
		"from %s as B INNER JOIN %s as S where b.bcId = s.bcId", mysql.BroadCategoryTable, mysql.SubCategoryTable)
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
		err = rows.Scan(&cat.bcId, &cat.bcName, &cat.bcIsActive, &cat.scId, &cat.scName, &cat.scIsActive)
		if err != nil {
			return nil, errors.New("error in scanning rows")
		}
		retVal = append(retVal, cat)
	}

	menu := getCategoriesInStructure(retVal)
	return menu, nil

}

func getCategoriesInStructure(categories []CategoryJoin) Categories {
	groupedMap := make(map[string]*CategoriesItems0)
	for _, category := range categories {
		if _, ok := groupedMap[category.bcId]; ok {
			// insert SC into it
			//groupedMap[category.bcId].BcID = "asd"
			groupedMap[category.bcId].SubCategories = append(groupedMap[category.bcId].SubCategories, &CategoriesItems0SubCategoriesItems0{
				ScID:       category.scId,
				ScName:     category.scName,
				ScIsActive: category.scIsActive,
				ScImageURL: "",
			})
		} else {
			// create a map entry
			groupedMap[category.bcId] = &CategoriesItems0{
				BcID:       category.bcId,
				BcName:     category.bcName,
				BcIsActive: category.bcIsActive,
				SubCategories: []*CategoriesItems0SubCategoriesItems0{
					{
						category.scId,
						"",
						category.scIsActive,
						category.scName,
					},
				},
			}
		}
	}
	var menuItem Categories
	for _, v := range groupedMap {
		menuItem = append(menuItem, v)
	}
	return menuItem
}
