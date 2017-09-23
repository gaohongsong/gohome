package main

import (
	"html/template"
	"os"
	"fmt"
	"strings"
)

type Person struct {
	UserName string
	email    string
}

type Human struct {
	UserName string
	Emails   []string
	Houses   []int
	Friends  []*Friend
}

type Friend struct {
	Fname string
}

// 模板函数（tag）
func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find @
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	//fmt.Println(strings.Replace(s, "@", "at", -1))

	// replace 2 by "at"
	return substrs[0] + "at" + substrs[1]
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}} {{.email}}!")
	p := Person{UserName: "pitou"}
	//p := Person{UserName:"pitou", email:"abc@root.com"}
	t.Execute(os.Stdout, p)

	//	=================================RANGE/WITH
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
		Houses:   []int{1, 2, 3},
	}
	t1.Execute(os.Stdout, p1)

	//	=================================IF
	tEmpty := template.New("def")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)

	tTplVar := template.New("template test")
	tTplVar = template.Must(tTplVar.Parse(`
				{{with $x := "<img src='abcdef'/>"}}{{$x | html}}{{end}}
				{{with $x := "abcdefg" | printf "%x"}}{{$x}}{{end}}
    			{{with $x := "abcdefg"}} {{printf "%s" $x}} {{$x | printf "%v"}} {{end}}
    `))
	tTplVar.Execute(os.Stdout, nil)

	//	=================================FuncMap
	t2 := template.New("test")
	t2 = t2.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t2 = template.Must(t2.Parse(`hello {{.UserName}}!
	{{range .Emails}}
		an email {{.|emailDeal}}
	{{end}}
	{{range .Houses}}
		an house {{.|emailDeal}}
	{{end}}
	`))
	t2.Execute(os.Stdout, p1)

	//	=================================Must
	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}
