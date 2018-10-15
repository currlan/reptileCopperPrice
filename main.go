/*
数据的抓取 ： 只抓取铜价信息
上色网数据更新时段无法抓取数据
*/
package main

import (
	"flag"
	"fmt"

	"./date"
	"./web"
)

var (
	fpathBefor        = "./date/"
	fpathAfter        = ".log"
	callAdministratir = "Can't down the date, Please call to Administratir"
)

func main() {

	//选项：是否重新创建数据表然后重新把数据存进表，默认不重新创建
	mysqlNo := flag.String("m", "no", "yes：重新载入数据库数据")
	//刷新缓存
	flag.Parse()
	if *mysqlNo == "yes" {
		listAll()
	}
	fmt.Println(*mysqlNo)

	//开进程按时下载数据
	go func() {
		date.Time()
	}()

	//接收下载的数据然后存进Mysql
	go func() {
		receive()
	}()

	//开启网站设置，会堵塞主进程
	web.CreateSocket()
}
