package controllers

import (
	"belog/models"
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
	_, err := models.DeleteArticle(artic_id)
	if err != nil {
		return
	}
	c.Redirect("/", 302)
}
