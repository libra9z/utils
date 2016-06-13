package utils

import (
    "testing"
    "fmt"
)

func TestCheckUser(t *testing.T) {
	appkey := "6a79527cca7b43b588330db3e4375826"
	appid := 7
	uri1 := "http://dev.laoyou99.cn:80/api/v1/system/time"
	uri2 := "http://dev.laoyou99.cn:80/api/v1/token/generate?appid=7"
	uri3 := "http://dev.laoyou99.cn:80/api/v1/user/check?appid=7&siteid=1&token="
	
	ti := GetSysTime(uri1)
	
	
	fmt.Printf("time=%d\n",ti)
	
	token := GetToken(uri2,appid,appkey,ti)
	
	fmt.Printf("token=%s\n",token)
	
	uri3 = uri3+token
	
	fmt.Println(uri3)
	userid,_ := CheckUser(uri3,appid,token,"15811397368","phone")	
	
	fmt.Printf("userid=%d\n",userid)
}


func TestCheckToken(t *testing.T) {
	appkey := "6a79527cca7b43b588330db3e4375826"
	appid := 7
	uri1 := "http://dev.laoyou99.cn:80/api/v1/system/time"
	uri2 := "http://dev.laoyou99.cn:80/api/v1/token/generate?appid=7"
	uri3 := "http://dev.laoyou99.cn:80/api/v1/token/check?appid=7&siteid=1&token="
	
	ti := GetSysTime(uri1)
	
	
	fmt.Printf("time=%d\n",ti)
	
	token := GetToken(uri2,appid,appkey,ti)
	
	fmt.Printf("token=%s\n",token)
	
	uri3 = uri3+token
	
	fmt.Println(uri3)
	bret := CheckToken(uri3,appid,token)	
	
	fmt.Printf("check token=%v\n",bret)
}

