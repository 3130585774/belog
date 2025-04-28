package controllers

import (
	"belog/models"
	"fmt"
	"time"
)

type AddArticleController struct {
	BaseController
}

// Get /*当访问/add路径的时候回触发AddArticleController的Get方法响应的页面是通过TpName*/
func (c *AddArticleController) Get() {
	c.TplName = "write_atrticle.gohtml"
}

// Post 通过this.ServerJSON()这个方法去返回json字符串
func (c *AddArticleController) Post() {

	//获取浏览器传输的数据，通过表单的name属性获取值
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")
	author := c.GetString("author")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中
	art := models.Article{Title: title, Tags: tags, Short: short, Content: content, Author: author, Createtime: time.Now().Unix()}
	_, err := models.AddArticle(art)

	//返回数据给浏览器
	var response map[string]interface{}
	if err == nil {
		//无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	c.Data["json"] = response
	c.ServeJSON()
}
