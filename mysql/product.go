package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"

	//
	_ "github.com/GO-SQL-Driver/MYSQL"
)

var (
	mu sync.RWMutex
	//Mysqlchan 格式：操作|数据
	//Mysqlchan = make(chan string, 1)
	db = &sql.DB{}
)

//MyInit 连接数据库
func init() {
	//打开数据库连接
	var err error
	db, err = sql.Open("mysql", "entre:545300@tcp(127.0.0.1:3306)/entredb?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}
}

//Insert 插入铜数据
func Insert(name, orange, price, change, date string) error {
	fmt.Println(name)
	mu.Lock()
	defer mu.Unlock()
	//插入数据
	fmt.Println("insert")
	stmt, err := db.Prepare("insert product set name=?,orange=?,price=?,change1=?,date=?,id=?")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(name, orange, price, change, date, 0)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(affect, id)
	return err
}

//Delete 删除数据
func Delete(date string) error {
	mu.Lock()
	defer mu.Unlock()
	stmt, err := db.Prepare("delete from product where date=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(date)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(affect, id)
	return err
}

//Select 查找数据
//nameM 名称
//num	显示起始位置
func Select(nameM, num string) (string, error) {

	fmt.Println("select start =======>")
	var retDate string

	numInt, err := strconv.Atoi(num)
	if err != nil {
		return "", err
	}
	numInt += 28

	mu.RLock()
	defer mu.RUnlock()
	rows, err := db.Query("select * from (" + "select * from product where name=" + "'" + nameM + "'" + " order by id desc limit " + num + "," + strconv.Itoa(numInt) + ") as new order by id")
	if err != nil {
		return "", err
	}
	for rows.Next() {
		var orange string
		var price string
		var change string
		var date string
		var name string
		var id string
		err = rows.Scan(&name, &orange, &price, &change, &date, &id)
		if err != nil {
			return "", err
		}
		fmt.Println(name, orange, price, change, date, "===>select")
		retDate = retDate + name + "|" + orange + "|" + price + "|" + change + "|" + date + "|"
		//fmt.Println(retDate)
	}
	fmt.Println("select over =======>")
	return retDate, err
}

//DropProduct 删除数据包再重新创建
func DropProduct() error {
	_, err := db.Query("DROP TABLE product")
	if err != nil {
		return err
	}
	_, err = db.Query("create table if not exists product(name    char(50) not null,orange  char(40) not null, price   char(40) not null,  change1  char(40) not null,date    char(40)   not null,id int not null primary key auto_increment)DEFAULT CHARSET=utf8;")
	if err != nil {
		return err
	}
	return err
}
