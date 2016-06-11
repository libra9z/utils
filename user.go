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
	//"os"
	"compress/gzip"
)

func GetSysTime(uri string)(st int64){
	
	resp, err := http.Get(uri)
	if err != nil {
		// handle error
		return 0
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		// handle error
		return 0
	}
	
	var vs map[string]interface{}
	
	err = json.Unmarshal(body,&vs)
	
	if err != nil {
		fmt.Printf("cannot convert json.\n")
		return 0
	}
	
	if vs["Time"] != nil {
		st = vs["Time"].(int64)
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
	
	v.Add("requesttoken",reqtoken)
	
	reqest, err := http.NewRequest("POST", uri, strings.NewReader(v.Encode()))

    if err != nil {
	    fmt.Println("Fatal error ", err.Error())
	    return ""
	}

  	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	client := &http.Client{nil, nil, nil,10000}
	response, err := client.Do(reqest)
	if response == nil {
		fmt.Println("no body return")
		return ""
	}
    defer response.Body.Close()

	if err != nil {
	    fmt.Println("Fatal error ", err.Error())
	    return ""
	}

    var body []byte
	
    if response.StatusCode == 200 {

	    switch response.Header.Get("Content-Encoding") {
	    case "gzip":
	        reader, _ := gzip.NewReader(response.Body)
	        reader.Read(body)
	    default:
	        bodyByte, _ := ioutil.ReadAll(response.Body)
	        body = bodyByte
	    }

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
	
	if vs["Userid"] != nil {
		userid = vs["Userid"].(int64)
	}	
	
	return userid, err
}