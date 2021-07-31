package util

import (
	"github.com/jameskeane/bcrypt"
)

// PasswordEncrypt 密码加密
// rawPwd 真实密码
func PasswordEncrypt(rawPwd string) string {
	encryptPwd, _ := bcrypt.Hash(rawPwd)
	return encryptPwd
}

// PasswordMatches 密码比对
// rawPwd 真实密码
// encryptPwd 加密后字符
func PasswordMatches(rawPwd, encryptPwd string) bool {
	return bcrypt.Match(rawPwd, encryptPwd)
}
