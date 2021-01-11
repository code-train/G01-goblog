package controllers

import (
	"fmt"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"gorm.io/gorm"
	"net/http"
)

type BaseController struct {

}

func (c BaseController) ResponseForSQLError(w http.ResponseWriter, err error)  {
	if err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Not Found")
	} else {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	}
}

func (c BaseController) ResponseForUnauthorized(w http.ResponseWriter, r *http.Request) {
	flash.Warning("未授权操作！")
	http.Redirect(w, r, "/", http.StatusFound)
}

