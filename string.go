package utils

import "strings"

/*
	去掉数字字符串中的逗号，比如
	 1,711,624,950 变成 1711624950
*/

func ClearComma(s string) string {
	if s == "" {
		return ""
	}
	sret := ""
	ss := strings.Split(s,",")

	for _,sv := range ss {
		sret += sv
	}

	return sret
}

/*
	去掉数字字符串的%,比如：
	3.65% 变为 3.65
*/

func ClearPercent(s string) string {
	if s == "" {
		return s
	}

	ss := strings.Split(s,"%")

	return ss[0]
}

func GetStringIndex(body []byte, bound []byte ) []int {

	blen := len(bound)

	bodylen := len(body)

	index := -1
	var idxs []int

	var j int
	for i := 0; i < bodylen; i++ {
		j = 0
		if body[i] == bound[0] {
			index = i
			for j = 0; j < blen && i+j < bodylen; j++ {
				if body[i+j] != bound[j] && j < blen {
					index = -1
					break
				}
			}
		}
		if j == blen {
			i = i + j
			idxs = append(idxs, index)
		}
	}

	return idxs
}
