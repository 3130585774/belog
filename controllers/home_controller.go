package controllers

import (
	"belog/models"
	"fmt"
)

type HomeController struct {
	//beego.Controller
	BaseController
}

func (c *HomeController) Get() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	var artList []models.Article
	artList, _ = models.FindArticleWithPage(page)
	c.Data["PageCode"] = 1
	c.Data["HasFooter"] = true
	fmt.Println("IsLogin:", c.IsLogin, c.Loginuser)
	c.Data["Content"] = models.MakeHomeBlocks(artList, c.IsLogin)
	c.TplName = "home.gohtml"
}
