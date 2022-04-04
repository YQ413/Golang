package controller

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func BaseVideoController(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	response.Header().Set("Access-Control-Allow-Origin", "*") //支持哪些来源的请求跨域
	response.Header().Set("Access-Control-Allow-Headers", "*")
	response.Header().Set("Access-Control-Allow-Methods", "*")  //支持哪些跨域方法
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Type", "application/json")
	// 此处调用别的Controller，
	if request.URL.Query().Get("action") == "videUpload" {
		fmt.Println("-------------videUpload--------------------")
		videUploadController(response, request)
	} else if request.URL.Query().Get("action") == "videoRead" {
		fmt.Println("--------------videoRead-------------------")
		videoReadController(response, request)
	}
}

func videoReadController(response http.ResponseWriter, request *http.Request) {
	file, err := os.Open("./test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
			} else {
				fmt.Println("Read file error!", err)
			}
			break
		}
		fmt.Println(string(line))
	}
}

func videUploadController(response http.ResponseWriter, request *http.Request) {
	//文件上传只允许POST方法
	if request.Method != http.MethodPost {
		response.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = response.Write([]byte("Method not allowed"))
		return
	}

	//从表单中读取文件
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		_, _ = io.WriteString(response, "Read file error")
		return
	}
	//defer 结束时关闭文件
	defer file.Close()
	log.Println("filename: " + fileHeader.Filename)

	//创建文件
	newFile, err := os.Create("./upload/" + fileHeader.Filename)
	if err != nil {
		_, _ = io.WriteString(response, "Create file error")
		return
	}
	//defer 结束时关闭文件
	defer newFile.Close()

	//将文件写到本地
	_, err = io.Copy(newFile, file)
	if err != nil {
		_, _ = io.WriteString(response, "Write file error")
		return
	}
}

/*type VideoController struct {
}

func (vc *VideoController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	ov := reflect.ValueOf(vc)
	fmt.Println(ov)
	if ov.IsNil() || ov.IsZero() {
		w.Write([]byte("BaseController is null pointer exception"))
		return
	}
	action := r.URL.Query().Get("action") //   拿到url后面的action参数
	if len(action) == 0 {
		w.Write([]byte("action is bull"))
		return
	}
	fmt.Println(action)
	method := ov.MethodByName(action) // 得到VideoController结构体中名为action的方法
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
func NewVideoController() http.Handler {
	return &VideoController{}
}
func (vc *VideoController) videUpload(response http.ResponseWriter, request *http.Request) {
	fmt.Println("文件上传-------------------")
	//文件上传只允许POST方法
	if request.Method != http.MethodPost {
		response.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = response.Write([]byte("Method not allowed"))
		return
	}

	//从表单中读取文件
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		_, _ = io.WriteString(response, "Read file error")
		return
	}
	//defer 结束时关闭文件
	defer file.Close()
	log.Println("filename: " + fileHeader.Filename)

	//创建文件
	newFile, err := os.Create("./upload/" + fileHeader.Filename)
	if err != nil {
		_, _ = io.WriteString(response, "Create file error")
		return
	}
	//defer 结束时关闭文件
	defer newFile.Close()

	//将文件写到本地
	_, err = io.Copy(newFile, file)
	if err != nil {
		_, _ = io.WriteString(response, "Write file error")
		return
	}
	_,_ = io.WriteString(response, "Upload success")
	/*if  request.Method == "post"{     //判断的请求
		request.ParseMultipartForm(128 << 20)
		file,header,err := request.FormFile("file") //接受到一个文本数据格式的文件，二进制
		defer file.Close()
		if err != nil{
			log.Fatal(err.Error())
		}
		os.Mkdir("./upload",os.ModePerm)  //创建目录文件
		cur,err := os.Create("./upload/" + header.Filename)  //创建本地文件
		defer file.Close()
		if err != nil{
			log.Fatal(err.Error())
		}
		io.Copy(cur,file)  //写入本地
	}else{
		t, _ := template.ParseFiles("")
		t.Execute(response,nil)
	}*/
	//mark := make(map[string]string, 10)
//}
/*func (vc *VideoController) videoRead(response http.ResponseWriter, request *http.Request) {
	file, err := os.Open("./test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
		for {
			line, err := r.ReadString('\n')
			line = strings.TrimSpace(line)
			if err != nil {
				if err == io.EOF {
					fmt.Println("File read ok!")
				} else {
					fmt.Println("Read file error!", err)
				}
				break
			}
			fmt.Println(string(line))
		}
}*/