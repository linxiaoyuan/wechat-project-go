package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
)
//('1', 'Bill', 'm', '1991-12-11', '1991-12-11', '1387782811113', 'wechat-abc', 'reverse')
type User struct {
	Id int64 `json:"id"`
	Name  string `json:"name"`
	Sex string `json:"sex"`
	Birth string `json:"birth"`
	Register string `json:"register"`
	Tel string `json:"tel"`
	Wechat string `json:"wechat"`
	Meta string `json:"meta"`
}
func userDBInit(db *sql.DB)  {
	
}

func userDBSelectAll(db *sql.DB) {
	print("select all\n")
	rows,err := db.Query("SELECT * FROM  user")
	if err != nil {
		//panic(err)
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Sex, &user.Birth, &user.Register, &user.Tel, &user.Wechat, &user.Meta)
		defer rows.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println(user)
	}
}

func userDBInsertFromJsonString(db *sql.DB, jsonString string) {
	var u User
	err := json.Unmarshal([]byte(jsonString), &u)
	if err != nil {
		print("param error")
		return
	}
	fmt.Println(u)
	stmt, es := db.Prepare("insert into user values(?,?,?,?,?,?,?,?)")
	if es != nil {
		panic(es.Error())
	}
	_, er := stmt.Exec(u.Id,u.Name,u.Sex,u.Birth,u.Register,u.Tel,u.Wechat,u.Meta)
	if er != nil {
		panic(er.Error())
	}

	defer stmt.Close()
}

func userDBInsertTestData2(db *sql.DB){
	jsonStr := `{"id":2, "name":"Tom", "sex":"f", "birth":"1991-12-11", "register":"1992-12-14", "tel":"123344113322", "wechat":"sfsaf", "meta":"reverse"}`
	userDBInsertFromJsonString(db, jsonStr)
}


func userDBInsertTestData(db *sql.DB){
	_,err := db.Exec("INSERT INTO user VALUES ('1', 'Bill', 'm', '1991-12-11', '1991-12-11', '1387782811113', 'wechat-abc', 'reverse')")
	if err != nil {
		//panic(err)
	}
}

