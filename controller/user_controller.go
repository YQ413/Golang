package controller

import (
	"deyuPersonal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// 该函数是所有用户有关操作的入口函数  根据请求的url后面的参数action的类来调用不同的controlle来实现1不同的功能
func BaseUserController(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Type", "application/json")
	// 此处调用别的Controller，
	if request.URL.Query().Get("action") == "login" {
		fmt.Println("-------------login--------------------")
		LoginController(response, request)
	} else if request.URL.Query().Get("action") == "register" {
		fmt.Println("--------------register-------------------")
		phone := request.FormValue("phone")
		name := request.FormValue("name")
		email := request.FormValue("email")
		password := request.FormValue("password")
		// TODO 此处调用checkuser 和 add  用于注册用户 ，我暂时把项目的架构给搭建好了  你自己写完一个方法去测试一下，注意 每次前端用求的url的都要带上action=xxx参数
		if models.CheckUserExist(phone) || models.CheckUserExist(email) { // 这里是用户存在返回true，注意
			fmt.Fprintf(response, "手机号码：%s,或邮箱：%s 已被注册", phone, email)
		} else {
			user := models.User{Name: name, Phone: phone, Email: email, Password: password}
			err := user.Add()
			if err != nil {
				response.Write([]byte(err.Error()))
			} else {
				response.Write([]byte("注册成功！"))
			}
		}
	} else if request.URL.Query().Get("action") == "update" {
		fmt.Println("--------------register-------------------")
		/*phone := request.FormValue("phone")
		name := request.FormValue("name")
		email := request.FormValue("email")
		password := request.FormValue("password")*/
	}
	//else if 等等
	// 比如用户登录   http:// localhost:8080/deyu/user?action=register  请求这个url会走注册这个controller  以此类推
}
func LoginController(response http.ResponseWriter, request *http.Request) {
	uname := request.FormValue("username")
	pwd := request.FormValue("password")
	user, err := models.Login(uname, pwd)
	if err != nil {
		response.Write([]byte(err.Error()))
		return
	}
	if user == nil {
		response.Write([]byte("用户名或密码错误"))
	}
	marshal, _ := json.Marshal(user)
	response.Write(marshal)
}

/*func RegisterController(response http.ResponseWriter,request *http.Request){
	phone := request.FormValue("phone")
	name := request.FormValue("name")
	email := request.FormValue("email")
	password := request.FormValue("password")
	user := models.User{}
	return
}*/
