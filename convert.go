package utils

import (
	"reflect"
	"strconv"
	"fmt"
	"encoding/json"
)

func Convert2Int64(v interface{}) int64 {

	t := reflect.TypeOf(v)
	var ret int64
	switch t.Name() {
	case "string":
		ret,_ = strconv.ParseInt(v.(string),10,64)
	case "float64":
		ret = int64(v.(float64))
	case "int64":
		ret = v.(int64)
	case "int":
		ret = int64(v.(int))
	}

	return ret
}


func Convert2Int(v interface{}) int {

	if v == nil {
		return 0
	}

	t := reflect.TypeOf(v)
	var ret int64
	switch t.Name() {
	case "string":
		ret,_ = strconv.ParseInt(v.(string),10,64)
	case "float64":
		ret = int64(v.(float64))
	case "int64":
		ret = v.(int64)
	case "int":
		ret = int64(v.(int))
	}

	return int(ret)
}

func Convert2Float64(v interface{}) float64 {

	t := reflect.TypeOf(v)
	var ret float64
	switch t.Name() {
	case "string":
		ret,_ = strconv.ParseFloat(v.(string),64)
	case "float64":
		ret = v.(float64)
	case "int64":
		ret = float64(v.(int64))
	case "int":
		ret = float64(v.(int))
	}

	return ret
}

func Convert2Float32(v interface{}) float32 {

	t := reflect.TypeOf(v)
	var ret float32
	var ret1 float64
	switch t.Name() {
	case "string":
		ret1,_ = strconv.ParseFloat(v.(string),64)
		ret = float32(ret1)
	case "float64":
		ret1 = v.(float64)
		ret = float32(ret1)
	case "int64":
		ret = float32(v.(int64))
	case "int":
		ret = float32(v.(int))
	}

	return ret
}

func ConvertToString(v interface{}) string {

	var ret string
	t := reflect.TypeOf(v)
	switch t.Name() {
	case "string":
		ret = v.(string)
	default:
		bb,err:=json.Marshal(v)
		if err != nil {
			fmt.Errorf("json 格式不正确 ：%v",err)
			return ""
		}
		ret=string(bb)
	}

	return ret
}