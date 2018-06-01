package data

import (
	"log"
)

type Url struct {
	Id          int
	Url         string
	Logo        string
	Introduce   string
	IsStar      int
	OrderNumber int
	CategoryId  int
}

// insert category to DB
func (url *Url) AddUrl() (insertId int64, err error) {
	// prepare the sql statement
	sqlStatement := "INSERT INTO url (url,logo,introduce,is_star,order_number,category_id) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err)
	}

	// defer for close stmt
	defer stmt.Close()

	// exec stmt
	result, err := stmt.Exec(url.Url, url.Logo, url.Introduce, url.IsStar, url.OrderNumber, url.CategoryId)
	if err != nil {
		log.Println(err)
	}

	insertId, err = result.LastInsertId()
	return
}

// query urls by categoryId
func QueryUrlsByCategoryId(categoryId int) (urls []Url, err error) {
	sqlStatement := "SELECT id,url,logo,introduce,is_star,order_number FROM url WHERE category_id=? ORDER BY order_number ASC"
	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err)
	}

	rows, err := stmt.Query(categoryId)
	defer rows.Close()

	for rows.Next() {
		url := Url{}
		if err = rows.Scan(&url.Id, &url.Url, &url.Logo, &url.Introduce, &url.IsStar, &url.OrderNumber); err != nil {
			return
		}
		urls = append(urls, url)
	}
	return
}

// query all urls
func QueryAllUrls() (urls []Url, err error) {
	sqlStatement := "SELECT id,url,logo,introduce,is_star,order_number FROM url"
	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err)
	}

	rows, err := stmt.Query()
	defer rows.Close()

	for rows.Next() {
		url := Url{}
		if err = rows.Scan(&url.Id, &url.Url, &url.Logo, &url.Introduce, &url.IsStar, &url.OrderNumber); err != nil {
			return
		}
		urls = append(urls, url)
	}
	return
}
