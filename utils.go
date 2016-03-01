package utils

import (
	"strings"
	"github.com/disintegration/imaging"
	"path/filepath"
	"crypto/md5"
	"fmt"
)

func ParseTagOption(str string) map[string]string {
	tags := strings.Split(str, ";")
	setting := map[string]string{}
	for _, value := range tags {
		v := strings.Split(value, ":")
		k := strings.TrimSpace(strings.ToUpper(v[0]))
		if len(v) == 2 {
			setting[k] = v[1]
		} else {
			setting[k] = k
		}
	}
	return setting
}



func GetImageFormat(url string) (*imaging.Format, error) {
	formats := map[string]imaging.Format{
		".jpg":  imaging.JPEG,
		".jpeg": imaging.JPEG,
		".png":  imaging.PNG,
		".tif":  imaging.TIFF,
		".tiff": imaging.TIFF,
		".bmp":  imaging.BMP,
		".gif":  imaging.GIF,
	}

	ext := strings.ToLower(filepath.Ext(url))
	if f, ok := formats[ext]; ok {
		return &f, nil
	} else {
		return nil, imaging.ErrUnsupportedFormat
	}
}

func CheckMD5Passwd( md5pwd string,orgpwd string )(isequal bool,pwd string ) {

	m5 := md5.Sum([]byte(orgpwd))
	
	str := fmt.Sprintf("%02x", m5)
	
	fmt.Printf("[util](CheckMD5Passwd):str=%s\n",str)
	
	if str == strings.ToLower(md5pwd) {
		isequal = true
		pwd = orgpwd
	}else{
		isequal = false
		pwd = ""
	}
	
	return isequal,pwd
}

func CheckMD5Passwd2( md5pwd string,orgpwdmd5 string )(isequal bool,pwd string ) {

	//m5 := md5.Sum([]byte(orgpwd))
	
	//str := fmt.Sprintf("%02x", m5)
	
	ss := strings.Split(orgpwdmd5,"|")
	if len(ss)<2 {
		return false,""
	} 
	
	str := strings.ToLower(ss[1])
	orgpwd := ss[0]
	
	fmt.Printf("[util](CheckMD5Passwd):str=%s\n",str)
	
	if str == strings.ToLower(md5pwd) {
		isequal = true
		pwd = orgpwd
	}else{
		isequal = false
		pwd = ""
	}
	
	return isequal,pwd
}