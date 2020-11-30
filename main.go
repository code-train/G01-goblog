package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r * http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "<h1>关于我</h1>")
	} else {
		fmt.Fprint(w, "<h1>未找到页面</h1>")
	}
}

func main(){
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}