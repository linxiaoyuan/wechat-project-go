package database

import (
	_ "database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/jmoiron/sqlx"
	_ "log"
)

var schema = `
CREATE TABLE user (
    user_id int64,
	tel string,
	wechat_id string,
	register_time time.Time,
    user_name string,
	sex string,	
	age int,
    email text
)
`

type User struct {
	userid int64 `db:"user_id"`
	userName  string `db:"user_name"`
	email  string
}

func articleDBInit()  {
	fmt.Println("Start Init Article DB")
	//db = sqlx.NewDb(sql.Open("sqlite3", ":learn:"), "sqlite3")
	////db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	////if err != nil {
	////	log.Fatalln(err)
	////}
	//db.MustExec(schema)
	//
	//tx := db.MustBegin()
	//tx.MustExec("INSERT INTO user (user_id, user_name, email) VALUES ($1, $2, $3)", "111", "Moiron", "jmoiron@jmoiron.net")
	//tx.Commit()
	//
	//people := []User{}
	//db.Select(&people, "SELECT * FROM person ORDER BY user_id ASC")
	//jason := people[0]

	//fmt.Printf("%#v", jason)

}