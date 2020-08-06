package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "log"
)

type DBManager struct {
	DB *sql.DB
}

func New() *DBManager{
	dbManager := &DBManager{

	}
	dbManager.initDB()


	return dbManager
}


func (dbManager *DBManager) initDB(){

	db, err := sql.Open("mysql", "root:fghjkl;'@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	//defer db.Close()


	_,err = db.Exec("USE everything")
	if err != nil {
		//panic(err)
	}

	articleDBInit(dbManager.DB)
	userDBInit(dbManager.DB)
	userArticleInit(dbManager.DB)
	dbManager.DB = db

	//userDBInsertTestData(db)
	//userDBSelectAll(db)
	userDBInsertTestData2(db)
}



func (dbManager *DBManager) get() {

}

func (dbManager *DBManager) Dealloc(){
	dbManager.DB.Close()

}
