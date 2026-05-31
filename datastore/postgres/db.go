package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgres_host     = "dpg-d8e1uv99rddc73efdp6g-a.singapore-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "courshub_db_user"
	postgres_password = "mZ7bowv2lDQVjJHrClK2H5tTiVkeCr35"
	postgres_dbname   = "courshub_db"
)

// initializing sql driver
var Db *sql.DB

// init is called before main
func init() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)
	var err error
	Db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	} else {
		log.Println("Database connected Successfully")
	}
}
