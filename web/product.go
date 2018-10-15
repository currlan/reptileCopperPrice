package web

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"../date"
	"../mysql"
)

type copperStruct struct {
	Name, App   string
	Price, Date string
	ToNum       []date.Product
}

func copperPost(w http.ResponseWriter, r *http.Request) {

}

func copperGet(w http.ResponseWriter, r *http.Request) {

	var coppS copperStruct
	coppS.ToNum = date.ToN

	t, err := template.ParseFiles("./template/table.html", "./template/head.html", "./template/end.html", "./template/title.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(w, coppS)
	if err != nil {
		fmt.Println(err)
	}
}

func copperGet1(w http.ResponseWriter, r *http.Request) {

	getURL := r.URL.Path[(len("/copper/")) : len(r.URL.Path)-1]
	//fmt.Println(getURL)
	copp, _ := mysql.Select(nameMap[getURL], "0")
	coopS := strings.Split(copp, "|")
	var coo copperStruct
	coo.ToNum = date.ToN
	coo.Name = coopS[0]

	numUint, _ := strconv.Atoi(getURL)
	coo.App = "价格： " + priceUnit[numUint-1]
	for i := 0; i < len(coopS)-1; i = i + 5 {
		coo.Price = coo.Price + coopS[i+2] + ","
		coo.Date = coo.Date + "'" + coopS[i+4] + "'" + ","
	}
	t, err := template.ParseFiles("./template/product.html", "./template/head.html", "./template/end.html", "./template/title.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(w, coo)
	if err != nil {
		fmt.Println(err)
	}
}

func copperGet2(w http.ResponseWriter, r *http.Request) {

	coppS := "密码错误"

	t, err := template.ParseFiles("./template/geren.html", "./template/head.html", "./template/end.html", "./template/title.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(w, coppS)
	if err != nil {
		fmt.Println(err)
	}
}

func copperGet3(w http.ResponseWriter, r *http.Request) {

	coppS := "密码错误"

	t, err := template.ParseFiles("./template/shine.html", "./template/head.html", "./template/end.html", "./template/title.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(w, coppS)
	if err != nil {
		fmt.Println(err)
	}
}

var nameMap = map[string]string{
	"1":  "SMM 1#电解铜",
	"2":  "平水铜",
	"3":  "升水铜",
	"4":  "贵溪铜",
	"5":  "湿法铜",
	"6":  "洋山铜溢价(仓单)",
	"7":  "洋山铜溢价(提单)",
	"8":  "人民币洋山铜溢价",
	"9":  "1#电解铜(华东)",
	"10": "1#电解铜(华北)",
	"11": "1#电解铜(华南)",
	"12": "硫酸铜",
	"13": "进口铜精矿TC(周)",
	"14": "进口铜精矿TC(月)",
	"15": "Φ3mm无氧铜丝(硬)",
	"16": "Φ3mm无氧铜丝(软)",
	"17": "Φ8mm无氧铜杆",
	"18": "漆包线",
	"19": "14%磷铜合金",
	"20": "30%砷铜合金",
	"21": "3.0-3.6%铍铜合金",
	"22": "沪 铜粉",
}

//priceUnit 价格单位
var priceUnit = [22]string{"元/吨", "元/吨", "元/吨", "元/吨", "元/吨", "美元/吨", "美元/吨", "元/吨", "元/吨", "元/吨", "元/吨", "元/吨", "美元/吨", "美元/吨", "元/吨", "元/吨", "元/吨", "元/吨", "元/吨", "元/吨", "元/吨", "元/公斤"}
