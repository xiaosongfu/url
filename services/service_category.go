package services

import "github.com/xiaosongfu/url/data"

// AddCategory
func AddCategory(category *data.Category) (insertId int64, err error) {
	insertId, err = category.AddCategory()
	return
}
