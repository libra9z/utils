package utils

import (
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	fmt.Println("start generate")
	iw, _ := NewIdWorker(2)
	var prevId int64 = 0
	for i := 0; i < 10; i++ {
		id, err := iw.NextId()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(id)
		}
		if prevId >= id {
			panic("prevId >= id")
		} else {
			prevId = id
		}
	}
	fmt.Println("end generate")
}

func TestSnowFlakeParseId(t *testing.T) {
	iw, _ := NewIdWorker(2)
	for i := 0; i < 200; i++ {
		id, err := iw.NextId()
		if err != nil {
			fmt.Println(err)
		} else {
			t, ts,hid, wid := ParseId(id)
			//输出ID
			fmt.Println(id)
			//输出时间
			fmt.Println(t)
			//输出时间戳
			fmt.Println(ts)
			//输出workid
			fmt.Println(wid)
			//输出hostid
			fmt.Println(hid)

		}
	}
}