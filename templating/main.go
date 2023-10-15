package main

import (
	"os"
	"text/template"
)

type ConfigItems struct {
	Envr string
	Vars string
}

func main() {
	devCfg := ConfigItems{"dev", "tst"}
	tmpl, err := template.New("dev").Parse("{{.Vars}} in environment {{.Envr}}")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, devCfg)
	if err != nil {
		panic(err)
	}
}
