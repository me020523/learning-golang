package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}
type Person struct {
	Username string
	Emails []string
	Friends []*Friend
	Single bool
}

func demoSingeField(){
	t := template.New("fieldName example")
	t, _ = t.Parse("hello, {{.Username}}")
	p := Person{
		Username : "bcshuai",
	}
	t.Execute(os.Stdout, p)
}
func demoLoop(){
	t := template.New("demo loop")
	t, _ = t.Parse(`
			hello {{.Username}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
			{{range .}}
				friend {{.Fname}}
			{{end}}
			{{end}}
		`)
	fs := []*Friend{
		&Friend{Fname: "TOM"},
		&Friend{Fname: "Lily"},
		&Friend{Fname: "Jack"},
	}
	p := Person{
		Username : "Mike",
		Emails: []string{"123@test.com", "456@test.com"},
		Friends: fs,
	}
	t.Execute(os.Stdout, p);
}
func demoCondition(){
	t := template.New("demo condition")
	t, _ = t.Parse(`
			hello {{.Username}}
			{{if .Single}}
				I'm Single
			{{else}}
				I have getten married
			{{end}}
		`)
	p := Person{
		Username : "Tome",
		Single : false,
	}
	t.Execute(os.Stdout, p)
}
func main(){
	//demoLoop()
	demoCondition()
}