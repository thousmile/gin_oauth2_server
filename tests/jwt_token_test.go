package tests

import (
	"fmt"
	"gin_oauth2_server/util"
	"testing"
	"time"
)

func TestJwtToken(t *testing.T) {
	token, _ := util.CreateAccessToken("123456")
	bearerToken := "Bearer " + token
	fmt.Println("BearerToken : ", bearerToken)
	id, err := util.GetIdFromToken(bearerToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("id : ", id)

}

func TestDS(t *testing.T) {
	fmt.Println("开始时间: ", time.Now().Format("2006-01-02 15:04:05"))
	bearerToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoiMTYyNzYyMzgyMCIsImxvZ2luSWQiOiIxMjM0NTYifQ.VrDR32p8o37mgbnXfZw6X3jarewzd44fLOv4vM8m8s0"
	id, err := util.GetIdFromToken(bearerToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("id : ", id)
}

func TestLoginId(t *testing.T) {
	fmt.Println(util.CreateLoginId())
}
