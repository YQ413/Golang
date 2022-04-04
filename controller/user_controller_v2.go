package controller

import (
	"deyuPersonal/models"
	"deyuPersonal/token"
	"deyuPersonal/verify"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type UserController struct {
}

// 每次请求/deyu/v2/user路径都会调用下列方法
func (uc *UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// // 捕获异常
	// defer func() {
	//  if err := recover(); err != nil {
	//   fmt.Println(err)
	//  }
	// }()
	r.ParseForm()
	w.Header().Set("Access-Control-Allow-Origin", "*") //支持哪些来源的请求跨域
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")  //支持哪些跨域方法
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Expose-Headers", "verify")  //跨域请求暴露的字段
	of := reflect.ValueOf(uc)
	fmt.Println(of)
	if of.IsNil() || of.IsZero() {
		w.Write([]byte("BaseController is null pointer exception"))
		return
	}
	action := r.URL.Query().Get("action") //   拿到url后面的action参数
	if len(action) == 0 {
		w.Write([]byte("action is bull"))
		return
	}
	fmt.Println(action)
	method := of.MethodByName(action) // 得到UserController结构体中名为action的方法
	fmt.Println(method)
	if method == (reflect.Value{}) {
		w.Write([]byte("action is failed"))
		return
	}
	method.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}) // Call 表示调用上述得到的方法
	// 忽略以下内容
	// path := method.Call([]reflect.Value{reflect.ValueOf(&w), reflect.ValueOf(r)}) // res 为调用方法的返回值
	// u, _ := url.Parse(r.Host + r.URL.String() + path[0].String())
	// proxy := httputil.NewSingleHostReverseProxy(u)
	// proxy.ServeHTTP(w, r)
}
func NewUserController() http.Handler {
	return &UserController{}
}
func (uc *UserController) Register(response http.ResponseWriter, request *http.Request) {
	fmt.Println("--------------register-------------------")
	phone := request.FormValue("phone")
	name := request.FormValue("name")
	email := request.FormValue("email")
	password := request.FormValue("password")
	fmt.Println("各数据：",phone,name,email,password)
	// TODO 此处调用checkuser 和 add  用于注册用户 ，我暂时把项目的架构给搭建好了  你自己写完一个方法去测试一下，注意 每次前端用求的url的都要带上action=xxx参数
	if len(phone) > 0 && len(name) > 0 && len(email) > 0 && len(password) > 0 {
		if models.CheckUserExist(phone) || models.CheckUserExist(email) { // 这里是用户存在返回true，注意
			fmt.Fprintf(response, "手机号码：%s,或邮箱：%s 已被注册", phone, email)
			fmt.Println(phone,email)
		} else {
			user := models.User{Name: name, Phone: phone, Email: email, Password: password}
			fmt.Println(user.Phone)
			err := user.Add()
			if err != nil {
				response.Write([]byte(err.Error()))
			} else {
				fmt.Println("注册成功")
				response.Write([]byte("注册成功！"))
			}
		}
	} else {
		fmt.Fprintln(response, "请完善所有信息！")
	}
}
func (uc *UserController) Login(response http.ResponseWriter, request *http.Request) {
	fmt.Println("-----------------------------login  v2 ------------------------------")
	key := request.Header.Get("verify")
	value := request.FormValue("verify")
	uname := request.FormValue("username")
	pwd := request.FormValue("password")
	fmt.Println("uname=",uname)
	fmt.Println("password=",pwd)
	fmt.Println("key=", key)
	fmt.Println("value=", value)
	if value != verify.GetVeri(key) {
		response.Write([]byte("验证码错误"))
		return
	} else {
		go verify.DeleteVerify(key)
	}
	user, err := models.Login(uname, pwd)
	if err != nil {
		response.Write([]byte(err.Error()))
		return
	}
	if user == nil {
		response.Write([]byte("用户名或密码错误"))
	}
	getToken, s := token.GetToken(user.Phone, 60*60*6)
	response.Header().Set("token", getToken)
	response.Header().Set("refreshToken", s)
	//fmt.Println(token)
	ud := &UserData{
		User:         user,
		Token:        getToken,
		RefreshToken: s,
	}
	marshal, _ := json.Marshal(ud)
	response.Write(marshal)
}

