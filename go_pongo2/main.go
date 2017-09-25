package main

import (
	"fmt"
	"github.com/flosch/pongo2"
	"net/http"
)

func init() {
}

func main() {

	//tpl, err := pongo2.FromString("Hello {{ name|capfirst }}!")
	// Compile the template at application start (for performance reasons)
	//if err != nil {
	//    panic(err)
	//}

	// 加载模板
	var tpl = pongo2.Must(pongo2.FromFile("index.html"))

	// 渲染模板
	//todos := []string{"dou", "ruai", "mi", "fa", "so", "la", "xi"}
	todos := []struct {
		name string
		ok   bool
	}{
		{"吃饭", false},
		{"睡觉", false},
		{"打怪", false},
		{"撩妹", false},
		{"呵呵", true},
	}
	out, err := tpl.Execute(pongo2.Context{
		"current_date": "florian",
		"todos":        todos,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out) // Output: Hello Florian!

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteWriter(pongo2.Context{
			"current_date": "florian",
			"todos":        todos,
		}, w)
	})

	http.ListenAndServe(":8000", nil)
}
