package controllers

import (
	"belog/models"
	"belog/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.gohtml"
}
func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("username:", username, ",password:", password)
	id := models.QueryUserWithParam(username, utils.MD5(password))
	//id := models.QueryUserWithParam(username, password)
	fmt.Println("id:", id)
	if id > 0 {
		/*
			设置了session后会将数据处理设置到cookie，然后再浏览器进行网络请求的时候回自动带上cookie
			因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置
		*/
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	c.SetSession("loginuser", username)
	c.ServeJSON()
}

//func (this *LoginController) Post() {
//	username := this.GetString("username")
//	password := this.GetString("password")
//	fmt.Println("username:", username, ",password:", password)
//	id := models.QueryUserWithParam(username, utils.MD5(password))
//	//id := models.QueryUserWithParam(username, password)
//	fmt.Println("id:", id)
//	if id > 0 {
//		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
//	} else {
//		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
//	}
//	this.ServeJSON()
//}
