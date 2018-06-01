package handlers

import (
	"net/http"

	"strconv"

	"github.com/xiaosongfu/url/data"
	"github.com/xiaosongfu/url/model"
	"github.com/xiaosongfu/url/services"
	"github.com/xiaosongfu/url/utils"
)

func AddCategory(writer http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	category := data.Category{}
	category.CategoryName = r.FormValue("categoryName")
	category.CategoryLogo = r.FormValue("categoryLogo")
	if id, err := strconv.Atoi(r.FormValue("parentCategoryId")); err != nil {
		category.ParentCategoryId = 1 // 1 是未分类
	} else {
		category.ParentCategoryId = id
	}

	resp := model.BasicResponse{}
	_, err := services.AddCategory(&category)
	if err != nil {
		resp.Code = model.ResponseCodeFailed
		resp.Message = err.Error()
	} else {
		resp.Code = model.ResponseCodeSuccess
		resp.Message = "add category success"
	}
	utils.WriteJsonResponse(writer, resp)
}
