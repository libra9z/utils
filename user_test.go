package utils

import (
    "testing"
)

func TestCheckUser(t *testing.T) {
	appkey := "6a79527cca7b43b588330db3e4375826"
	appid := 7
	uri1 := "dev.laoyou99.cn:80/api/v1/system/time"
	uri2 := "dev.laoyou99.cn:80/api/v1/token/generate?appid=7"
	uri3 := "dev.laoyou99.cn:80/api/v1/user/check?appid=7&siteid=1&token="
	
	ti := GetSysTime(uri1)
	
	
	fmt.Printf("time=%d\n",ti)
	
	token := GetToken(uri2,appid,appkey,ti)
	
	fmt.Printf("token=%s\n",token)
	
	uri3 = uri3+token
	
	userid,_ := CheckUser(uri3,appid,token,"0275976165","rfid")	
	
	fmt.Printf("userid=%d\n",userid)
}

