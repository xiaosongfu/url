package handlers

import (
	"html/template"
	"net/http"

	"github.com/xiaosongfu/url/services"
)

func Index(writer http.ResponseWriter, r *http.Request) {
	// 1 获取数据
	urls, err := services.QueryAllUrls()
	if err != nil {
		writer.Write([]byte("server error"))
		return
	}

	// 2 渲染模板
	temp, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		writer.Write([]byte("server error"))
		return
	}
	temp.Execute(writer, urls)
}
