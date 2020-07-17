package main

import (
	"fmt"
	"github.com/linxiaoyuan/wechat-project/webserver"
	"github.com/linxiaoyuan/wechat-project/database"
)

func main() {
	fmt.Println("abc")
	initDatabase()
	//startWebServer()
}

func startWebServer()  {
	web := webserver.New()
	web.Run()
}

func initDatabase()  {
	dbManager := database.New()
	dbManager.Init()

}