# url
url collection

---

## 静态文件服务器  
```go
// 静态文件服务器
fileServer := http.FileServer(http.Dir("web/public"))
http.Handle("/static/", http.StripPrefix("/static/", fileServer))
```
使用 FileServer 函数创建了一个能够为指定目录(<project root>/web/public)中的静态文件服务的处理器，
并将这个处理器传递给多路复用器，此外，还使用 StripPrefix 函数去除请求 url 中的 static 前缀。  

当服务器接收到一个以 /static/ 开头的 URL 请求时，以上两行代码会移除 URL 中的 /static/ ，
然后在 <project root>/web/public 目录中查找被请求的文件，如接收到： http://localhost/static/css/basic.css 请求时，
它将会在 <project root>/web/public 目录中查找一下文件： css/basic.css 。
