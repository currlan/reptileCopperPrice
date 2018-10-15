package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"./date"
	"./mysql"
)

/*
//date.Product 产品信息结构体 已经在date包定义
type Product struct {
	Category string //种类
	Name     string //名称
	Orange   string //价格范围
	Price    string //均价
	Change   string //涨跌
	Uint     string //单位
	Date     string //时间
}
*/
//receive 接收下载的数据然后存进Mysql
func receive() {
	go func() {
		for prodate := range date.ExChan {
			func() {
				for i := 0; i < len(prodate); i++ {
					mysql.Insert(prodate[i].Name, prodate[i].Orange, prodate[i].Price, prodate[i].Change, prodate[i].Date)
				}
			}()
		}
	}()
}

//单独把某个文件内容放进sql上
func one(name string) {
	prodate, _ := date.Get(name)
	fmt.Println(prodate)
	func() {
		fmt.Println(prodate)
		for i := 0; i < len(prodate); i++ {
			mysql.Insert(prodate[i].Name, prodate[i].Orange, prodate[i].Price, prodate[i].Change, prodate[i].Date)
		}
	}()
}

//listAll 把结构化的文件内容存近sql上
func listAll() {

	err := mysql.DropProduct()
	if err != nil {
		fmt.Println("mysql.DropProduct", err)
		return
	}
	path := "date/date"
	files, _ := ioutil.ReadDir(path)
	for _, fi := range files {
		if fi.IsDir() {
		} else {
			str := strings.Split(fi.Name(), ".")[0]
			println(str)
			one(str)
		}
	}
}
