package main

import (
	_ "belog/routers"
	"belog/utils"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	utils.InitMysql()
	beego.Run()

}
