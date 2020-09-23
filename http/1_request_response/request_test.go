package tig_http_router

import (
	"fmt"
	gconv "github.com/og/x/conv"
	"log"
	"net/http"
	"testing"
)

var _ = `
# 启动HTTP服务
`

func TestNewServeMux(t *testing.T) {
	// 创建一个 http 服务器
	serve := http.NewServeMux()
	// 处理 url 为 / 的请求
	serve.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Write 传入 []byte 来响应请求 []byte("") 是将字符串转换为 []byte 的语法
		n, err := w.Write([]byte("og/thinking-in-go"))
		log.Print("write byte length: ", n)
		if err != nil { panic(err) }
	})
	// 处理 url 为 /news 的请求
	serve.HandleFunc("/news", func(w http.ResponseWriter,  r *http.Request) {
		n, err := w.Write([]byte("news page"))
		log.Print("write byte length: ", n)
		if err != nil { panic(err) }
	})
	addr := ":1001"
	log.Print("Listen http://127.0.0.1"+ addr)
	// 监听 1001 端口，http://127.0.0.1:1001
	err := http.ListenAndServe(addr, serve)
	if err != nil { panic(err) }
}

var _ = `
以上代码是使用 go 官方提供的 net/http 启动的 http 服务器
^http.ResponseWriter^ 用于控制响应内容
^*http.Request^ 用于获取 http 请求

接下来逐步介绍在 go 如何使用 http
`

var _ = `
## 绑定请求

### 获取 query (GET 参数)

`

func TestQuery(t *testing.T) {
	serve := http.NewServeMux()
	log.Print("打开 http://127.0.0.1:1002/values?title=og&names=a&names=b 查看响应")
	serve.HandleFunc("/query", func(w http.ResponseWriter,  r*http.Request) {
		// 获取 GET 请求
		query := r.URL.Query()
		// fmt.Sprintf 是将各种类型的变量转换为字符串的方法
		response := fmt.Sprintf("%#v", query)
		_, err := w.Write([]byte(response))
		if err != nil { panic(err) }
	})
	log.Print("打开 http://127.0.0.1:1002/query_get?title=og&names=a&names=b 查看响应")
	serve.HandleFunc("/query_get", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		response := fmt.Sprintf("%#v", "title:(" + query.Get("title") + ")" + "names:(" + query.Get("names") + ")")
		_, err := w.Write([]byte(response))
		if err != nil { panic(err) }
	})
	addr := ":1002"
	log.Print("Listen http://127.0.0.1"+ addr)
	err := http.ListenAndServe(addr, serve)
	if err != nil { panic(err) }
}

var _ = `
上述代码主要内容是 ^r.URL.Query().Get("title")^

net/url 包将 "title=og&names=a" 字符串转换为 ^type Values map[string][]string^。
通过给 Values 实现了 Get Set Add Del 等方法来快捷的获取GET参数(query string)。
比如 Get 的实现是：

^^^go
type Values map[string][]string

func (v Values) Get(key string) string {
	if v == nil {
		return ""
	}
	vs := v[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}
^^^

注意 Values 不是 map[string]string 而是 map[string]**[]**string
这是因为 url 中的 GET 参数会出现 names=a&names=b 这种多个相同 key 的情况。
但平时很少会使用多个相同key,故此 ^url.Values{}.Get(key string) string^ 返回的是 string 而不是 []string

但日常开发过程中逐行 query.Get(key) 比较繁琐，且容易出错，例如：

`

func TestStructAndQuery(t *testing.T) {
	serve := http.NewServeMux()
	log.Print("打开 http://127.0.0.1:1003/struct_and_query?name=nimoc&age=27 查看响应")
	serve.HandleFunc("/struct_and_query", func(w http.ResponseWriter,  r*http.Request) {
		// 获取 GET 请求
		query := r.URL.Query()
		type Req struct {
			Name string
			Age int
		}
		age, err := gconv.StringInt(query.Get("age"))
		if err != nil {
			_, err := w.Write([]byte("age 格式错误")) ; if err != nil {panic(err)}
			return
		}
		req := Req{
			Name: query.Get("name"),
			Age: age,
		}
		response := fmt.Sprintf("我是%s，今年%d岁", req.Name, req.Age)
		_, err = w.Write([]byte(response))
		if err != nil { panic(err) }
	})

	addr := ":1003"
	log.Print("Listen http://127.0.0.1"+ addr)
	err := http.ListenAndServe(addr, serve)
	if err != nil { panic(err) }
}


var _ = `

^^^go
age, err := gconv.StringInt(query.Get("age"))
// 省略 if err != nil 代码 
req := Req{
	Name: query.Get("name"),
	Age: age,
}
^^^

通过 query.Get(key) 取值非常繁琐，还要处理 string 到 int 的转换。

基于 github.com/og/juice 可快速绑定请求


`
type JE interface {RequestParse(r *http.Request, v interface{})}
var je JE

func TestOrangeParse(t *testing.T) {
	type Req struct {
		Name string `query:"name"`
		Age int `query:"age"`
	}
	req := Req{}
	// name=nimoc&age=27
	r := &http.Request{}
	je.RequestParse(r, &req)
}
