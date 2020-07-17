package database

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

type DBManager struct {

}

func New() *DBManager{
	db := &DBManager{

	}

	return db
}

var DB *sql.DB


func (db *DBManager) Init(){

	var err error
	DB, err = sql.Open("mysql", dataBase)
	if err != nil {
		log.Fatalln("open db fail:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalln("ping db fail:", err)
	}

	articleDBInit()
	userDBInit()
	userArticleInit()
}