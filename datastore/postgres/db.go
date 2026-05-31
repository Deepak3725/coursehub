package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgres_host     = "db"
	postgres_port     = 5432
	postgres_user     = "postgres"
	postgres_password = "postgres"
	postgres_dbname   = "postgres"
)

// initializing sql driver
var Db *sql.DB

// init is called before main
func init() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		//                             ↑          ↑          ↑              ↑
		//                           spaces    spaces      spaces         spaces  ✅
		postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)
	var err error
	Db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	} else {
		log.Println("Database connected Successfully")
	}
}
