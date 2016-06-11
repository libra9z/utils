package utils

import (
	"net/http"
	"net/url"
	"strconv"
	"fmt"
	"crypto/md5"
	"io/ioutil"
	"encoding/json"
	"strings"

)

func GetSysTime(uri string)(st int64){
	
	resp, err := http.Get(uri)
	if err != nil {
		// handle error
		fmt.Println(err)
		return 0
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		// handle error
		fmt.Println(err)
		return 0
	}
	
	var vs map[string]interface{}
	
	err = json.Unmarshal(body,&vs)
	
	if err != nil {
		fmt.Printf("cannot convert json.\n")
		return 0
	}
	
	if vs["Time"] != nil {
		st = int64(vs["Time"].(float64))
	}	
	
	return st
}

func GetToken(uri string,appid int,appkey string,ti int64) (token string ){
	
	v := url.Values{}
	
	t := strconv.FormatInt(ti,10)
	
	v.Add("time",t)
	
	b := appkey+t
	m5 := md5.Sum([]byte(b))

	reqtoken := fmt.Sprintf("%02x", m5)
	
	fmt.Println("reqtoken="+reqtoken)
	
	v.Add("requesttoken",reqtoken)
	
	resp, err := http.Post(uri,"application/x-www-form-urlencoded",strings.NewReader(v.Encode()))
	
	if err != nil {
		fmt.Println(err)
	}
	
	if resp == nil {
		fmt.Printf("no response body return.\n")
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(string(body))
	}

	var vs map[string]interface{}
	err = json.Unmarshal(body,&vs)
	
	if err != nil {
		fmt.Printf("cannot convert json.\n")
		return ""
	}
	
	if vs["Token"] != nil {
		token = vs["Token"].(string)
	}	
			
	return token
}

func CheckUser(uri string,appid int,token string,value string,typ string)(userid int64,err error){

	v := url.Values{}
	
	v.Add("type",typ)
	v.Add("value",value)
	
	resp, err := http.Post(uri,"application/x-www-form-urlencoded",strings.NewReader(v.Encode()))
	
	if err != nil {
		fmt.Println(err)
	}
	
	if resp == nil {
		fmt.Printf("no response body return.\n")
		return 0,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(string(body))
	}
	
	
	var vs map[string]interface{}
	err = json.Unmarshal(body,&vs)
	
	if err != nil {
		fmt.Printf("cannot convert json.\n")
		return 0,err
	}

	//fmt.Printf("body = %v.\n",vs)
	
	if vs["UserId"] != nil {
		userid = int64(vs["UserId"].(float64))
	}	
	
	return userid, err
}
