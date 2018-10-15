package mysql

import (
	"fmt"
)

//InsertName 插入用户数据
func InsertName(name, pwd string) error {
	mu.Lock()
	defer mu.Unlock()
	//插入数据
	fmt.Println("insert")
	stmt, err := db.Prepare("insert username set user=?,pwd=?")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(name, pwd)
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

//SelectName 查找数据
//nameM 名称
//num	显示起始位置
func SelectName(nameM string) (string, error) {

	var user string
	var pwd string
	fmt.Println("select start =======>")

	mu.RLock()
	defer mu.RUnlock()
	rows, err := db.Query("select * from username where user='" + nameM + "'")
	if err != nil {
		return "", err
	}
	for rows.Next() {
		err = rows.Scan(&user, &pwd)
		if err != nil {
			return "", err
		}
	}
	return pwd, err
}

//DeleteName 删除数据
func DeleteName(user string) error {
	mu.Lock()
	defer mu.Unlock()
	stmt, err := db.Prepare("delete from username where user=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user)
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
