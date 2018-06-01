package services

import "github.com/xiaosongfu/url/data"

// AddUrl
func AddUrl(url *data.Url) (insertId int64, err error) {
	insertId, err = url.AddUrl()
	return
}

// QueryAllUrls
func QueryAllUrls() (urls []data.Url, err error) {
	urls, err = data.QueryAllUrls()
	return
}
