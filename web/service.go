package web

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"../mysql"
)

type serviceStruct struct {
	User  string
	Spand string
}

// service server
func serviceGet(w http.ResponseWriter, r *http.Request) {

	cook, err := r.Cookie("luoshanshan")
	if err != nil {
		severGeterr(w, r, "", "")
		return
	}
	err = cookieBase64J(cook)
	if err != nil {
		severGeterr(w, r, "", "")
		return
	}
	//跳转
	http.Redirect(w, r, "/", 302)
	fmt.Println(cook)
}

func servicePost(w http.ResponseWriter, r *http.Request) {

	user := r.PostFormValue("user")
	pwdret := r.PostFormValue("password")
	if pwdret == "" {
		severGeterr(w, r, "密码不能为空", user)
		return
	}
	pwd, err := mysql.SelectName(user)
	if err != nil {
		return
	}
	if pwd != r.PostFormValue("password") {
		severGeterr(w, r, "密码错误", user)
		return
	}

	cookieValue := cookieBase64(user, pwd)
	cookie := http.Cookie{
		Name:     user,
		Value:    cookieValue,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/service/", 302)
}

func severGeterr(w http.ResponseWriter, r *http.Request, spand, user string) {

	t, err := template.ParseFiles("./template/service.html", "./template/head.html", "./template/end.html", "./template/title.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	var serS serviceStruct
	serS.Spand = spand
	serS.User = user
	err = t.Execute(w, serS)
	if err != nil {
		fmt.Println(err)
	}
}

//cookieBase64 对cookie加密
func cookieBase64(user, pwd string) string {

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString([]byte(pwd))

	return encodeString
}

func cookieBase64J(cook *http.Cookie) error {

	pwd, err := mysql.SelectName(cook.Name)
	if err != nil {
		return err
	}
	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(cook.Value)
	if err != nil {
		return err
	}
	if string(decodeBytes) != pwd {
		return errors.New("cookie 不一致")
	}
	return err
}
