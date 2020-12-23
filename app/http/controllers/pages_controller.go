package controllers

import (
	"fmt"
	"net/http"
)

// PageController struct
type PageController struct {
}

// Home 首页
func (*PageController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, 欢迎来到 goblog!</h1>")
}

// About 关于我们
func (*PageController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>此博客是用以记录编程笔记，如您有反馈或建议，请联系</h1>")
}

// NotFound 404
func (*PageController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404")
}
