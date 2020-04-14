package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"twirp/db_proxy/module"
)

var (
	PgSqlConn *gorm.DB
)

func init() {
	uri := os.Getenv("PGURI")
	if uri == "" {
		uri = "host=0.0.0.0 port=5432 user=postgres dbname=postgres password=root sslmode=disable"
		//log.Fatalln("pgsql uri is null")
	}

	// host=0.0.0.0 port=5432 user=postgres dbname=postgres password=root sslmode=disable
	// docker run --name postgres --restart=always -e POSTGRES_PASSWORD=root -d -p5432 postgres
	db, err := gorm.Open("postgres", uri)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.DB().Ping()
	if err != nil {
		log.Fatalln(err)
	}
	PgSqlConn = db

	err = PgSqlConn.AutoMigrate(&module.Url{}).Error
	if err != nil {
		log.Fatalln(err)
	}
}
