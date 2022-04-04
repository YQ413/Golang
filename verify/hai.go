package verify

import (
	"bytes"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

var verifys map[string]string

func init() {
	verifys = make(map[string]string, 10)
}
func GetVerify(k string) []byte {
	weight := 525
	height := 221
	dc := gg.NewContext(weight, height) //设置图片
	dc.SetRGB255(63, 64, 87)            //设置图片颜色
	dc.Clear()
	face := loadfontface(120)
	dc.SetFontFace(face)
	dc.SetRGB255(238, 241, 247)
	rand.Seed(time.Now().UnixNano()) //设置随机数种子，加上这行代码，可以保证每次随机都是随机的
	verify := getVerify(4)
	verifys[k] = verify
	for i := 0; i < len(verify); i++ {
		x := rand.Float64()
		dc.DrawStringWrapped(verify[i:i+1], float64(115+i*100), 90, 0.5, x, 300, 1.5+x, gg.AlignCenter) //写字
	}
	buffer := bytes.NewBuffer(nil)
	dc.EncodePNG(buffer)
	return buffer.Bytes()
}
func loadfontface(size float64) font.Face {
	fontBytes, err := ioutil.ReadFile("C:\\Users\\袁琼\\go\\src\\deyuPersonal\\verify\\FZZYJW.TTF") //读取文件字体
	if err != nil {
		log.Println(err) //输出会有日期和时间，其余与fmt.println效果相同
		return nil
	}
	font, err := truetype.Parse(fontBytes) //识别ttf字体？
	face := truetype.NewFace(font, &truetype.Options{Size: size, DPI: 72})
	return face
}
func getVerify(n int) string {
	var name = "abcdefghijkmnlopqrstuvwxyzABCDEFGHIJKMNLOPQRSTUVWXYZ1234567890"
	var randData string
	for j := 0; j < n; j++ {
		data := rand.Intn(62)
		randData += name[data : data+1]
	}
	return randData
}
func DeleteVerify(key string) {
	_, exit := verifys[key]
	if exit {
		delete(verifys, key)
	}
}
func GetVeri(key string) string {
	return verifys[key]
}
func SetVerify(key string, value string) {
	verifys[key] = value
}
