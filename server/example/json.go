package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	/*
		序列化只能说输出 public
		Id "id,string"以string格式(反)序列化
		json `json:name`取别名
		omitempty,序列化后有空值时不输出该字段，int型0即为空值
		"-"不参与(反)序列化
	*/
	Id      int    `json:"id,string"`
	Name    string `json:"username"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"-"`
}

func main() {
	u := User{
		Id:      12,
		Name:    "wendell",
		Age:     0,
		Address: "广州"}
	data, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	//data:=[]byte(`{"id":"12","username":"wendell","age":1,"Address":"广州"}`)
	u2 := User{}
	err = json.Unmarshal(data, &u2)
	if err != nil {
		fmt.Println(err)
	}
	//由于Address没有被序列号，所以是空值
	fmt.Printf("%+v \n", u2)
}
