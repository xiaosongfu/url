package data

import (
	"log"
)

type Category struct {
	Id               int
	CategoryName     string
	CategoryLogo     string
	ParentCategoryId int
}

// insert category to DB
func (category *Category) AddCategory() (insertId int64, err error) {
	sqlStatement := "INSERT INTO category (category_name,category_logo,parent_category_id) VALUES (?, ?, ?)"
	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(category.CategoryName, category.CategoryLogo, category.ParentCategoryId)
	if err != nil {
		log.Println(err)
	}

	insertId, err = result.LastInsertId()
	return
}

// query all categorys
func Categorys() (categorys []Category, err error) {
	sqlStatement := "SELECT id,category_name,category_logo FROM category"
	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	defer rows.Close()

	for rows.Next() {
		category := Category{}
		if err = rows.Scan(&category.Id, &category.CategoryName, &category.CategoryLogo); err != nil {
			return
		}
		categorys = append(categorys, category)
	}
	return
}
