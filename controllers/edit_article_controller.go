package controllers

import (
	"belog/models"
	"fmt"
)

type EditArticleController struct {
	BaseController
}

// Get /*当访问/add路径的时候回触发AddArticleController的Get方法响应的页面是通过TpName*/
func (c *EditArticleController) Get() {
	var art models.Article
	articleId, _ := c.GetInt("id")
	art = models.QueryArticlesWithId(articleId)
	c.Data["title"] = art.Title
	c.Data["short"] = art.Short
	c.Data["content"] = art.Content
	c.Data["author"] = art.Author
	c.Data["tags"] = art.Tags
	c.TplName = "edit_article.gohtml"
}

// Post 通过this.ServerJSON()这个方法去返回json字符串
func (c *EditArticleController) Post() {

	//获取浏览器传输的数据，通过表单的name属性获取值
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")
	author := c.GetString("author")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中
	art := models.Article{Title: title, Tags: tags, Short: short, Content: content, Author: author}
	_, err := models.EditArticle(art)

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
