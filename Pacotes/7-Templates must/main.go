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
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}}\nCarga Horária: {{.CargaH}}h"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
