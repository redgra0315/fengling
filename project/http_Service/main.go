package main

import (
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(resw http.ResponseWriter, r *http.Request) {
	str, err := ioutil.ReadFile("project/http_Service/ttt.html")
	if err != nil {
		resw.Write([]byte("没有文件内容！"))
	}
	//str := ``
	resw.Write(str)

}
func main() {
	http.HandleFunc("/fengling", f1)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
