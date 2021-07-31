package tests

import (
	"fmt"
	"github.com/jameskeane/bcrypt"
	"testing"
)

func TestBcrypt(t *testing.T) {
	pwd := "123456"
	fmt.Println("pwd: ", pwd)

	hashPwd, _ := bcrypt.Hash(pwd)
	fmt.Println("hashPwd: ", hashPwd)
	if bcrypt.Match(pwd, hashPwd) {
		fmt.Println("They match")
	}
}

func TestBcrypt2(t *testing.T) {
	pwd := "doudou"
	fmt.Println("pwd: ", pwd)

	hashPwd := "$2a$10$BynoLxqTC0d6WZRnSfmRLeBTtsdM0XZ/rWkmPX.vepX5Y5t0KTF.2"
	fmt.Println("hashPwd: ", hashPwd)
	if bcrypt.Match(pwd, hashPwd) {
		fmt.Println("They match")
	}
}
