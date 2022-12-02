package controllers

import (
	"belog/models"
	"fmt"
	"strconv"
)

type ShowArticleController struct {
	//beego.Controller
	BaseController
}

func (this *ShowArticleController) Get() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:", id)
	//获取id所对应的文章信息
	art := models.QueryArticlesWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Content"] = art.Content
	//this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
	this.TplName = "show_article.gohtml"
}
