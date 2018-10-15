package web

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"

	"github.com/drone/routes"
	"golang.org/x/net/websocket"

	"../mysql"
)

//Echo 网页套接字接口--接收与发送
func echo(ws *websocket.Conn) {

	var reply string

	fmt.Println("websocket=======>")
	for {
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println(err.Error())
			break
		}
		//fmt.Println(reply)
		//check date
		date, err := mysql.Select("1", "0")
		if err != nil {
			break
		}
		if date == "" {
			date = "false"
		}
		fmt.Println(date)
		if err := websocket.Message.Send(ws, date); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func getuser(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./template/index.html", "./template/head.html", "./template/title.html", "./template/end.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	ra := rand.Intn(7)
	err = t.Execute(w, ra)
	if err != nil {
		fmt.Println(err)
	}
}

func modifyuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are modify user %s", uid)
}
func deleteuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are delete user %s", uid)
}
func adduser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are add user %s", uid)
}

//CreateSocket 创建网页连接和websocket
func CreateSocket() {
	//加载JS CSS等静态库
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	mux := routes.New()
	//  /
	mux.Get("/", getuser)
	mux.Post("/", modifyuser)
	mux.Del("/", deleteuser)
	mux.Put("/", adduser)

	//service
	mux.Get("/service/", serviceGet)
	mux.Post("/service/", servicePost)

	//copper
	mux.Get("/copper/", copperGet)
	mux.Post("/copper/", copperPost)

	//product
	for i := 1; i < 23; i++ {
		mux.Get("/copper/"+strconv.Itoa(i)+"/", copperGet1)
	}

	//geren
	mux.Get("/geren/", copperGet2)
	mux.Get("/ge/", copperGet3)

	http.Handle("/", mux)
	http.Handle("/server_client/", websocket.Handler(echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
