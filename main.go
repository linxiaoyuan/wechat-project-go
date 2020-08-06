package main

import (
	"fmt"
	"github.com/linxiaoyuan/wechat-project/database"
	"github.com/linxiaoyuan/wechat-project/web"
)

var dbManager *database.DBManager
var webServer *web.Webserver
func main() {
	fmt.Println("======Start=======")
	initDatabase()
	//startWebServer()

	fmt.Println("======End=========")
}

func startWebServer()  {
	webServer = web.New()
	webServer.Run()
}

func initDatabase()  {
	dbManager = database.New()
}