package main

import (
	"deyuPersonal/controller"
	"net/http"
)

func main() {
	http.Handle("/deyu/v2/user", controller.NewUserController())
	// 这里拦截url可以分类  比如用户类的就走deyu/user   如何要做什么事情就在后面加参数 action=什么事情就好了
	http.HandleFunc("/deyu/user", controller.BaseUserController)     // 注意url前面加斜线
	http.HandleFunc("/deyu/getverifycode", controller.GetVerifyCode) // 注意url前面加斜线
	// 这里拦截url可以分类  比如视频操作的就走deyu/video   如何要做什么事情就在后面加参数 action=什么事情就好了
	//http.Handle("/deyu/video", controller.NewVideoController())
	http.HandleFunc("deyu/video",controller.BaseVideoController)  //这个你模仿前面的user来写就好了
	http.ListenAndServe(":8080", nil)
}
