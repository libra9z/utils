package utils

import (
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"errors"
)

func ServiceGet(uri string)(interface{},error) {

	resp, err := http.Get(uri)

	if resp == nil || err != nil  {
		fmt.Printf("no response body return or has error(%v).\n",err)
		return 0,errors.New("no response body return or has error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(string(body))
		return 0,err
	}


	var vs interface{}
	err = json.Unmarshal(body,&vs)

	if err != nil {
		fmt.Printf("cannot convert json.\n")
		return 0,err
	}

	return vs,err
}

func ServicePost(uri string,params string)(interface{},error) {

	resp, err := http.Post(uri,"application/json",strings.NewReader(params))

	if resp == nil || err != nil  {
		fmt.Printf("no response body return or has error(%v).\n",err)
		return nil,errors.New("no response body return or has error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(string(body))
		return nil,err
	}else {
		fmt.Println(string(body))
	}


	var vs map[string]interface{}
	err = json.Unmarshal(body,&vs)

	if err != nil {
		fmt.Printf("cannot convert json.\n")
		return nil,err
	}

	return vs,err
}

func ServicePut(uri string,params string)(interface{},error) {

	client := &http.Client{}
	req,err := http.NewRequest(http.MethodPut,uri,strings.NewReader(params))
	resp, err := client.Do(req)

	if resp == nil || err != nil  {
		fmt.Printf("no response body return or has error(%v).\n",err)
		return 0,errors.New("no response body return or has error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(string(body))
		return 0,err
	}


	var vs map[string]interface{}
	err = json.Unmarshal(body,&vs)

	if err != nil {
		fmt.Printf("cannot convert json.\n")
		return 0,err
	}

	return vs,err
}

func ServiceDelete(uri string,params string)(interface{},error) {

	client := &http.Client{}
	req,err := http.NewRequest(http.MethodDelete,uri,strings.NewReader(params))
	resp, err := client.Do(req)

	if resp == nil || err != nil  {
		fmt.Printf("no response body return or has error(%v).\n",err)
		return 0,errors.New("no response body return or has error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(string(body))
		return 0,err
	}


	var vs map[string]interface{}
	err = json.Unmarshal(body,&vs)

	if err != nil {
		fmt.Printf("cannot convert json.\n")
		return 0,err
	}

	return vs,err
}
