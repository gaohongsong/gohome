package main
/*
	net/http包的基本使用方法
*/
import (
    "fmt"
    "net/http"
)

/* 自定义路由 */
type CustomRouter struct {
}

/* 实现 Handler 接口 */
func (p *CustomRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        helloWorld(w, r)
        return
    }
    http.NotFound(w, r)
    return
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func main() {
    //router := &CustomRouter{}
    //http.ListenAndServe(":8080", router)
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}