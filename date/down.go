package date

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
	"errors"

	"github.com/PuerkitoBio/goquery"
)

var accept = "ext/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
var acceptEncoding = "gzip, deflate, br"
var acceptLanguage = "zh-CN,zh;q=0.9"
var cacheControl = "max-age=0"
var connection = "keep-alive"
var cookie = "_ga=GA1.2.640308922.1517208694; _gid=GA1.2.1969781450.1517208694; Hm_lvt_9734b08ecbd8cf54011e088b00686939=1517208694; LXB_REFER=www.baidu.com; Hm_lvt_50b0b45724f4f39e2a94cb8af0e9b547=1517208709; SMM_auth_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjZWxscGhvbmUiOiIxMzAyNzk0MTg4MCIsImNvbXBhbnlfaWQiOjAsImNvbXBhbnlfc3RhdHVzIjowLCJjcmVhdGVfYXQiOjE1MTcyMDkxMDksImVtYWlsIjoiIiwiZW5fZW5kX3RpbWUiOjAsImVuX3JlZ2lzdGVyX3N0ZXAiOjEsImVuX3JlZ2lzdGVyX3RpbWUiOjAsImVuX3N0YXJ0X3RpbWUiOjAsImVuX3VzZXJfdHlwZSI6MCwiZW5kX3RpbWUiOjAsImlzX21haWwiOjAsImlzX3Bob25lIjoxLCJsYW5ndWFnZSI6ImhxIiwibHlfZW5kX3RpbWUiOjAsImx5X3N0YXJ0X3RpbWUiOjAsImx5X3VzZXJfdHlwZSI6MCwicmVnaXN0ZXJfdGltZSI6MTUxNzIwOTEwOSwicyI6MjAsInN0YXJ0X3RpbWUiOjAsInVzZXJfaWQiOjEyNDM5ODUsInVzZXJfbmFtZSI6IlNNTTE1MTcyMDkxMDlaQyIsInVzZXJfdHlwZSI6MCwienhfZW5kX3RpbWUiOjAsInp4X3N0YXJ0X3RpbWUiOjAsInp4X3VzZXJfdHlwZSI6MH0.bPvDyGFCPxIu4lYIh0RrLjUozu-vtKN_GcePDo20lm4; _gat=1; _gat_UA-102039857-2=1; Hm_lpvt_9734b08ecbd8cf54011e088b00686939=1517210593; Hm_lpvt_50b0b45724f4f39e2a94cb8af0e9b547=1517210593"
var host = "hq.smm.cn"
var upgradeInsecureRequests = "1"

//var userAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"

var url = "https://hq.smm.cn/tong"

//Product 产品信息结构体
type Product struct {
	Category string //种类
	Name     string //名称
	Orange   string //价格范围
	Price    string //均价
	Change   string //涨跌
	Uint     string //单位
	Date     string //时间
	URL      int    //链接
}

//productChan 内部数据通道
//var productChan = make(chan Product, 20)

//ExChan 外部部数据通道
var ExChan = make(chan []Product, 28)

//限制进程通道
//var dateChan = make(chan []string, 5)

var userAgent = [...]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

//GetRandomUserAgent 设置header user-Agent
func GetRandomUserAgent() string {
	return userAgent[r.Intn(len(userAgent))]
}

//Down 获取上海有色金属网的铜价信息
//获取的信息按参数保存在本目录下
//每天保存一个文档
func Down(filename string) error {

	//获取数据文件的时间
	//fi := strings.Split(filepath.Base(filename), ".")[0]
	var fi string
	now := time.Now()
	bo := now.Before(time.Date(now.Year(), now.Month(), now.Day(), 10, 51, 0, 0, now.Location()))
	if bo {
		tb := now.Add(time.Hour * (-24))
		fi = tb.Format("2006-01-02")
	} else {
		fi = now.Format("2006-01-02")
	}
	//链接计数
	uRL := 1

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		setEmail(err.Error())
		return err
	}

	req.Header.Add("Accept", accept)
	//req.Header.Add("Accept-Encoding", acceptEncoding)	//加上会乱码
	req.Header.Add("Accept-Language", acceptLanguage)
	req.Header.Add("Cache-Control", cacheControl)
	req.Header.Add("Connection", connection)
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Host", host)
	req.Header.Add("Upgrade-Insecure-Requests", upgradeInsecureRequests)
	req.Header.Add("User-Agent", GetRandomUserAgent())

	resp, err := client.Do(req)
	if err != nil {
		setEmail(err.Error())
		return err
	}
	defer resp.Body.Close()

	var productSlice []Product
	query, err := goquery.NewDocumentFromReader(resp.Body)
	//s, _ := query.Find("ul.tab-item-ul").Find("li").Find("a").Attr("title")
	sq := query.Find("ul.tab-item-ul").Find("li")
	sq.Each(func(index int, sel *goquery.Selection) {

		name, _ := sel.Find("a").Attr("title")
		if name == "" {
			return 
		}
		categoryName := sel.Parent().Prev().Text()
		orange := sel.Find(".value2").Find("span").Text()
		price := sel.Find(".value3").Text()
		change := sel.Find(".value4").Text()
		uInt := sel.Find(".value5").Text()

		var dateProduct Product
		dateProduct.Category = categoryName
		dateProduct.Name = name
		dateProduct.Orange = orange
		dateProduct.Price = price
		dateProduct.Change = change
		dateProduct.Uint = uInt
		dateProduct.Date = fi
		dateProduct.URL = uRL
		uRL++

		productSlice = append(productSlice, dateProduct)
	})
	fmt.Println(productSlice)
	if productSlice == nil {
		return errors.New("not data")
	}
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		setEmail("Can't OpenFile :" + filename + ":" + time.Now().Format("2006/01/02 15:04:05"))
		return err
	}
	defer file.Close()
	dec := gob.NewEncoder(file)
	err = dec.Encode(productSlice)
	if err != nil {
		file.Close()
		_ = os.Remove(filename)
		setEmail("Can't save :" + filename + ":" + time.Now().Format("2006/01/02 15:04:05"))
		return err
	}
	setEmail("数据下载成功" + filename)
	//数据传送至sql
	ExChan <- productSlice
	return err
}