type UserData struct {
	User         *models.User `json:"user"`
	Token        string       `json:"token"`
	RefreshToken string       `json:"refresh_token"`
}

// 以后每增加user的操作 写一个方法即可，注意方法没有返回值，且方法的阐述为response http.ResponseWriter, request *http.Request，例如
//func (uc *UserController) 操作名(response http.ResponseWriter, request *http.Request)  {
// 要返回数据给前端，直接调用write方法或fmt.Fprintln()
// 如果返回的是json格式的数据，要先用json.Marshal()函数转换
//response.Write()
//}
func (uc *UserController) FindPassword(response http.ResponseWriter, request *http.Request) {
	fmt.Println("---------------FindPassword---------------")
	phone := request.FormValue("phone")
	email := request.FormValue("email")
	password := request.FormValue("password")
	if models.FindPassword(phone, email, password) {
		response.Write([]byte("修改成功！"))
	} else {
		fmt.Fprintln(response, "输入的信息有误！")
	}
	/*uname := request.FormValue("username1")
	username :=request.FormValue("username2")
	password := request.FormValue("password")
	if models.CheckUserExist(uname){
		if models.CheckPhoneEmail(username,uname){
			err := models.FindPassword(uname,password)
			if err != nil{
				fmt.Fprintln(response,err)
			}
			response.Write([]byte("修改成功！"))
		}
	}*/
	/*preToken := request.Header.Get("token")
	preRefreshToken := request.Header.Get("refreshToken")
	err1 := token.CheckToken(preToken)
	if err1 != nil{
		err2 := token.CheckToken(preRefreshToken)
		if err2 != nil{
			fmt.Fprintln(response,"请登录！")
		}
	}*/
	//pawString := strconv.FormatInt(pswint64, 30)  //int64转化为字符串
	//fmt.Println(pawString)
	/*MapId := session.Session{}.GetFlag()
	fmt.Println(MapId)
	for key = range MapId {
		if user.Id == key {
			fmt.Fprintln(response,models.FindPassword(uname))
		} else{
			fmt.Fprintln(response,"请先登录")
		}
	}*/
}
func (uc *UserController) UpdatePassword(response http.ResponseWriter, request *http.Request) {
	//uname := request.FormValue("username")
	password := request.FormValue("password")
	preToken := request.Header.Get("token")
	preRefreshToken := request.Header.Get("refreshToken")
	err := token.CheckToken(preToken)
	if err != nil {
		err1 := token.CheckToken(preRefreshToken)
		if err1 != nil {
			fmt.Fprintln(response, "请登录！")
		}
	}
	getUsername, err2 := token.GetUsername(preToken)
	if err2 != nil {
		fmt.Fprintln(response, err2)
	}
	err3 := models.UpdatePassword(getUsername, password)
	if err3 != nil {
		fmt.Fprintln(response, err3)
	}
	response.Write([]byte("修改成功！"))
}
func (uc *UserController) UpdateUser(response http.ResponseWriter, request *http.Request) {
	fmt.Println("---------------UpdateUser---------------")
	var newUser models.User
	newUser.Phone = request.FormValue("phone")
	newUser.Name = request.FormValue("name")
	newUser.Email = request.FormValue("email")
	newUser.Id = models.GainId(newUser.Phone)
	preToken := request.Header.Get("token")
	preRefreshToken := request.Header.Get("refreshToken")
	err := token.CheckToken(preToken)
	if err != nil {
		err1 := token.CheckToken(preRefreshToken)
		if err1 != nil {
			fmt.Fprintln(response, "请登录！")
		}
	}
	err2 := newUser.Update()
	if err2 != nil {
		fmt.Fprintln(response, err)
	}
	response.Write([]byte("修改成功！"))
	/*MapId := session.Session{}.GetFlag()
	fmt.Println(MapId)
	for key = range MapId {
		if updateUser.Id == key {
			if updateUser.Update() != nil{
				fmt.Fprintln(response,updateUser.Update())
			}else{
				fmt.Fprintln(response,"修改成功！")
			}
		} else{
			fmt.Fprintln(response,"请先登录")
		}
	}*/
}

//randData = append(randData,name[data:data+1])
/*name[data:data+1]
randData :=append(randData,data)*/
