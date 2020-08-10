package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type MsgData struct {
	X, Y, Z int
	Name []string
	_            [2]byte
}
var network bytes.Buffer //网络传递的数据载体
func main() {
	err := senMsg()
	if err!=nil {
		fmt.Println("编码错误")
		return
	}
	err = revMsg()
	if err!=nil {
		fmt.Println("解码错误")
		return
	}
}

func senMsg()error {
	fmt.Print("开始执行编码（发送端）")
	//var fullPath = path.Join("~", "data.dat")
	f, _ := os.OpenFile("/Users/tuo/data.dat", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	defer f.Close()
	enc := gob.NewEncoder(f)
	var names = []string {"fff", "ddd", "gggg"}
	sendMsg:=MsgData{X:3, Y:4, Z:5, Name:names}
	fmt.Println("原始数据：",sendMsg)
	err := enc.Encode(&sendMsg)
	return  err
}
func revMsg()error {
	f, _ := os.Open("/Users/tuo/data.dat")

	defer f.Close()
	var revData MsgData
	dec:=gob.NewDecoder(f)
	err:= dec.Decode(&revData) //传递参数必须为 地址
	fmt.Println("解码之后的数据为：",revData)
	return err
}
