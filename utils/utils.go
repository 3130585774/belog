package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	//切记：导入驱动包
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() {

	fmt.Println("InitMysql....")
	//driverName := beego.AppConfig.String("driverName")

	//driverName:=beego.AppConfig.String("driverName")
	driverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	//orm.RegisterDriver(driverName, orm.DRMySQL). 2

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/myblog?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreateTableWithUser()
		CreateTableWithArticle()
	}

}
func CreateTableWithUser() {
	s := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`

	_, err := ModifyDB(s)
	if err != nil {
		return
	}
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// QueryRowDB 查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// CreateTableWithArticle 创建文章表
func CreateTableWithArticle() {
	s := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	_, err := ModifyDB(s)
	if err != nil {
		return
	}
}
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

// SwitchMarkdownToHtml /**
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}
