package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	myTemplate := `
	Name:  &lt {{.Name}}
	Age: {{.Age}}
	Robbies:
	{{range .Robbies}}
	Robbie Name: {{.Name}}
	{{end}}
	`
	personInfo := template.Must(template.New("info").Parse(myTemplate))

	person := struct {
		Name    string
		Age     int
		Robbies []struct{ Name string }
	}{Name: "Davi", Age: 22, Robbies: []struct{ Name string }{{Name: "Play"}, {Name: "Sleep"}}}

	if err := personInfo.Execute(os.Stdout, person); err != nil {
		log.Fatal(err)
	}

	const temp = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("exemplohtml").Parse(temp))
	var data struct {
		A string        // valor literal
		B template.HTML // renderiza o html
	}
	data.A = "<b>Hello</b>"
	data.B = "<b>Hello</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Panic(err)
	}
}
