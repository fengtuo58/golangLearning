package main

import "html/template"
import "os"
import "log"
func main() {
	const temp1 = `<p>A:{{.A}}</p><p>B:{{.B}}</p>`
	t := template.Must(template.New("escape").Parse(temp1))
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>hello!</b>"
	data.B = "<b>hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
