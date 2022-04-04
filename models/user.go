package models

import (
	"deyuPersonal/utils"
)

// models 表示这里面放的结构体都是和数据库的表一一对应的,主义  要想将结构体转换为json，需要添加json tag
type User struct {
	Id       int    `json:"id"` // 用户的唯一标识
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 表示添加用户 也就是用户注册或者管理员添加用户
func (user *User) Add() error {
	// TODO 此处将user的数据添加到数据库
	// 例如  insert(sql,user.name,user.phone)
	sqlStr := "insert into user(name,phone,email,password) values (?,?,?,?) "
	_, err := utils.DB.Exec(sqlStr, user.Name, user.Phone, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// 表示修改用户信息 也就是用户修改自己的信息或者管理允修改用户的信息
func (user *User) Update() error {
	// TODO 此处将user的数据添加到数据库
	// 例如  update(sql,user.name,user.phone)
	sqlStr := "update user set name=?,phone=?,email=? where id=?" // 明显有问题，用id来进行查询用户  这里最好不要修改用户的密码，只修改基础资料
	_, err := utils.DB.Exec(sqlStr, user.Name, user.Phone, user.Email, user.Id)
	if err != nil {
		return err
	}
	return nil
}
func Login(username, pwd string) (*User, error) {
	// TODO 此处将根据用户名密码去查数据库,登录查询用户的信息，不要查出密码
	// sqlStr := "select * from user where pwd = ? and phone = ? or email = ?"
	sqlStr := "select id,name,phone,email from user where password=? and (phone =? or email =?)"
	var u User
	err := utils.DB.QueryRow(sqlStr, pwd, username, username).Scan(&u.Id, &u.Name, &u.Phone, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func GetUserList() []*User {
	// TODO 此处将查询所有用户 如果项目没有后台管理，可先不实现该方法
	return nil
}
func CheckUserExist(username string) bool {
	// TODO 此处将查询用户是否存在
	sqlStr := "select id from user where (phone = ? or email = ?)"
	row := utils.DB.QueryRow(sqlStr, username, username)
	var id int
	err := row.Scan(&id)
	if err != nil {
		//fmt.Println()
		return false
	}
	if id > 0 {
		return true
	}
	return false
}
func UpdatePassword(username string, password string) error {
	// TODO 此处用于修改用户密码（设置新密码）
	sqlStr := "update user set password where (phone=? or email = ?)"
	_, err := utils.DB.Exec(sqlStr, password, username, username)
	if err != nil {
		return err
	}
	return nil
}
func GainId(username string) int {
	// TODO 此处用于获取用户的ID
	sqlStr := "select id from user where (phone=? or email = ?)"
	row := utils.DB.QueryRow(sqlStr, username, username)
	var id int
	err := row.Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}
func FindPassword(phone string, email string, password string) bool {
	// TODO 此处用于找回密码
	sqlStr := "update user set password=? where phone =? and email =? "
	row := utils.DB.QueryRow(sqlStr, password, phone, email)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	return true
}
func CheckPhoneEmail11(phone string, email string) bool {
	// TODO 此处用于检验用户输入的电话邮箱是否正确
	sqlStr := "select id from user where (phone=? or email=?) and (phone=? or email=?)"
	row := utils.DB.QueryRow(sqlStr, phone, email)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	return true
}
