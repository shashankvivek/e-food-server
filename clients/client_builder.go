package clients

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type ClientBuilder struct {
	// this place can be used to initialize the auth client which can be used to talk to other micro-services
}

func NewClientBuilder() ClientBuilder {
	return ClientBuilder{}
}

func (b ClientBuilder) BuildSqlClient() *sql.DB {
	// The password can come from secrets.json of GKE
	db, err := sql.Open("mysql", "root:root@/ecommerce?parseTime=True")
	if err != nil {
		log.Fatal("error connecting DB : ", err.Error())
	}
	return db
}
