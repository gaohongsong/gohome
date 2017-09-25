package main

/*
	net/http包的基本使用方法
*/
import (
	"fmt"
	"net/http"
	"io"
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
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

func getParam(w http.ResponseWriter, r *http.Request) {
	//reqGet := r.URL.Query()
	//// get multiple key to array
	//names := reqGet["name"]
	//age := reqGet.Get("age")
	//// get none exist key to ""
	//notExist := reqGet.Get("balabala")
	//if notExist == "" {
	//	fmt.Println("notExist not exist in param")
	//}
	//for _, name := range names {
	//	fmt.Println(name)
	//}
	//fmt.Println(r.URL.Path, r.URL.String(), r.URL.RawPath, r.URL.RawQuery, names, age, notExist == "")
	//
	// 获取GET请求参数
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(r.Form.Get("age"), r.Form.Get("name"), r.Form["name"], r.Form.Get("EMTPY"))
	//fmt.Fprintf(w, "%s", r.Form)
	fmt.Printf("name: %s, age=%s\n", r.Form["name"], r.Form.Get("age"))
	fmt.Printf("name=%s, age=%s\n", r.FormValue("name"), r.FormValue("age"))
	fmt.Printf("name=%s, age=%s\n", r.PostFormValue("name"), r.PostFormValue("age"))
	fmt.Fprintf(w, "name=%s, age=%s\n", r.FormValue("name"), r.FormValue("age"))
	/*
	name: [pitou miya], age=123
	name=pitou, age=123
	name=, age=
	*/
}

func postFormData(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		log.Fatal(err)
	}
	//fmt.Fprintf(w, "Form: %s\n", r.Form)
	//fmt.Fprintf(w, "PostForm: %s\n", r.PostForm)
	//fmt.Fprintf(w, "MultipartForm: %s\n", r.MultipartForm)
	//fmt.Println(r.Form.Get("name"), r.Form.Get("age"))
	//fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	fmt.Fprintf(w, "name=%s, age=%s\n", r.PostFormValue("name"), r.PostFormValue("age"))
	fmt.Printf("name=%s, age=%s\n", r.PostFormValue("name"), r.PostFormValue("age"))
	fmt.Printf("name=%s, age=%s\n", r.FormValue("name"), r.FormValue("age"))
	fmt.Printf("name=%s, age=%s\n", r.PostForm["name"], r.PostForm.Get("age"))
	/*
	name=miya, age=12345
	name=miya, age=12345
	name=[miya pitou], age=12345
	*/

	// file upload
	file, handler, err := r.FormFile("test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)

	//f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
	f, err := os.Create("./" + handler.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	copySize, ok := io.Copy(f, file)
	fmt.Println(copySize, ok)
}

type Message struct {
	Name string `json:"name"`
	Body string    `json:"body"`
	Time int64    `json:"time,string"`
}


func postJsonData(w http.ResponseWriter, r *http.Request) {
	// dumps
	m := Message{"Alice", "Hello", 1294706395881547000}
	bb, err := json.Marshal(m)
	if err != nil {
		log.Fatal("%s", err)
	}
	fmt.Println(bb)
	//w.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(w, "%s", b)

	// loads
	//var data Message
	var data map[string]interface{}
	//bb := `{
	//"name":"rsj217@gmail.com",
	//"body":"123",
	//"time":100.5
	//}`
	err = json.Unmarshal([]byte(bb), &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)

	// read from body and return
	// curl -v -XPOST localhost:8000/json/ -H "Content-Type:application/json" -d '{"name":"miya","body":"body","time":"11234"}'
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Read Err: %s", err)
	}

	//err = json.Unmarshal(body, &msg)
	err = json.Unmarshal(body, &data)
	fmt.Println(data, data["name"])
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", body)
}

func main() {
	//router := &CustomRouter{}
	//http.ListenAndServe(":8080", router)
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/get/", getParam)
	//curl -XPOST localhost:8000/post/ -F "name=miya" -F "age=123" -F "test=@test.txt"
	http.HandleFunc("/post/", postFormData)
	http.HandleFunc("/json/", postJsonData)
	http.ListenAndServe(":8000", nil)
}
