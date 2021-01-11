package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	config2 "goblog/pkg/config"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	config.Initialize()
}

func main() {

	bootstrap.SetUpDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":" + config2.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
