package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email    string
}

type Human struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

type Friend struct {
	Fname string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}} {{.email}}!")
	p := Person{UserName: "pitou"}
	//p := Person{UserName:"pitou", email:"abc@root.com"}
	t.Execute(os.Stdout, p)

	//	=================================
	f1 := Friend{Fname: "miya"}
	f2 := Friend{Fname: "alpha"}
	t1 := template.New("abc")

	t1, _ = t1.Parse(`hello {{.UserName}}!
	{{range .Emails}}
		an email {{.}}
	{{end}}
	{{with .Friends}}
	{{range .}}
		my friend name is {{.Fname}}
	{{end}}
	{{end}}
	`)
	p1 := Human{
		UserName: "hongsong",
		Emails:   []string{"a@a.com", "b@b.com"},
		Friends:  []*Friend{&f1, &f2},
	}
	t1.Execute(os.Stdout, p1)

	//	=================================
    tEmpty := template.New("def")
    tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
    tEmpty.Execute(os.Stdout, nil)

    tWithValue := template.New("template test")
    tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
    tWithValue.Execute(os.Stdout, nil)

    tIfElse := template.New("template test")
    tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
    tIfElse.Execute(os.Stdout, nil)
}
