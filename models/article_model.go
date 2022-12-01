package models

import (
	"belog/utils"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
	//Status int //Status=0为正常，1为删除，2为冻结
}

//---------数据处理-----------
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}

//-----------数据库操作---------------
//插入一篇文章
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}
func DeleteArticle(id int) (int64, error) {
	return utils.ModifyDB("DELETE FROM article WHERE (id = ?)", id)
}
