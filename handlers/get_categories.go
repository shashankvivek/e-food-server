package handlers

import (
	"database/sql"
	"e-food/restapi/operations/menu"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type BroadCategory struct {
	BroadCategoryId string
	Name            string
	Description     string
	ImageUrl        string
	IsActive        bool
}

type Menu []BroadCategory

type menuImpl struct {
	dbClient *sql.DB
}

func NewMenuCategoryHandler(dbClient *sql.DB) menu.CategoryListHandler {
	return &menuImpl{
		dbClient: dbClient,
	}
}

func (impl *menuImpl) Handle(param menu.CategoryListParams) middleware.Responder {
	var q = "SELECT * FROM ecommerce.broadcategories"
	rows, err := impl.dbClient.Query(q)
	if err != nil {
		fmt.Println("Error fetching broad category")
	}
	defer rows.Close()
	retVal := Menu{}
	cols, _ := rows.Columns()
	fmt.Println("Col detected: ", cols)

	for rows.Next() {
		cat := BroadCategory{}
		err = rows.Scan(&cat.BroadCategoryId, &cat.Name, &cat.Description, &cat.ImageUrl, &cat.IsActive)
		fmt.Println(cat)
		if err != nil {
			fmt.Println("Error in rows")
		}
		retVal = append(retVal, cat)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error in rows asda")
	}
	fmt.Println(retVal)

	return menu.NewCategoryListOK().WithPayload(nil)
}
