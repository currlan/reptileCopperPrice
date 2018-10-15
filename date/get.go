package date

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

//Get 返回指定日期数据
func Get(filepath string) ([]Product, error) {
	//返回的数据
	var productSlice []Product
	file, err := os.OpenFile("./date/date/"+filepath+".log", os.O_RDONLY, 0766)
	if err != nil {
		setEmail(filepath)
		return nil, err
	}
	defer file.Close()
	dec := gob.NewDecoder(file)
	err = dec.Decode(&productSlice)
	if err != nil {
		setEmail("Can't get the" + filepath + "date:" + time.Now().Format("2006/01/02  15:04:05"))
		return nil, err
	}
	return productSlice, err
}

//setEmail 把记录写入记录文档
func setEmail(filepath string) {
	f, err := os.OpenFile("./date/record.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0766)
	if err != nil {
		fmt.Println("Can't create or find the record.log ")
		return
	}
	defer f.Close()
	fmt.Println(filepath)
	f.WriteString(filepath + "\r\n")
}
