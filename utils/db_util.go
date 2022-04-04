package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

var DB *sql.DB

func init() {
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	db, err := sql.Open("mysql", "root:yuan@tcp(127.0.0.1:3306)/deyuperson") // 设置连接数据库的参数
	//连接数据库
	if err != nil {
		//fmt.Println("连接失败！")
		panic(err)
	}
	DB = db
	err = db.Ping()
	fmt.Println("数据库连接成功！")
	if err != nil {
		fmt.Println("发生错误")
	}
}
func Myexec(sqlS string, e interface{}) []interface{} {
	indext1 := strings.Index(sqlS, "(")
	indext2 := strings.Index(sqlS, ")")
	w := sqlS[indext1+1 : indext2]
	t := strings.Split(w, ",") //以逗号分割字符串生成一个字
	yy := reflect.ValueOf(e)
	yy = reflect.Indirect(yy)
	u := make([]interface{}, 0, len(t))
	for _, cont := range t {
		co := strings.TrimSpace(cont)                              //去除首尾的空格
		u = append(u, yy.FieldByName(toCamelCase(co)).Interface()) //按照驼峰命名格式的姓名查找其值，并添加到u切片里面
	}
	fmt.Println(u)
	return u
}
func toCamelCase(name string) string {
	pre := name[0] - 32
	if pre < 65 {
		pre += 32
	}
	for i, s := range name {
		if s == '_' {
			u := name[i+1] - 32
			if u < 65 {
				u += 32
			}
			temp := fmt.Sprintf("%c", pre) + name[1:] + fmt.Sprintf("%c", u) + name[i+2:]
			return temp
		}
	}
	return fmt.Sprintf("%c", pre) + name[1:]
}
