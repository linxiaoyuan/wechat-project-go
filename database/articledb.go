package database

import (
	"database/sql"
	"fmt"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "log"
)


func articleDBInit(db *sql.DB)  {
	fmt.Println("Start Init Article DB")

}