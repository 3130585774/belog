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
	articId, _ := c.GetInt("id")
	fmt.Println(articId)
	fmt.Println()
	_, err := models.DeleteArticle(articId)
	if err != nil {
		return
	}
	c.Redirect("/", 302)
}
