package clients

import (
	"database/sql"
	"e-food/constants"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/go-sql-driver/mysql"
	razorpay "github.com/razorpay/razorpay-go"
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
	// move values to constants
	cfg := mysql.Cfg("qwiklabs-gcp-00-7f4ecd010334:us-central1:e-db", "root", "root")
	cfg.DBName = "e_db"
	db, err := mysql.DialCfg(cfg)
	//db, err := sql.Open(constants.DriveName, "root:root@/ecommerce?parseTime=True")
	//TestCloudSQL()
	if err != nil {
		log.Fatal("error connecting DB : ", err.Error())
	}
	return db
}

func TestCloudSQL() {
	cfg := mysql.Cfg("qwiklabs-gcp-00-7f4ecd010334:us-central1:e-db", "root", "root")
	cfg.DBName = "e_db"
	db, err := mysql.DialCfg(cfg)
	if err != nil {
		log.Println("Error in CloudSQL")
		log.Fatal(err)
	}
	log.Println(db)
}

func (b ClientBuilder) BuildRazorPayClient() *razorpay.Client {
	// move keys to secrets.json
	client := razorpay.NewClient(constants.MyRazorKey, constants.MyRazorSecret)
	return client
}
