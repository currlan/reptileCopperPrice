package date

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//Time 按时下载数据
func Time() {

	now := time.Now()
	//网页数据格式
	weekNot(now)
	fmt.Println(ToN)
	bo := now.Before(time.Date(now.Year(), now.Month(), now.Day(), 10, 50, 0, 0, now.Location()))
	if bo {
		tb := now.Add(time.Hour * (-24))
		if weekday(tb) {
			//fmt.Println(tb.Format("20160102"))
			fn := tb.Format("20060102")
			stat(fn)
			//fmt.Println(fn)
		}
	}
	now = time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
	t := time.NewTimer(next.Sub(now))
	<-t.C
	if weekday(now) {
		fdn := now.Format("20060102")
		stat(fdn)
		//fmt.Println(fdn)
	}
	for {
		now = time.Now()
		next = now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 12, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		if weekday(next) {
			fr := time.Now().Format("20060102")
			stat(fr)
			//fmt.Println(fr)
		}
	}
}

//ToN 当天总数据
var ToN []Product

func stat(filename string) {
	filename = "./date/date/" + filename + ".log"
	num := 0
LOOP:
	_, err := os.Stat(filename)
	if err == nil {
		setEmail("数据已经下载 :" + time.Now().Format("2006/01/02 15:04:05"))
		fi := strings.Split(filepath.Base(filename), ".")[0]
		if err != nil {
			fmt.Println("can't open file", fi)
			return
		}
		fmt.Println(fi, ToN)
		return
	}
	//重复下载超过十次/终止下载
	if num > 10 {
		setEmail("数据未下载成功 :" + time.Now().Format("2006/01/02 15:04:05"))
		return
	}

	Down(filename)
	//等待一分钟然后再去检测是否已经数据
	next := time.Now().Add(time.Second * 5)
	t := time.NewTimer(next.Sub(time.Now()))
	<-t.C
	goto LOOP
}

//weekday 判断是否是周六日
//周六日铜价无更新
func weekday(ti time.Time) bool {
	we := ti.Weekday().String()
	if we == "Sunday" || we == "Saturday" {
		setEmail("Today is " + we)
		//weekNot(ti)
		return false
	}
	return true

}

//weekNot 无铜价更新 抓取以往铜价数据来做全局变量
func weekNot(ti time.Time) {
	tb := ti.Add(time.Hour * (-24))
	fn := tb.Format("20060102")
	//fmt.Println(fn)
	toNhere, err := Get(fn)
	//fmt.Println(toNhere, err)
	if err != nil {
		weekNot(tb)
		return
	}
	if toNhere == nil {
		weekNot(tb)
	} else {
		ToN = toNhere
	}
}
