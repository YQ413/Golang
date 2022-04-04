package controller

import (
	"deyuPersonal/verify"
	"github.com/gofrs/uuid"
	"net/http"
)

func GetVerifyCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Headers", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Expose-Headers", "verify")
	w.Header().Set("Content-Type", "image/jpeg")
	key := r.Header.Get("verify")
	if len(key) == 0 {
		u2, _ := uuid.NewV4()
		s := u2.String()
		w.Header().Set("verify", s)
	}
	getVerify := verify.GetVerify(key)
	//fmt.Println(key)
	w.Write(getVerify)
}
