package main

import (
	"html/template"
	"log"
	"os"
	"fmt"
	"path/filepath"
	"io"
	"github.com/kataras/iris/core/errors"
)

var templates map[string]*template.Template

// 模板渲染
func renderTemplate(wr io.Writer, name string, data interface{}) error {
	tpl, ok := templates[name]
	if !ok {
		return errors.New("Template File Not Exist: " + name)
	}
	return tpl.ExecuteTemplate(wr, name, data)
}

// 模板预加载
func init() {
	// map初始化，分配内存
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	// 模板目录
	templateDir := "./app/"
	// 视图模板
	layouts, err := filepath.Glob(filepath.Join(templateDir, "layouts/*.html"))
	if err != nil {
		log.Fatal(err)
	}
	// 基础模板
	widgets, err := filepath.Glob(filepath.Join(templateDir, "widgets/*.html"))
	if err != nil {
		log.Fatal(err)
	}
	// index.html footer.html header.html
	for _, layout := range layouts {
		files := append([]string{layout}, widgets...)
		log.Printf("Init【%s】: %v", layout, files)
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(files...))
	}
}

func main() {

	// 以第一个为主模板，根据主模板构建嵌套树，注释部分将无法正确渲染嵌套关系
	//t := template.Must(template.ParseFiles(
	//	"app/content.html",
	//	"app/header.html",
	//	"app/footer.html",
	//	"app/content.html"))

	t := template.Must(template.ParseFiles(
		"app/index.html",
		"app/header.html",
		"app/footer.html",
		"app/content.html"))
	data := struct {
		Title string
	}{
		Title: "Alpha",
	}
	// 模板相互独立
	fmt.Println("================================================")
	t.ExecuteTemplate(os.Stdout, "header", data)
	fmt.Println()
	t.ExecuteTemplate(os.Stdout, "content", data)
	fmt.Println()
	t.ExecuteTemplate(os.Stdout, "footer", data)
	fmt.Println("================================================")

	// 一步到位
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, map[string]string{
		"Title": "Miya",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n================================= 预加载")
	err = renderTemplate(os.Stdout, "index.html", map[string]string{
		"Title": "Sleep?",
	})
	if err != nil {
		log.Fatal(err)
	}

}
