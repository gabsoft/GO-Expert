package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome   string
	CargaH int
}

func main() {
	curso := Curso{"Go", 40}
	tmpl := template.New("CursoTemplate")
	tmpl, _ = tmpl.Parse("Curso: {{.Nome}}\nCarga Horária: {{.CargaH}}h")
	err := tmpl.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
