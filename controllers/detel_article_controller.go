package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type DetelArticlerController struct {
	beego.Controller
}

func (c *DetelArticlerController) Get() {
	artic_id, _ := c.GetInt("id")
	fmt.Println(artic_id)
	fmt.Println()
	c.Redirect("/", 302)
}
