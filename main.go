package main

import (
	"net/http"

	"log"

	"github.com/xiaosongfu/url/config"
	"github.com/xiaosongfu/url/handlers"
)

func main() {
	// 静态文件服务器
	fileServer := http.FileServer(http.Dir("web/public"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// 配置路由
	// --> 首页
	http.HandleFunc("/", handlers.Logger(handlers.Index))
	// --> addCategory
	http.HandleFunc("/api/addCategory", handlers.Logger(handlers.AddCategory))
	// --> addUrl
	http.HandleFunc("/api/addUrl", handlers.Logger(handlers.AddUrl))

	// 获取配置
	addr := config.Conf.Server[config.Conf.Env].Addr

	// 打印日志
	log.Println("server started,and listening on 0.0.0.0" + addr)

	// 启动服务
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(nil)
	}
}
