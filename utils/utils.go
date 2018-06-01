package utils

import (
	"encoding/json"
	"net/http"
)

// 向 http.ResponseWriter 写入 json 响应
func WriteJsonResponse(writer http.ResponseWriter, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	jsonResp, _ := json.Marshal(data)
	writer.Write(jsonResp)
}
