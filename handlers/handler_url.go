package handlers

import (
	"net/http"
	"strconv"

	"github.com/xiaosongfu/url/data"
	"github.com/xiaosongfu/url/model"
	"github.com/xiaosongfu/url/services"
	"github.com/xiaosongfu/url/utils"
)

func AddUrl(writer http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := data.Url{}
	url.Url = r.FormValue("url")
	url.Logo = r.FormValue("logo")
	url.Introduce = r.FormValue("introduce")

	// is_star
	if isStar, err := strconv.Atoi(r.FormValue("isStar")); err != nil {
		url.IsStar = 0 // 0 是默认值
	} else {
		url.IsStar = isStar
	}

	// order_number
	if orderNumber, err := strconv.Atoi(r.FormValue("orderNumber")); err != nil {
		url.OrderNumber = 0 // 0 是默认值
	} else {
		url.OrderNumber = orderNumber
	}

	// category_id
	if categoryId, err := strconv.Atoi(r.FormValue("categoryId")); err != nil {
		url.CategoryId = 1 // 0 是未分类
	} else {
		url.CategoryId = categoryId
	}

	resp := model.BasicResponse{}
	_, err := services.AddUrl(&url)
	if err != nil {
		resp.Code = model.ResponseCodeFailed
		resp.Message = err.Error()
	} else {
		resp.Code = model.ResponseCodeSuccess
		resp.Message = "add url success"
	}
	utils.WriteJsonResponse(writer, resp)
}
