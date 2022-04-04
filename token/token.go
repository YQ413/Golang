package token

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const secretKey = "YQ123" // 表示常量
/*
用户登录后根据用户的手机号或者邮箱生成一个token和一个refreshToken，refreshToken 是在token失效后重新获取token，从而不必让用户再一次登录 故refreshToken一般过期时间比较长
*/
func GetToken(key string, expiration int) (string, string) { // expiration 过期时间 单位是秒
	token := "yq111" + key + "1a1" + fmt.Sprintf("%d", time.Now().Unix()) + "1a1" + fmt.Sprintf("%d", expiration) + secretKey
	refreshToken := "yq111" + key + "1a1" + fmt.Sprintf("%d", time.Now().Unix()) + "1a1" + fmt.Sprintf("%d", expiration*7) + secretKey
	return base64.StdEncoding.EncodeToString([]byte(token)), base64.StdEncoding.EncodeToString([]byte(refreshToken))
}
func GetTokenByRefreshToken(refreshToken string) (string, error) {
	err := CheckToken(refreshToken)
	if err != nil {
		return "", err
	}
	decodeString, _ := base64.StdEncoding.DecodeString(refreshToken)
	key := string(decodeString)
	key = key[5 : len(key)-5]
	splits := strings.Split(key, "1a1")
	expiration, err := strconv.Atoi(splits[2])
	if err != nil {
		return "", err
	}
	token := "yq111" + splits[0] + "1a1" + fmt.Sprintf("%d", time.Now().Unix()) + "1a1" + fmt.Sprintf("%d", expiration/7) + secretKey
	return base64.StdEncoding.EncodeToString([]byte(token)), nil
}
func CheckToken(key string) error {
	decodeString, err := base64.StdEncoding.DecodeString(key)
	key = string(decodeString)
	if err != nil {
		return err
	}
	if key[len(key)-5:] != secretKey {
		return errors.New("key 不合法!") // 英语不好的无奈
	}
	key = key[5 : len(key)-5]
	splits := strings.Split(key, "1a1")
	expiration, err := strconv.Atoi(splits[2])
	if expiration <= 0 { // 过期时间小于0为永久生效的token 但是一般绝对不会让过期时间小于零 这是非常愚蠢的做法
		return nil
	}
	preTime, err := strconv.Atoi(splits[1])
	if err != nil {
		return err
	}
	if time.Now().Unix() >= int64(preTime+expiration) {
		return errors.New("the key was expiration !")
	}
	return nil
}
func GetUsername(key string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(key)
	key = string(decodeString)
	if err != nil {
		return "", err
	}
	key = key[5 : len(key)-5]
	splits := strings.Split(key, "1a1")
	return splits[0], nil
}
