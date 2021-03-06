package dao

import (
	"database/sql"
	"e-food/api/models"
	"e-food/constants"
	"fmt"
)

type MenuHandler interface {
	GetMenuItems(dbClient *sql.DB) (models.Categories, error)
	getCategoriesInStructure(categories []category) models.Categories
}

type category struct {
	bcId       int64
	bcName     string
	bcIsActive bool
	scId       int64
	scName     string
	scIsActive bool
}

func CreateMenuHandler() MenuHandler {
	return &category{}
}

func (c *category) GetMenuItems(dbClient *sql.DB) (models.Categories, error) {
	q := fmt.Sprintf("select b.bcID ,bcNAME,bcIsActive,scId,scName,scIsActive "+
		"from %s as B INNER JOIN %s as S where b.bcId = s.bcId", constants.BroadCategoryTable, constants.SubCategoryTable)
	fmt.Println(q)
	rows, err := dbClient.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var retVal []category
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		cat := category{}
		err = rows.Scan(&cat.bcId, &cat.bcName, &cat.bcIsActive, &cat.scId, &cat.scName, &cat.scIsActive)
		if err != nil {
			return nil, err
		}
		retVal = append(retVal, cat)
	}

	menu := c.getCategoriesInStructure(retVal)
	return menu, nil

}

func (c *category) getCategoriesInStructure(categories []category) models.Categories {
	groupedMap := make(map[int64]*models.Category)
	for _, category := range categories {
		if _, ok := groupedMap[category.bcId]; ok {
			// insert SC into it
			groupedMap[category.bcId].SubCategories = append(groupedMap[category.bcId].SubCategories, &models.SubCategory{
				ScID:       category.scId,
				ScName:     category.scName,
				ScIsActive: category.scIsActive,
				ScImageURL: "",
			})
		} else {
			// create a map entry
			groupedMap[category.bcId] = &models.Category{
				BcID:       category.bcId,
				BcName:     category.bcName,
				BcIsActive: category.bcIsActive,
				SubCategories: []*models.SubCategory{
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
	var menuItem models.Categories
	for _, v := range groupedMap {
		menuItem = append(menuItem, v)
	}
	return menuItem
}
