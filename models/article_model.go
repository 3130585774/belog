package models

import (
	"belog/utils"
	"fmt"
	"github.com/astaxie/beego"
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

// AddArticle ---------数据处理-----------
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}
func EditArticle(article Article) (int64, error) {
	i, err := UpdateArticle(article)
	return i, err
}

//-----------数据库操作---------------
//插入一篇文章
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}
func UpdateArticle(article Article) (int64, error) {
	return utils.ModifyDB("UPDATE article SET title=?, tags=?,short=?,content=?,author=? WHERE id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Id)
}

// DeleteArticle 删除文章
func DeleteArticle(id int) (int64, error) {
	return utils.ModifyDB("DELETE FROM article WHERE (id = ?)", id)
}

//-----------查询文章---------

// FindArticleWithPage 根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	//从配置文件中获取每页的文章数量
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("---------->page", page)
	return QueryArticleWithPage(page, num)
}
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		err := rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		if err != nil {
			return nil, err
		}
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}

// QueryArticlesWithId TODO 根据ID查文章
func QueryArticlesWithId(id int) Article {
	var article []Article
	sql := fmt.Sprintf("where id = %d", id)
	article, _ = QueryArticlesWithCon(sql)
	return article[0]
}
