package models

import (
	"belog/utils"
	"fmt"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}

//--------------数据库操作-----------------

// InsertUser 插入
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

// QueryUserWightCon 按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("SELECT id FROM users %s ", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	err := row.Scan(&id)
	if err != nil {
		return 0
	}
	return id
}

// QueryUserWithUsername QueryUserWithUsername 根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

// QueryUserWithParam QueryUserWithParam 根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}
