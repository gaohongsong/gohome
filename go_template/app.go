package main

import (
	"html/template"
	"log"
	"os"
	"fmt"
)

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
	t.ExecuteTemplate(os.Stdout, "header", data);fmt.Println()
	t.ExecuteTemplate(os.Stdout, "content", data);fmt.Println()
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
}
